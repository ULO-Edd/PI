---
work_package_id: WP02
title: Orchestration & Networking
lane: "done"
dependencies: [WP01]
base_branch: 005-infrastructure-connectivity-router-modernization-WP01
base_commit: 0a460b46a2055778804ac99c7b3d4ee374df978a
created_at: '2026-06-08T13:18:16.567677+00:00'
subtasks: [T006, T007, T008, T009]
requirement_refs:
- FR-001
- FR-002
- FR-003
- FR-007
shell_pid: "3792"
agent: "gemma-4-31b-it"
reviewed_by: "Tiago Ramos"
review_status: "approved"
---

# Work Package: Orchestration & Networking

## Objective
Configure Docker Compose for network isolation, service discovery, and a single API Gateway entry point.

## Context & Constraints
- **Goal**: A one-command startup that ensures correct sequencing and connectivity.
- **Standard**: Docker Compose V2, Nginx for proxying.
- **Tools**: `docker-compose.yml`, `nginx.conf`.

## Subtasks & Detailed Guidance

### Subtask T006: Configure `docker-compose.yml` with network bridge and PostgreSQL
**Purpose**: Establish the core infrastructure layout.
**Steps**:
1. Create a `docker-compose.yml` file.
2. Define a network bridge named `admin-portal-network`.
3. Add the `postgres` service using the official image.
4. Add the three backend services, linking them to the bridge and providing their respective `.env` files.
**Validation**:
- [ ] `docker-compose up` starts all containers without network errors.

### Subtask T007: Implement container health checks and startup sequencing
**Purpose**: Prevent services from crashing during startup due to missing dependencies.
**Steps**:
1. Add a `healthcheck` to the `postgres` service using `pg_isready`.
2. Add `healthcheck` blocks to the backend services using their `/health` endpoints.
3. Use `depends_on` with `condition: service_healthy` for all backend services depending on `postgres`.
**Validation**:
- [ ] Services start in the correct order (DB $\rightarrow$ Apps).

### Subtask T008: Configure Nginx API Gateway (`nginx.conf`)
**Purpose**: Provide a single point of entry for the frontend.
**Steps**:
1. Create an `nginx.conf` file.
2. Configure a server listening on port 80.
3. Create location blocks to proxy requests:
   - `/api/identity` $\rightarrow$ `http://identity-service:80`
   - `/api/events` $\rightarrow$ `http://events-service:80`
   - `/api/registrations` $\rightarrow$ `http://registrations-service:80`
4. Set appropriate proxy headers (`Host`, `X-Real-IP`).
**Validation**:
- [ ] Nginx config is syntactically correct.

### Subtask T009: Add Nginx Gateway to `docker-compose.yml` and verify connectivity
**Purpose**: Expose the API to the host machine.
**Steps**:
1. Add the `gateway` service to `docker-compose.yml` using the Nginx image and mounting `nginx.conf`.
2. Map host port `8080` to container port `80`.
3. Run `docker-compose up -d`.
4. Test connectivity using `curl http://localhost:8080/api/identity/health` (or similar).
**Validation**:
- [ ] Requests to `localhost:8080` are correctly routed to internal services.

## Definition of Done
- [ ] A single `docker-compose up -d` command starts the entire system.
- [ ] All services are healthy.
- [ ] Frontend can reach all backend services via `localhost:8080`.

## Risks & Mitigations
- **Risk**: Nginx config errors can bring down the entire API.
- **Mitigation**: Use a simple, tested proxy configuration and verify with `nginx -t` during build.

## Activity Log

- 2026-06-08T13:18:23Z – gemma-4-31b-it – shell_pid=3392 – lane=doing – Assigned agent via workflow command
- 2026-06-08T13:19:32Z – gemma-4-31b-it – shell_pid=3392 – lane=for_review – Ready for review: Implemented Docker Compose orchestration with service health checks and Nginx API Gateway mapping host port 8080.
- 2026-06-08T13:31:18Z – gemma-4-31b-it – shell_pid=3792 – lane=doing – Started review via workflow command
- 2026-06-08T13:31:44Z – gemma-4-31b-it – shell_pid=3792 – lane=done – Review passed: Implemented Docker Compose orchestration with service health checks and Nginx API Gateway mapping host port 8080.
