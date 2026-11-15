---
work_package_id: WP04
title: Event Registration & Capacity Gates
dependencies:
- WP03
requirement_refs:
- FR-005
- FR-006
- FR-007
planning_base_branch: feature/sdd-plus-project1
merge_target_branch: feature/sdd-plus-project1
branch_strategy: Planning artifacts for this feature were generated on feature/sdd-plus-project1. During /spec-kitty.implement this WP may branch from a dependency-specific base, but completed changes must merge back into feature/sdd-plus-project1 unless the human explicitly redirects the landing branch.
subtasks:
- T014
- T015
- T016
- T017
- T018
history: []
authoritative_surface: internal/domain/registration/
execution_mode: code_change
owned_files:
- internal/domain/registration/**
- internal/api/handlers/registration_handler.go
tags: []
agent: "opencode"
shell_pid: "1224"
---

# WP04: Event Registration & Capacity Gates

## Objective
Implement event registration and unregistration endpoints with strict validation checks ensuring registrations only occur on published events, block duplicates, and respect maximum event capacity.

---

## Subtask Breakdown

### Subtask T014: Implement registration and unregistration logic in domain service
* **Purpose**: Code register and unregister logic in `internal/domain/registration/service.go` and `repository.go`.
* **Guidance**:
  - Implement service and repository logic to add/remove registration entries in the `registrations` database table.
  - Generate a secure new UUID for each registration record on insertion.
* **Validation**:
  - Insert query successfully stores attendee mapping, and delete removes it.

### Subtask T015: Add domain validators for Published status, duplicates, and capacity
* **Purpose**: Integrate business logic validation checking.
* **Guidance**:
  - In `RegisterForEvent(eventID, attendeeID)`:
    - Load the event. Validate that its status is `published`. If not (e.g. `draft` or `cancelled`), return a validation error.
    - Check if a registration already exists for the same `eventID` and `attendeeID`. If so, return a duplicate error.
    - Query the total count of registrations for the event. If `currentCount >= event.Capacity`, block the registration and return `ErrEventFull`.
* **Validation**:
  - Ensure validations prevent registration violations.

### Subtask T016: Complete Registration endpoints extracting Attendee identity from claims
* **Purpose**: Map HTTP payload, extract authenticated attendee ID, and convert domain errors.
* **Guidance**:
  - Update `internal/api/handlers/registration_handler.go`.
  - Extract the logged-in attendee's user ID from JWT context (rather than using placeholder dummy UUID).
  - Call the registration domain service.
  - Map custom errors:
    - Return `400 Bad Request` if event status is not published.
    - Return `409 Conflict` if duplicate registration is attempted, or if `ErrEventFull` is returned.
* **Validation**:
  - Response codes and payloads conform to spec rules.

### Subtask T017: Connect registration endpoints to Router with authentication
* **Purpose**: Secure endpoints using the existing authentication middleware.
* **Guidance**:
  - Review routes configuration in `internal/api/router.go`.
  - Ensure `POST /events/{eventId}/register` and `DELETE /events/{eventId}/register` are wired with JWTAuth and require the role `attendee` (or `administrator`).
* **Validation**:
  - Sending unauthenticated requests to registration paths fails with `401 Unauthorized` or `403 Forbidden`.

### Subtask T018: Write unit tests verifying registration capacity gates
* **Purpose**: Create thorough test suites for registration limits.
* **Guidance**:
  - Write unit/service test cases in `internal/domain/registration/service_test.go` covering:
    - Successful registration for available capacity.
    - Blocking registration when event is full (asserting `ErrEventFull`).
    - Blocking duplicate registration.
    - Unregistering attendee.
  - Ensure coverage is ≥70%.
* **Validation**:
  - All registration tests pass successfully.

---

## Branch Strategy & Lane Execution
* Run execution using: `spec-kitty agent action implement WP04 --agent <name>`
* Workspace and worktree are allocated per lane under `.worktrees/`.
* Code changes are committed back to `feature/sdd-plus-project1`.

## Activity Log

- 2026-05-22T21:52:15Z – opencode – shell_pid=6224 – Started implementation via action command
- 2026-05-22T21:54:15Z – opencode – shell_pid=6224 – Ready for review: Event registration endpoints have been fully implemented with capacity and duplicate validation.
- 2026-05-22T21:56:51Z – opencode – shell_pid=1224 – Started review via action command
- 2026-05-22T21:57:00Z – opencode – shell_pid=1224 – Review passed: Event registration logic with capacity and status gates correctly implemented.
