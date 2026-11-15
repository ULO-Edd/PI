# Implementation Plan: Fix Event and User Registration Core

**Mission ID**: `01KS8SHC2ETY63EWAVDBNA8PP0`
**Mission Type**: `software-dev`

This document details the step-by-step implementation plan to fix core database migrations and implement complete user and event registration features with domain-level validation.

---

## 1. Technical Context

### Existing System & Limitations
- The PostgreSQL database is configured to run in Docker, but the existing up-migration script contains both creation and destruction statements. This causes migrations to drop tables immediately after creating them in the same transaction.
- The user registration (`POST /register`), login, and JWT generation logic are partially implemented but fail due to the missing `users` table and stubbed database behavior.
- Event management CRUD endpoints exist in code but lack full database queries, organiser association, and filter capabilities.
- Registration endpoints are stubbed and do not enforce event capacity limit constraints or correct authentication mapping (e.g. attendee ID is hardcoded).

---

## 2. Charter Check

- **Testing Standard**: Minimum 70% unit test coverage using Go's built-in `testing` library. Every handler and domain service must cover happy paths and error/edge cases.
- **Quality Gates**: Maintain code quality conforming to `golangci-lint` standards.
- **Deployment**: Must be fully deployable via Docker and Docker Compose.
- **Branch Strategy**: Development and merging target is `feature/sdd-plus-project1`.

---

## 3. Work Packages

### WP01: Database Schema Migration Correction
* **Goal**: Ensure tables are correctly and durably created in the database.
* **Tasks**:
  1. Modify `db/migrations/000001_initial_schema.up.sql` to remove the `-- +migrate Down` block and the `DROP TABLE` statements.
  2. Modify `db/migrations/000001_initial_schema.down.sql` to contain the correct `DROP TABLE` statements (`DROP TABLE IF EXISTS registrations; DROP TABLE IF EXISTS events; DROP TABLE IF EXISTS users;`).
  3. Verify that running migrations creates the schemas and that tables persist upon container restart.

### WP02: Domain Services & Validations
* **Goal**: Implement business rules and validation constraints at the Go domain service layer.
* **Tasks**:
  1. Implement user registration and password hashing in `internal/domain/user/service.go`. Add checks for unique username and valid role inputs.
  2. Implement event CRUD business logic in `internal/domain/event/service.go` ensuring valid status transitions and negative capacity constraints are blocked.
  3. Implement event registration/unregistration validations in `internal/domain/registration/service.go`. Enforce the capacity limits and block duplicates.
  4. Write comprehensive unit tests in `service_test.go` for all three domains, aiming for ≥70% coverage.

### WP03: API Endpoints & Handlers
* **Goal**: Complete handlers and route wiring with correct error status mappings.
* **Tasks**:
  1. Complete user registration (`POST /register`) and login (`POST /login`) handlers in `internal/api/handlers/user_handler.go`. Map domain conflicts (e.g., duplicate username) to `409 Conflict`.
  2. Implement search and filter parameters (`status`, `location`, `organiser_id`) in `GET /events` handler.
  3. Ensure event write operations (`POST`, `PUT`, `DELETE` events) extract the authenticated user's ID from the JWT claims to assign ownership and verify permissions.
  4. Implement registration (`POST /events/{eventId}/register`) and unregistration (`DELETE /events/{eventId}/register`) handlers, extracting attendee ID from authenticated JWT.
  5. Route endpoints correctly through auth middleware in `internal/api/router.go`.

### WP04: Integration Testing & Verification
* **Goal**: Validate end-to-end functionality and achieve quality gate standards.
* **Tasks**:
  1. Write API integration tests covering registration, login, event creation, search filtering, and capacity validation.
  2. Run the test suite and verify coverage stats.
  3. Build and test inside Docker containers to ensure clean environment isolation.

---

## 4. Verification Plan

### Test Commands
```bash
# Run all Go unit and integration tests
go test ./... -coverprofile=coverage.out
go tool cover -func=coverage.out
```

### Manual Verification Flow
1. Boot the environment via `docker compose up --build`.
2. Inspect the startup logs for `Migrations ran successfully`.
3. Use the Swagger API UI (`http://localhost:8080/swagger/index.html`) to register a new user, log in, create a published event, and register attendees up to capacity. Ensure subsequent registrations fail with `409 Conflict`.
