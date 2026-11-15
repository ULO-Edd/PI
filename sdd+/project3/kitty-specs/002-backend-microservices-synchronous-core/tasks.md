# Work Packages: Backend Microservices Synchronous Core

**Inputs**: Design documents from `kitty-specs/002-backend-microservices-synchronous-core/`
**Prerequisites**: plan.md (required), spec.md (user stories), research.md, data-model.md, contracts/, quickstart.md

**Organization**: Fine-grained subtasks (`Txxx`) roll up into work packages (`WPxx`). Each work package must be independently deliverable and testable.

---

## Work Package WP01: Perimeter API Gateway and Infrastructure (Priority: P0)

**Goal**: Setup solution scaffolding, API Gateway project, JwtBearer validation, and perimeter rate-limiting.
**Independent Test**: The Gateway project compiles, runs, intercept requests, validates security credentials, and enforces request quotas.
**Prompt**: `/tasks/WP01-perimeter-api-gateway-and-infrastructure.md`

**Requirements Refs**: FR-001, FR-002, FR-003

### Included Subtasks
- [x] T001 Scaffold solution structure, ApiGateway, Services projects, and Shared contracts layout
- [x] T002 Configure path routing and proxy forwarding in `ApiGateway`
- [x] T003 Implement JWT token signature validation guards in gateway middleware
- [x] T004 Implement role-based rate-limiting (quotas for Admin, Organiser, and Attendees)

### Implementation Notes
- Initialize C# .NET 10 solution and create standard directory structures.
- Use native ASP.NET Core rate-limiting middleware to configure quota policies (e.g. 100 req/min vs 1000 req/min).

### Parallel Opportunities
- Gateways routing (T002) and authentication checks (T003) can proceed concurrently once scaffolding is complete.

### Dependencies
- None (starting package).

### Risks & Mitigations
- Inconsistent port mapping; pin local gateway and microservice ports explicitly in `launchSettings.json`.

---

## Work Package WP02: Identity Service and Role Contracts (Priority: P1) 🎯 MVP

**Goal**: Setup Identity Service with EF Core, secure user registrations, and claims token generation.
**Independent Test**: API consumers can register, authenticate via `/api/auth/login`, and receive valid JWT credentials.
**Prompt**: `/tasks/WP02-identity-service-and-role-contracts.md`

**Requirements Refs**: FR-004

### Included Subtasks
- [x] T005 Scaffold IdentityService, set up relational Database Context, and apply initial database migrations
- [x] T006 Implement credentials registration with cryptographic password hashing (bcrypt)
- [x] T007 Build the `/api/auth/login` endpoint returning role-based JWT claims

### Implementation Notes
- Use EF Core 10 migrations. Use BCrypt.Net for password hashing.
- Include Roles (Admin, Organiser, Attendee) in JWT claims payload.

### Parallel Opportunities
- Credentials hashing (T006) can be implemented concurrently with DB context setup (T005) before migrations are applied.

### Dependencies
- Depends on WP01.

### Risks & Mitigations
- Compromised token signatures; ensure signature keys are loaded securely from configuration secrets at container runtime.

---

## Work Package WP03: Events Catalogue & CRUD Services (Priority: P1)

**Goal**: Deliver Events Service with public catalogue queries and secure organiser event management.
**Independent Test**: Public event lists are queries, and organisers can create/update events with strict date and title validations.
**Prompt**: `/tasks/WP03-events-catalogue-and-crud-services.md`

**Requirements Refs**: FR-005

### Included Subtasks
- [x] T008 Scaffold EventsService, set up EF Core DBContext, and apply initial database migrations
- [x] T009 Implement responsive, public event query endpoint `/api/events`
- [x] T010 Build secured, role-restricted event creation and update CRUD endpoints

### Implementation Notes
- Titles must be validated to be <= 100 characters, and dates must be validated to be in the future before saving.
- Secure these endpoints with authorisation filters verifying the `Organiser` claim.

### Parallel Opportunities
- Public queries (T009) and secured CRUD forms (T010) can proceed in parallel once DB schema exists.

### Dependencies
- Depends on WP01, WP02.

### Risks & Mitigations
- Unauthorized event mutations; apply global API Authorisation Filters to ensure only Admin or Organiser roles bypass checks.

---

## Work Package WP04: Registrations Service & Inter-Service Client (Priority: P1)

