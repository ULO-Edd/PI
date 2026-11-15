---
work_package_id: WP02
title: Database & Domain Models
dependencies:
- WP01
requirement_refs:
- FR-001
- FR-002
- FR-003
- FR-005
- FR-006
- FR-007
- FR-008
- FR-011
planning_base_branch: feature/sdd-plus-project1
merge_target_branch: feature/sdd-plus-project1
branch_strategy: Planning artifacts for this feature were generated on feature/sdd-plus-project1. During /spec-kitty.implement this WP may branch from a dependency-specific base, but completed changes must merge back into feature/sdd-plus-project1 unless the human explicitly redirects the landing branch.
subtasks:
- T004
- T005
- T006
history: []
authoritative_surface: internal/domain/
execution_mode: code_change
owned_files:
- internal/platform/database/
- internal/domain/user/user.go
- internal/domain/event/event.go
- internal/domain/registration/registration.go
- db/migrations/
tags: []
agent: "opencode"
shell_pid: "7416"
---

## 1. Objective

Establish the database connection, set up a migration system, and define the core domain entities (User, Event, Registration) in Go.

## 2. Detailed Guidance

### Subtask T004: Implement database connection and migration setup
- In `internal/platform/database`, write the code to connect to the PostgreSQL database using the configuration from WP01.
- Choose and implement a database migration library (e.g., `golang-migrate/migrate`). Create an initial migration script.

### Subtask T005: Create database schema for `users`, `events`, and `registrations`
- Create a new SQL migration file.
- Write the `CREATE TABLE` statements for the `users`, `events`, and `registrations` tables as defined in `data-model.md`.

### Subtask T006: Create domain models/entities
- In the `internal/domain/` subpackages (`user`, `event`, `registration`), create the Go structs that represent the database tables.
- These structs should have fields that map to the table columns, with appropriate data types and JSON tags for API responses.

## 3. Definition of Done
- The application can successfully connect to the database.
- The migration system is in place and the initial schema migration can be applied successfully.
- Go structs for User, Event, and Registration are defined.

## Activity Log

- 2026-05-22T13:30:41Z – opencode – shell_pid=7416 – Started implementation via action command
- 2026-05-22T13:31:15Z – opencode – shell_pid=7416 – Ready for review
- 2026-05-22T13:32:52Z – opencode – shell_pid=7416 – Review passed
