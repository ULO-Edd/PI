---
work_package_id: WP03
title: Event CRUD Domain & Handler Implementation
dependencies:
- WP02
requirement_refs:
- FR-003
- FR-004
planning_base_branch: feature/sdd-plus-project1
merge_target_branch: feature/sdd-plus-project1
branch_strategy: Planning artifacts for this feature were generated on feature/sdd-plus-project1. During /spec-kitty.implement this WP may branch from a dependency-specific base, but completed changes must merge back into feature/sdd-plus-project1 unless the human explicitly redirects the landing branch.
subtasks:
- T008
- T009
- T010
- T011
- T012
- T013
history: []
authoritative_surface: internal/domain/event/
execution_mode: code_change
owned_files:
- internal/domain/event/**
- internal/api/handlers/event_handler.go
- internal/api/router.go
tags: []
agent: "opencode"
shell_pid: "19840"
---

# WP03: Event CRUD Domain & Handler Implementation

## Objective
Implement fully functional event CRUD operations with database-backed repository queries, search/filtering support, and organizer-only write permission checks.

---

## Subtask Breakdown

### Subtask T008: Implement Event CRUD domain business service validations
* **Purpose**: Validate event structures inside `internal/domain/event/service.go`.
* **Guidance**:
  - Enforce validations on creation and updates:
    - Title must be non-empty and less than 255 characters.
    - Start time must be in the future (relative to create/update time if applicable, or logically valid).
    - End time must be strictly after Start time.
    - Capacity must be > 0.
    - Status must be a valid state: `draft`, `published`, or `cancelled`.
* **Validation**:
  - Ensure improper payloads are blocked early in the service layer.

### Subtask T009: Add database repository queries supporting full event CRUD
* **Purpose**: Write the SQL database queries in `internal/domain/event/repository.go` for all CRUD functions.
* **Guidance**:
  - Implement SQL statements inside the database repository:
    - Create: `INSERT INTO events ...`
    - GetByID: `SELECT ... FROM events WHERE id = $1`
    - List/GetAll: support querying without filtering, returning a slice of events.
    - Update: `UPDATE events SET title=$1, description=$2, location=$3, start_time=$4, end_time=$5, capacity=$6, status=$7 WHERE id=$8`
    - Delete: `DELETE FROM events WHERE id=$1`
* **Validation**:
  - Ensure rows are read and mapped cleanly into domain structures.

### Subtask T010: Add search and query filtering support for Listing Events handler
* **Purpose**: Implement optional filtering logic inside the list handler.
* **Guidance**:
  - Update `GetEvents` inside `internal/api/handlers/event_handler.go`.
  - Parse optional query parameters: `status`, `location`, `organiser_id`.
  - Pass these filters to the repository layer (or build a dynamic SQL query safely) to filter events list accordingly.
* **Validation**:
  - Performing `GET /events?status=published` returns only published events.

### Subtask T011: Implement event handlers verifying organizer identities
* **Purpose**: Restrict event editing and deletion.
* **Guidance**:
  - Update `CreateEvent`, `UpdateEvent`, and `DeleteEvent` in `internal/api/handlers/event_handler.go`.
  - Extract the authenticated user's ID and role from the JWT context (using JWT claims middleware).
  - For Create: Assign `organiser_id` dynamically from the logged-in user's claims instead of generating a placeholder UUID.
  - For Update/Delete: Verify that the logged-in user is either the original organizer or an Administrator. Return `403 Forbidden` if unauthorized.
* **Validation**:
  - Unauthorized edits on other organizers' events are correctly blocked with `403 Forbidden`.

### Subtask T012: Wire Event CRUD routes in the API Router
* **Purpose**: Ensure router multiplexes correctly to all CRUD methods.
* **Guidance**:
  - Open `internal/api/router.go`.
  - Wire all new endpoints:
    - `GET /events` (Public)
    - `GET /events/{eventId}` (Public)
    - `POST /events` (Protected - Organiser/Admin roles)
    - `PUT /events/{eventId}` (Protected - Organiser/Admin roles)
    - `DELETE /events/{eventId}` (Protected - Organiser/Admin roles)
* **Validation**:
  - Endpoint routing routes correctly, enforcing middleware security rules.

### Subtask T013: Add domain service unit tests for event model and validations
* **Purpose**: Verify event services thoroughly.
* **Guidance**:
  - Implement robust unit tests in `internal/domain/event/service_test.go` covering:
    - Happy path event creation and CRUD updates.
    - Creation/update failures (end_time before start_time, capacity non-positive).
    - Search filtering logic.
  - Achieve ≥70% coverage.
* **Validation**:
  - Run `go test ./internal/domain/event/...` to confirm success.

---

## Branch Strategy & Lane Execution
* Run execution using: `spec-kitty agent action implement WP03 --agent <name>`
* Workspace and worktree are allocated per lane under `.worktrees/`.
* Code changes are committed back to `feature/sdd-plus-project1`.

## Activity Log

- 2026-05-22T21:46:54Z – opencode – shell_pid=3732 – Started implementation via action command
- 2026-05-22T21:51:40Z – opencode – shell_pid=3732 – Ready for review: Event CRUD endpoints have been fully implemented with query filtering, validation constraints, and role-based ownership checks.
- 2026-05-22T21:56:30Z – opencode – shell_pid=19840 – Started review via action command
- 2026-05-22T21:56:42Z – opencode – shell_pid=19840 – Review passed: Event CRUD operations and filtering correctly implemented.