**Goal**: Deliver Registrations database context and synchronous capacity reservation HTTP client wraps.
**Independent Test**: Ticket booking requests trigger inter-service HTTP capacity deduction checks in the Events database successfully.
**Prompt**: `/tasks/WP04-registrations-service-and-inter-service-client.md`

**Requirements Refs**: FR-006, FR-007, FR-008

### Included Subtasks
- [x] T011 Scaffold RegistrationsService, set up EF Core DBContext, and apply initial database migrations
- [x] T012 Implement atomic capacity deduction endpoint `/api/events/{id}/reserve` in EventsService
- [x] T013 Implement secure HTTP inter-service client wrapper in RegistrationsService calling EventsService

### Implementation Notes
- No direct queries or cross-database joins are allowed between the Events and Registrations databases.
- Ensure the Events reservation endpoint `/reserve` executes atomic counter updates to avoid concurrency races.

### Parallel Opportunities
- Events reservation API (T012) and Registrations DB setup (T011) can run in parallel.

### Dependencies
- Depends on WP01, WP03.

### Risks & Mitigations
- Concurrency race conditions on the last available ticket; implement EF Core optimistic concurrency control or atomic SQL counter update statements.

---

## Work Package WP05: Synchronous Capacity Booking & Rollbacks (Priority: P1)

**Goal**: Complete transactional booking, automatic inventory rollback endpoints, and integration tests.
**Independent Test**: System books tickets under high-load concurrency, and automatically rolls back capacity on database failures.
**Prompt**: `/tasks/WP05-synchronous-capacity-booking-and-rollbacks.md`

**Requirements Refs**: FR-009, FR-010

### Included Subtasks
- [x] T014 Build booking coordinator transaction `/api/registrations/book` in RegistrationsService
- [x] T015 Build capacity rollback endpoint `/api/events/{id}/release` in EventsService
- [x] T016 Write xUnit unit tests verifying controller roles and database sync transactions
- [x] T017 Write Playwright .NET E2E integration test suite for concurrency, rollbacks, and rate-limiting

### Implementation Notes
- If Registrations DB fails to save, immediately post to `/api/events/{id}/release` to restore ticket inventory.
- Write E2E test verifying: mock registration database timeout causes Events inventory to be fully rolled back.

### Parallel Opportunities
- Unit testing (T016) and E2E resilience checks (T017) can be prepared concurrently.

### Dependencies
- Depends on WP04.

### Risks & Mitigations
- Network dropouts midway through rollback requests; implement exponential backoff retries with structured correlation IDs.

---

## Dependency & Execution Summary

- **Sequence**: WP01 -> WP02 (Identity MVP) -> WP03 (Events catalogue) -> WP04 (Registrations HTTP) -> WP05 (Booking transactions & rollbacks).
- **Parallelization**: WP02 and WP03 can proceed concurrently once the gateway is configured.
- **MVP Scope**: WP01 (gateway), WP02 (identity/auth), WP03 (catalogue queries), and WP04 (inter-service communication).

---

## Subtask Index (Reference)

| Subtask ID | Summary | Work Package | Priority | Parallel? |
|------------|---------|--------------|----------|-----------|
| T001       | Scaffold solutions and Shared layers | WP01         | P0       | No        |
| T002       | Configure API Gateway routing | WP01         | P0       | Yes       |
| T003       | Implement JWT validation guards | WP01         | P0       | Yes       |
| T004       | Implement role-based rate-limits | WP01         | P0       | Yes       |
| T005       | Scaffold Identity DB context & migrations | WP02         | P1       | No        |
| T006       | Implement credentials hashing | WP02         | P1       | Yes       |
| T007       | Build login claims endpoint | WP02         | P1       | Yes       |
| T008       | Scaffold Events DB context & migrations | WP03         | P1       | No        |
| T009       | Implement public catalogue queries | WP03         | P1       | Yes       |
| T010       | Build organiser event CRUD API | WP03         | P1       | Yes       |
| T011       | Scaffold Registrations DB & migrations | WP04         | P1       | Yes       |
| T012       | Build Events capacity reserve API | WP04         | P1       | Yes       |
| T013       | Build inter-service HTTP sync client | WP04         | P1       | Yes       |
| T014       | Build Registrations book coordinator API | WP05         | P1       | No        |
| T015       | Build Events capacity release API | WP05         | P1       | No        |
| T016       | Write xUnit unit tests | WP05         | P1       | Yes       |
| T017       | Write Playwright E2E rollback checks | WP05         | P1       | No        |
