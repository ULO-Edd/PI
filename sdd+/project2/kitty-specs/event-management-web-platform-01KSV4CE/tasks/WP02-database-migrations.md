---
work_package_id: WP02
title: Database Extensions & Migration Scripts
dependencies:
- WP01
requirement_refs:
- C-001
planning_base_branch: feature/sdd-plus-project2
merge_target_branch: feature/sdd-plus-project2
branch_strategy: Planning artifacts for this feature were generated on feature/sdd-plus-project2. During /spec-kitty.implement this WP may branch from a dependency-specific base, but completed changes must merge back into feature/sdd-plus-project2 unless the human explicitly redirects the landing branch.
subtasks:
- T006
- T007
- T008
- T009
- T010
phase: Phase 2 - Database Schema Extension
assignee: ''
agent: "opencode"
shell_pid: "13624"
history:
- timestamp: '2026-05-30T00:31:40Z'
  agent: system
  action: Prompt generated via /spec-kitty.tasks
authoritative_surface: backend/db/migrations/
execution_mode: code_change
owned_files:
- backend/db/migrations/**
tags: []
---

# Work Package Prompt: WP02 – Database Extensions & Migration Scripts

## Objective
Apply schema migrations to support Categories, Organizer Approval statuses, Rich Event Metadata, and Registration Presence states.

---

## Subtask Details

### Subtask T006: Create Categories Schema Migrations
*   **Purpose:** Implement `categories` table and seed foundational categories (e.g., Technology, Business, Design, Academic).
*   **Steps:**
    1. Create standard UP and DOWN SQL migration files inside `backend/db/migrations/`.
*   **Files:** `backend/db/migrations/*_create_categories.up.sql`, `backend/db/migrations/*_create_categories.down.sql`
*   **Validation:** Tables exist after migrations complete.

### Subtask T007: Organizer Verification Column Schema
*   **Purpose:** Extend `users` table with status column.
*   **Steps:**
    1. Write migration script adding `status` attribute (Pending, Approved, Suspended) with a check constraint on roles.
*   **Files:** `backend/db/migrations/*_add_organizer_status.up.sql`
*   **Validation:** Column added with defaults.

### Subtask T008: Rich Event Metadata Columns
*   **Purpose:** Add categories and metadata fields to `events` table.
*   **Steps:**
    1. Write migration file linking events to categories, and adding `status`, `contact_email`, and `location`.
*   **Files:** `backend/db/migrations/*_add_event_rich_metadata.up.sql`
*   **Validation:** Fields exist and references match category table.

### Subtask T009: Registration Presence status
*   **Purpose:** Add status attributes to tracking records.
*   **Steps:**
    1. Create migration scripts to extend registrations with attendance state checks (`Registered`, `Present`, `No-Show`, `Cancelled`).
*   **Files:** `backend/db/migrations/*_add_presence_to_registrations.up.sql`
*   **Validation:** Database verifies status inputs correctly.

### Subtask T010: Migrations Verification
*   **Purpose:** Test entire migration sequence up and down.
*   **Steps:**
    1. Execute migration tools against active PostgreSQL database to verify rollback.
*   **Files:** `backend/db/migrations/**`
*   **Validation:** Migrations complete cleanly with no failures.

## Activity Log

- 2026-05-30T00:45:31Z – opencode – shell_pid=7884 – Started implementation via action command
- 2026-05-30T00:46:37Z – opencode – shell_pid=7884 – Database migrations written for categories, organizer verification status, rich event metadata, and registration presence status.
- 2026-05-30T00:46:51Z – opencode – shell_pid=13624 – Started review via action command
- 2026-05-30T00:47:06Z – opencode – shell_pid=13624 – Review passed: Category schema, organizer status, event metadata, and registration presence status columns successfully migrated and verified.
