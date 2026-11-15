# Implementation Plan: Infrastructure Connectivity & Router Modernization

## 1. Technical Context
### Current State
- Backend services (Identity, Events, Registrations) and PostgreSQL are intended to be containerized but currently suffer from `ERR_CONNECTION_REFUSED` errors.
- Frontend runs locally via `npm run dev` and uses an outdated Vue Router navigation guard pattern causing console warnings.
- Connectivity failures are not gracefully handled in the UI.

### Technical Choices
- **Orchestration**: Docker Compose with internal DNS service names.
- **Connectivity**: Single API Gateway port mapped to the host for all backend traffic.
- **Health Monitoring**: HTTP `/health` endpoints for service readiness.
- **Configuration**: Service-specific `.env` files.
- **Frontend**: Vue Router 4 navigation guards and Axios interceptors for error handling.
- **UI**: Toast Notification system for connectivity alerts.

### Constitution Check
- **Frontend (Vue 3, Tailwind, PrimeVue)**: Aligns with the requirement for router modernization and toast notifications.
- **Backend (.NET 10)**: Aligns with the target for containerized ASP.NET Core services.
- **Deployment (Docker/K8s)**: Using Docker Compose as the local development prerequisite for K8s.
- **Resilience (Backoff/Correlation)**: The requested resilient Axios error handling aligns with the resilience standards.

## 2. Implementation Phases

### Phase 0: Outline & Research
**Goal**: Resolve technical unknowns regarding container orchestration and router patterns.

- [ ] **Research Task 1**: Docker Compose health check patterns for .NET 10 and PostgreSQL (e.g., `pg_isready` and custom HTTP health checks).
- [ ] **Research Task 2**: Implementation of a lightweight API Gateway for local dev (e.g., YARP or Nginx) to provide a single host port.
- [ ] **Research Task 3**: Vue Router 4 navigation guard patterns for protecting routes and eliminating deprecation warnings.
- [ ] **Research Task 4**: Axios interceptor implementation for detecting and handling `ERR_CONNECTION_REFUSED` specifically.

**Output**: `research.md` documenting decisions and rationales.

### Phase 1: Design & Contracts
**Goal**: Define the infrastructure layout and API boundaries.

- [ ] **Data Model**: Define the infrastructure entities (Service, Port, Network, HealthCheck) in `data-model.md`.
- [ ] **API Contracts**: Define the Gateway routing and health check endpoints in `/contracts/`.
- [ ] **Quickstart Guide**: Create `quickstart.md` documenting the `docker-compose up` flow and environment variable setup.
- [ ] **Agent Context Update**: Update agent-specific context files with new infrastructure patterns.

**Output**: `data-model.md`, `/contracts/*`, `quickstart.md`, updated agent context.

## 3. Gates
- **Gate 1 (Post-Research)**: All research tasks completed and decisions documented in `research.md`.
- **Gate 2 (Post-Design)**: API contracts defined and `quickstart.md` reviewed for clarity.

## 4. Assumptions
- The developer has Docker and Docker Compose installed.
- The backend services already have a basic `/health` endpoint or can easily implement one.
- The frontend uses a standard Axios instance for API requests.
