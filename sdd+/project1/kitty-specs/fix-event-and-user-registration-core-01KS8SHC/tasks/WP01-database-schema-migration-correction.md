---
work_package_id: WP01
title: Database Schema Migration Correction
dependencies: []
requirement_refs:
- FR-001
planning_base_branch: feature/sdd-plus-project1
merge_target_branch: feature/sdd-plus-project1
branch_strategy: Planning artifacts for this feature were generated on feature/sdd-plus-project1. During /spec-kitty.implement this WP may branch from a dependency-specific base, but completed changes must merge back into feature/sdd-plus-project1 unless the human explicitly redirects the landing branch.
base_branch: kitty/mission-fix-event-and-user-registration-core-01KS8SHC
base_commit: 131444376bd760abd219ea503f52095bbe9d03fd
created_at: '2026-05-22T21:42:03.506454+00:00'
subtasks:
- T001
- T002
- T003
shell_pid: "6440"
history: []
authoritative_surface: db/migrations/
execution_mode: code_change
owned_files:
- db/migrations/**
tags: []
agent: "opencode"
---

# WP01: Database Schema Migration Correction

## Objective
Correct the SQL migrations configuration so that `users`, `events`, and `registrations` tables are permanently and successfully created and persisted when the database runs.

---

## Subtask Breakdown

### Subtask T001: Fix initial up-migration syntax in SQL scripts
* **Purpose**: Remove drop script block from `000001_initial_schema.up.sql`.
* **Guidance**:
  - Open `db/migrations/000001_initial_schema.up.sql`.
  - Remove lines from `-- +migrate Down` onwards. Keep ONLY the three `CREATE TABLE` blocks for `users`, `events`, and `registrations`.
* **Validation**:
  - File contains no `DROP TABLE` statements.

### Subtask T002: Implement robust down-migration syntax for teardown
* **Purpose**: Provide clean destruction statements in the down migration file.
* **Guidance**:
  - Open `db/migrations/000001_initial_schema.down.sql`.
  - Replace the contents with proper `DROP TABLE IF EXISTS` statements in reverse dependency order:
    ```sql
    DROP TABLE IF EXISTS registrations;
    DROP TABLE IF EXISTS events;
    DROP TABLE IF EXISTS users;
    ```
* **Validation**:
  - Ensure the tables are dropped in this exact sequence to avoid foreign key violations.

### Subtask T003: Verify migrations execution against uninitialized database
* **Purpose**: Test that the schema builds correctly from scratch and persists.
* **Guidance**:
  - Run `docker compose down -v` to wipe the persistent postgres volume.
  - Run `docker compose up --build` and watch the `api` container output logs for `Migrations ran successfully`.
  - Connect to the containerized database or verify via API that query actions don't return `relation "users" does not exist`.
* **Validation**:
  - Server boots up cleanly and connects to a valid, persistent schema.

---

## Branch Strategy & Lane Execution
* Run execution using: `spec-kitty agent action implement WP01 --agent <name>`
* Workspace and worktree are allocated per lane under `.worktrees/`.
* Code changes are committed back to `feature/sdd-plus-project1`.

## Activity Log

- 2026-05-22T21:42:06Z – opencode – shell_pid=3528 – Assigned agent via action command
- 2026-05-22T21:42:59Z – opencode – shell_pid=3528 – Ready for review: Database migrations have been split successfully into clean up and down scripts.
- 2026-05-22T21:54:38Z – opencode – shell_pid=6440 – Started review via action command
- 2026-05-22T21:56:05Z – opencode – shell_pid=6440 – Review passed: Migration scripts correctly separated.
