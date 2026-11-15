# Tasks: Infrastructure Connectivity & Router Modernization

## Work Packages

### WP01: Backend Containerization Foundation
**Summary**: Containerize all backend services and implement health endpoints to enable reliable orchestration.
**Priority**: High
**Requirements Refs**: FR-001, FR-004
**Included Subtasks**:
- [x] T001: Implement .env files for all backend services and Postgres. [P]
- [x] T002: Create Dockerfiles for Identity, Events, and Registrations services. [P]
- [x] T003: Implement `/health` endpoint in IdentityService.
- [x] T004: Implement `/health` endpoint in EventsService.
- [x] T005: Implement `/health` endpoint in RegistrationsService.

**Implementation Sketch**:
1. Define environment variables in separate `.env` files for each service.
2. Create optimized Dockerfiles for the .NET 10 services.
3. Add a simple health check controller/endpoint to each service that returns 200 OK.

**Parallel Opportunities**: T001 and T002 can be done in parallel.
**Dependencies**: None.
**Risks**: Inconsistent health check implementations across services.
**Estimated Prompt Size**: ~350 lines.

---

### WP02: Orchestration & Networking
**Summary**: Configure Docker Compose for network isolation, service discovery, and a single API Gateway entry point.
**Priority**: High
**Requirements Refs**: FR-001, FR-002, FR-003, FR-007
**Included Subtasks**:
- [x] T006: Configure `docker-compose.yml` with network bridge and PostgreSQL.
- [x] T007: Implement container health checks and startup sequencing in `docker-compose.yml`.
- [x] T008: Configure Nginx API Gateway (`nginx.conf`) for routing services via port 8080.
- [x] T009: Add Nginx Gateway to `docker-compose.yml` and verify connectivity to one service.

**Implementation Sketch**:
1. Set up the `admin-portal-network` bridge in `docker-compose.yml`.
2. Configure PostgreSQL and the three backend services with the correct network and environment files.
3. Use `depends_on` with `service_healthy` based on the `/health` endpoints.
4. Configure Nginx to proxy `/api/identity`, `/api/events`, and `/api/registrations` to the respective internal containers.

**Parallel Opportunities**: None.
**Dependencies**: WP01
**Risks**: Port conflicts on the host machine; Nginx configuration errors.
**Estimated Prompt Size**: ~400 lines.

---

### WP03: Frontend Stability & Modernization
**Summary**: Refactor navigation guards to eliminate framework warnings and implement resilient network error handling.
**Priority**: Medium
**Requirements Refs**: FR-005, FR-006
**Included Subtasks**:
- [x] T010: Audit and refactor Vue Router navigation guards to eliminate deprecation warnings. [P]
- [x] T011: Implement a global Axios interceptor for network error detection.
- [x] T012: Implement a Toast Notification system for connectivity alerts.
- [x] T013: Integrate Axios interceptor with Toast system.
- [x] T014: Verify full end-to-end flow: `docker-compose up` -> Frontend access -> API communication -> Error handling.

**Implementation Sketch**:
1. Update Vue Router guards to return values instead of using `next()`.
2. Create a global Axios instance with a response interceptor that detects `ERR_NETWORK`.
3. Build a simple Toast component and a store/service to trigger it.
4. Link the interceptor to the Toast service.
5. Perform a full system walk-through to validate all connectivity and stability goals.

**Parallel Opportunities**: T010 can be done in parallel with backend work.
**Dependencies**: WP02
**Risks**: Breaking existing authorization logic during guard refactor.
**Estimated Prompt Size**: ~450 lines.
