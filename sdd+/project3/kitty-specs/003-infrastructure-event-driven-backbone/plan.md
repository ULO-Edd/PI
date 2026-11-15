# Implementation Plan: Infrastructure Event-Driven Backbone

**Branch**: `003-infrastructure-event-driven-backbone` | **Date**: 2026-06-02 | **Spec**: [spec.md](spec.md)  
**Input**: Feature specification from `kitty-specs/003-infrastructure-event-driven-backbone/spec.md`

## Summary

This feature delivers the decoupled, resilient asynchronous messaging infrastructure and background processing layer for the platform. It establishes a RabbitMQ message broker topology for routing check-in and notification events. Background Analytics and Notification worker services, implemented with ASP.NET Core (.NET 10), process events asynchronously. Reliability is guaranteed via a 3-attempt exponential backoff retry policy, with permanent failures routed to a dedicated Dead-Letter Queue (DLQ). Failed messages are persisted in an isolated relational database, and secure API Gateway management endpoints are provided for Admins to inspect, replay (retaining original correlation IDs), or purge DLQ entries. Local orchestration is handled via Docker Compose, with production-ready Kubernetes YAML manifests for container deployment, secret mounting, and health checks.

## Technical Context

- **Language/Version**: C# (.NET 10)  
- **Primary Dependencies**: ASP.NET Core (.NET 10), RabbitMQ.Client (v7.0.0+), Microsoft.EntityFrameworkCore.Sqlite (v10) for local failure tracking, Microsoft.EntityFrameworkCore (v10)  
- **Storage**: Isolated SQLite databases per worker / management service (local development)  
- **Testing**: xUnit for unit and worker logic testing, Playwright for .NET for integration flow testing  
- **Target Platform**: Linux (Docker container images, Kubernetes orchestration)  
- **Project Type**: Web Application / Web APIs and Background Services (under `/backend` root directory)  
- **Performance Goals**: Message transmission latency under 50ms at p95 under standard load; messaging bus ingests and buffers bursts of 500+ concurrent events without message loss  
- **Constraints**: Maximum of exactly 3 retries with exponential backoff before routing to DLQ; replayed messages must retain original correlation ID; strict role-restricted Admin access for DLQ management  
- **Scale/Scope**: Background worker processing rates scale dynamically to consume up to 1,000 events per second; zero-downtime rolling updates supported via container orchestration  

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

- **Backend Tech Stack**: ASP.NET Core (.NET 10) used for background workers matches the backend standards. (PASS)
- **Database Isolation Constraints**: Dead-Letter/Failure storage maintains an isolated relational database separate from other services. (PASS)
- **Deployment and Containerization**: Local Docker Compose and production Kubernetes manifests match the Docker-only targeting Kubernetes requirement. (PASS)
- **Resilience**: RabbitMQ integration utilizes structured correlation IDs and exponential backoff retry policies matching the resilience standards. (PASS)

All gates successfully passed. No constitution violations detected.

## Project Structure

### Documentation (this feature)

```
kitty-specs/003-infrastructure-event-driven-backbone/
├── spec.md              # Feature specification
├── plan.md              # This file (/spec-kitty.plan command output)
├── research.md          # Phase 0 output
├── data-model.md        # Phase 1 output
├── quickstart.md        # Phase 1 output
├── contracts/           # Phase 1 output
└── checklists/
    └── requirements.md  # Quality checklist
```

### Source Code Layout

The infrastructure and services are organized under the centralized `/backend` subdirectory:

```
backend/
├── src/
│   ├── ApiGateway/              # Perimeter gateway router with routes for DLQ management
│   ├── Services/
│   │   ├── AnalyticsWorker/     # Background worker consuming check-in events for dashboards
│   │   │   ├── Program.cs
│   │   │   └── Workers/
│   │   │       └── CheckInEventConsumer.cs
│   │   ├── NotificationWorker/  # Background worker consuming check-in events for notifications
│   │   │   ├── Program.cs
│   │   │   └── Workers/
│   │   │       └── CheckInEventConsumer.cs
│   │   └── DeadLetterService/   # Relational storage service for failed DLQ messages
│   │       ├── Controllers/
│   │       │   └── DeadLetterController.cs
│   │       ├── Data/            # EF Core DBContext for failure logging
│   │       └── Program.cs
│   └── Shared/                  # Shared integration event schemas and correlation helpers
├── deploy/
│   ├── local/
│   │   └── docker-compose.yml   # Multi-container local execution manifest
│   └── k8s/
│       ├── rabbitmq.yaml        # Kubernetes StatefulSet / Service for RabbitMQ
│       ├── workers.yaml         # Kubernetes Deployments for Analytics & Notification workers
│       ├── deadletter.yaml      # Kubernetes Deployment and Service for Dead-Letter service
│       └── configmaps-secrets.yaml # Production environment secret mapping configurations
└── tests/
    └── WorkerIntegrationTests/  # Worker consumer integration and retry policy tests
```

**Structure Decision**: Workers and DLQ persistence service structured under `/backend/src/Services/` along with Shared and ApiGateway components. Orchestration configurations are centralized in `/backend/deploy/` directory to manage local development (docker-compose) and production environments (k8s manifests).

## Complexity Tracking

*No constitution violations or complex architecture deviations detected.*
