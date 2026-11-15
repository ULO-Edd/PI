# Implementation Plan: Backend Microservices Synchronous Core

**Branch**: `002-backend-microservices-synchronous-core` | **Date**: 2026-06-01 | **Spec**: [spec.md](spec.md)  
**Input**: Feature specification from `kitty-specs/002-backend-microservices-synchronous-core/spec.md`

## Summary

This feature delivers the high-performance, strictly typed synchronous backend core for SDD-Project3. Built with ASP.NET Core (.NET 10), it contains an API Gateway routing entry, an Identity Service managing role contracts, and two distinct domain services (Events and Registrations) housing core business logic. All services maintain fully isolated relational databases and leverage Entity Framework Core (EF Core 10) Code-First migrations, ensuring zero cross-database queries or joins. Synchronous capacity validation is strictly enforced via secure inter-service HTTP requests (Option A) with automated transactional rollback on failure.

## Technical Context

- **Language/Version**: C# (.NET 10)  
- **Primary Dependencies**: ASP.NET Core, Microsoft.EntityFrameworkCore (v10), Microsoft.AspNetCore.Authentication.JwtBearer, Swashbuckle.AspNetCore (OpenAPI generation)  
- **Storage**: Isolated relational databases per microservice managed by EF Core 10 Code-First migrations  
- **Testing**: Playwright for .NET (for E2E and HTTP integration checks), xUnit (for controller and domain logic unit testing)  
- **Target Platform**: Linux (Docker container images, Kubernetes orchestration)  
- **Project Type**: Web Application / Web APIs (under `/backend` root directory)  
- **Performance Goals**: API Gateway endpoints respond in < 200ms at p95; capacity reservation and ticket booking requests process in < 500ms  
- **Constraints**: No cross-database joins or queries; strict JWT authentication guards; role-based access control (Admin, Organiser, Attendee); rate limiting configured by identity  
- **Scale/Scope**: System rate-limiting (e.g., 100 req/min for Attendees, 1000 req/min for Organisers) to handle event catalogue loads of up to 10,000 public records  

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

- **Backend Tech Stack**: ASP.NET Core (.NET 10) with isolated relational databases per microservice matches exactly. (PASS)
- **Database Isolation Constraints**: 0 cross-database joins or queries enforced. (PASS)
- **Deployment and Containerization**: Docker-only targeting Kubernetes matches exactly. (PASS)
- **Security & Resilience standards**: JWT-based role check validation, and structured correlation ID with transaction rollback meet standards. (PASS)

No constitution violations detected.

## Project Structure

### Documentation (this feature)

```
kitty-specs/002-backend-microservices-synchronous-core/
├── spec.md              # Feature specification
├── plan.md              # This file (/spec-kitty.plan command output)
├── research.md          # Phase 0 output
├── data-model.md        # Phase 1 output
├── quickstart.md        # Phase 1 output
├── contracts/           # Phase 1 output (OpenAPI schemas)
│   ├── identity-contracts.json
│   ├── events-contracts.json
│   └── registrations-contracts.json
└── checklists/
    └── requirements.md  # Quality checklist
```

### Source Code Layout

The backend services will be organized under a centralized `/backend` subdirectory:

```
backend/
├── src/
│   ├── ApiGateway/              # Perimeter router, Jwt authentication, and rate limiting
│   │   ├── Program.cs
│   │   └── appsettings.json
│   ├── Services/
│   │   ├── IdentityService/     # User registrations and role mappings (Admin, Organiser, Attendee)
│   │   │   ├── Controllers/
│   │   │   ├── Data/            # EF Core DBContext, entities, and migrations
│   │   │   ├── Models/
│   │   │   └── Program.cs
│   │   ├── EventsService/       # Event catalog management and capacity track validation
│   │   │   ├── Controllers/
│   │   │   ├── Data/            # EF Core DBContext, entities, and migrations
│   │   │   ├── Models/
│   │   │   └── Program.cs
│   │   └── RegistrationsService/ # Ticket bookings and transactions handling
│   │       ├── Controllers/
│   │       ├── Data/            # EF Core DBContext, entities, and migrations
│   │       ├── Models/
│   │       └── Program.cs
│   └── Shared/                  # Shared models, validation filters, and correlation middleware
└── tests/
    ├── UnitTests/               # xUnit controller and store logic units
    └── IntegrationTests/        # Playwright .NET API integration and database checks
```

**Structure Decision**: Microservices project structures under `/backend/src/Services/` combined with a shared infrastructure gateway project. Each microservice contains its own Data context folder to completely isolate schema migrations and databases.

## Complexity Tracking

*No constitution violations or complex architecture deviations detected.*
