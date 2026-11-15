---
work_package_id: WP02
title: User Registration & Login Services
dependencies:
- WP01
requirement_refs:
- FR-002
planning_base_branch: feature/sdd-plus-project1
merge_target_branch: feature/sdd-plus-project1
branch_strategy: Planning artifacts for this feature were generated on feature/sdd-plus-project1. During /spec-kitty.implement this WP may branch from a dependency-specific base, but completed changes must merge back into feature/sdd-plus-project1 unless the human explicitly redirects the landing branch.
subtasks:
- T004
- T005
- T006
- T007
history: []
authoritative_surface: internal/domain/user/
execution_mode: code_change
owned_files:
- internal/domain/user/**
- internal/api/handlers/user_handler.go
tags: []
agent: "opencode"
shell_pid: "12336"
---

# WP02: User Registration & Login Services

## Objective
Implement secure user registration with role checks, password hashing, and login credentials authentication in the domain service and HTTP handlers.

---

## Subtask Breakdown

### Subtask T004: Implement register domain service with security & uniqueness rules
* **Purpose**: Validate user fields, hash the password safely, and insert records into the database.
* **Guidance**:
  - Open `internal/domain/user/service.go`.
  - Validate parameters in `Register`:
    - Username: Minimum 3 characters, alphanumeric.
    - Password: Minimum 6 characters.
    - Role: Must be one of `attendee`, `organiser`, or `administrator`.
  - Use `golang.org/x/crypto/bcrypt` (or standard hashing package verified in imports) to hash the password securely before saving.
  - Return `ErrUsernameTaken` or equivalent domain error if the database insert fails due to unique constraints.
* **Validation**:
  - Registration inputs are parsed and validated correctly before database queries.

### Subtask T005: Implement login validation and authentication logic
* **Purpose**: Verify credentials against the stored hash and return user metadata.
* **Guidance**:
  - Implement the `Login(username, password)` method in `internal/domain/user/service.go`.
  - Query the user by username. Compare the input password against the stored `password_hash` using bcrypt constant-time comparison (`bcrypt.CompareHashAndPassword`).
  - Return a structured `User` entity or domain auth error on mismatch.
* **Validation**:
  - Secure credential comparison prevents timing and brute force leakages.

### Subtask T006: Implement and correct Register and Login handlers with error maps
* **Purpose**: Capture JSON requests and map domain errors to specific HTTP responses.
* **Guidance**:
  - Update `internal/api/handlers/user_handler.go`.
  - For `Register`: Map duplicate username error to `409 Conflict`, input validation failure to `400 Bad Request`. On success, return `201 Created` with the JSON-serialized user (raw password omitted).
  - For `Login`: On success, call the JWT helper to generate a token and return `200 OK` with `{ "token": "<JWT>" }`. Return `411 Unauthorized` or `401 Unauthorized` on credential mismatch.
* **Validation**:
  - Responses conform exactly to the OpenAPI contracts.

### Subtask T007: Write domain service tests for User model and authentication
* **Purpose**: Achieve high test coverage checking validations and error handling.
* **Guidance**:
  - Write test cases in `internal/domain/user/service_test.go` covering:
    - Successful registration with hashed password confirmation.
    - Registration failures (short username, invalid role).
    - Authentication successes and failure.
  - Ensure test coverage for the user domain is ≥70%.
* **Validation**:
  - Run `go test ./internal/domain/user/... -v` and verify pass.

---

## Branch Strategy & Lane Execution
* Run execution using: `spec-kitty agent action implement WP02 --agent <name>`
* Workspace and worktree are allocated per lane under `.worktrees/`.
* Code changes are committed back to `feature/sdd-plus-project1`.

## Activity Log

- 2026-05-22T21:44:05Z – opencode – shell_pid=7080 – Started implementation via action command
- 2026-05-22T21:45:36Z – opencode – shell_pid=7080 – Ready for review: User registration and login services have been fully implemented with domain validations and status-mapped endpoints.
- 2026-05-22T21:56:13Z – opencode – shell_pid=12336 – Started review via action command
- 2026-05-22T21:56:22Z – opencode – shell_pid=12336 – Review passed: User registration and login services correctly implemented.
