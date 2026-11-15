---
work_package_id: WP01
title: Backend Containerization Foundation
lane: "done"
dependencies: []
base_branch: main
base_commit: a66bdf89ffad2bb6cfe9fad6c8d6f0f57260a5d2
created_at: '2026-06-08T13:13:13.699177+00:00'
subtasks: [T001, T002, T003, T004, T005]
requirement_refs:
- FR-001
- FR-004
shell_pid: "3240"
agent: "gemma-4-31b-it"
reviewed_by: "Tiago Ramos"
review_status: "approved"
---

# Work Package: Backend Containerization Foundation

## Objective
Containerize all backend services and implement health endpoints to enable reliable orchestration.

## Context & Constraints
- **Goal**: Create the baseline containers and health checks required for Docker Compose sequencing.
- **Standard**: .NET 10 best practices for containerization.
- **Tools**: Docker, .NET SDK.

## Subtasks & Detailed Guidance

### Subtask T001: Implement .env files for all backend services and Postgres
**Purpose**: Externalize configuration for environment-specific variables.
**Steps**:
1. Create `.env.postgres`, `.env.identity`, `.env.events`, and `.env.registrations`.
2. Define necessary variables (e.g., `DB_PASSWORD`, `CONNECTION_STRING`, `PORT`).
3. Ensure no secrets are hardcoded in the service logic.
**Validation**:
- [ ] All four `.env` files exist with required variables.

### Subtask T002: Create Dockerfiles for Identity, Events, and Registrations services
**Purpose**: Define the build process for the .NET 10 microservices.
**Steps**:
1. Create a `Dockerfile` for each service.
2. Use a multi-stage build (build stage with SDK, runtime stage with ASP.NET runtime).
3. Ensure the app is published in `Release` mode.
4. Expose the internal port (typically 80).
**Validation**:
- [ ] Services can be built successfully using `docker build`.

### Subtask T003: Implement `/health` endpoint in IdentityService
**Purpose**: Provide a mechanism for orchestration tools to verify service readiness.
**Steps**:
1. Add `Microsoft.Extensions.Diagnostics.HealthChecks` NuGet package.
2. Configure the health check service in `Program.cs`.
3. Map the `/health` endpoint to return a simple `200 OK` when the service is running.
**Validation**:
- [ ] `GET /health` returns 200 when the service is healthy.

### Subtask T004: Implement `/health` endpoint in EventsService
**Purpose**: Provide a mechanism for orchestration tools to verify service readiness.
**Steps**:
1. Add `Microsoft.Extensions.Diagnostics.HealthChecks` NuGet package.
2. Configure the health check service in `Program.cs`.
3. Map the `/health` endpoint to return a simple `200 OK` when the service is running.
**Validation**:
- [ ] `GET /health` returns 200 when the service is healthy.

### Subtask T005: Implement `/health` endpoint in RegistrationsService
**Purpose**: Provide a mechanism for orchestration tools to verify service readiness.
**Steps**:
1. Add `Microsoft.Extensions.Diagnostics.HealthChecks` NuGet package.
2. Configure the health check service in `Program.cs`.
3. Map the `/health` endpoint to return a simple `200 OK` when the service is running.
**Validation**:
- [ ] `GET /health` returns 200 when the service is healthy.

## Definition of Done
- [ ] All three backend services have a valid `Dockerfile`.
- [ ] All services have a working `/health` endpoint.
- [ ] Environment variables are fully externalized to `.env` files.

## Risks & Mitigations
- **Risk**: Health checks might be too simple (only checking if app is up, not if DB is connected).
- **Mitigation**: Focus on "Readiness" here; deeper DB checks are handled by `pg_isready` in the orchestrator.

## Activity Log

- 2026-06-08T13:13:20Z – gemma-4-31b-it – shell_pid=13196 – lane=doing – Assigned agent via workflow command
- 2026-06-08T13:16:24Z – gemma-4-31b-it – shell_pid=13196 – lane=for_review – Ready for review: Implemented .env files, Dockerfiles for backend services, and /health endpoints for readiness checks.
- 2026-06-08T13:28:55Z – gemma-4-31b-it – shell_pid=3240 – lane=doing – Started review via workflow command
- 2026-06-08T13:29:18Z – gemma-4-31b-it – shell_pid=3240 – lane=done – Review passed: Implementation correctly containerizes backend services and implements /health endpoints.
