---
work_package_id: WP03
title: Backend Core Logic & Admin/Organizer Auth Guarding
dependencies:
- WP02
requirement_refs:
- FR-001
- FR-002
- FR-003
- FR-013
planning_base_branch: feature/sdd-plus-project2
merge_target_branch: feature/sdd-plus-project2
branch_strategy: Planning artifacts for this feature were generated on feature/sdd-plus-project2. During /spec-kitty.implement this WP may branch from a dependency-specific base, but completed changes must merge back into feature/sdd-plus-project2 unless the human explicitly redirects the landing branch.
subtasks:
- T011
- T012
- T013
- T014
- T015
phase: Phase 3 - Backend Core & Authentication
assignee: ''
agent: "opencode"
shell_pid: "1684"
history:
- timestamp: '2026-05-30T00:31:40Z'
  agent: system
  action: Prompt generated via /spec-kitty.tasks
authoritative_surface: backend/internal/
execution_mode: code_change
owned_files:
- backend/internal/auth/**
- backend/internal/models/**
tags: []
---

# Work Package Prompt: WP03 – Backend Core Logic & Admin/Organizer Auth Guarding

## Objective
Update User authentication models, parse JWT claims for roles, and build administrative approval interfaces for organizer validation.

---

## Subtask Details

### Subtask T011: Go User Model Role Extensions
*   **Purpose:** Extend the User model in Go with `role` and `status` attributes.
*   **Steps:**
    1. Update model definitions in `backend/internal/models/user.go`.
*   **Files:** `backend/internal/models/user.go`
*   **Validation:** Model fields compile.

### Subtask T012: Role Claims JWT Integration
*   **Purpose:** Include role and status claims inside standard issued JWT and enforce middleware checks.
*   **Steps:**
    1. Update token generation with claims.
    2. Add dynamic checking middleware for Admin, Organizer, and Attendee roles.
*   **Files:** `backend/internal/auth/jwt.go`, `backend/internal/auth/middleware.go`
*   **Validation:** Middleware rejects requests missing correct role claims.

### Subtask T013: Admin Organizer Approval Handlers
*   **Purpose:** Build Administrator endpoint to approve/suspend Organizer accounts.
*   **Steps:**
    1. Write GET `/api/admin/organizers` and PUT `/api/admin/organizers/{id}/status`.
*   **Files:** `backend/internal/auth/admin_controller.go`
*   **Validation:** Handlers correctly save status changes to the DB.

### Subtask T014: Admin Category CRUD
*   **Purpose:** Implement category creation and listing endpoints.
*   **Steps:**
    1. Implement POST `/api/categories` and GET `/api/categories` restricted to Admins.
*   **Files:** `backend/internal/auth/category_controller.go`
*   **Validation:** Categories successfully persist in database.

### Subtask T015: Admin Endpoint Tests
*   **Purpose:** Add backend integration tests for administrative actions.
*   **Steps:**
    1. Implement testing suites inside `backend/internal/auth/admin_test.go`.
*   **Files:** `backend/internal/auth/admin_test.go`
*   **Validation:** Run suites using `go test`.

## Activity Log

- 2026-05-30T20:50:29Z – opencode – shell_pid=12508 – Started implementation via action command
- 2026-05-30T20:57:45Z – opencode – shell_pid=12508 – Backend role mapping, JWT status parsing, and Admin organizer verification endpoints fully implemented and verified with integration tests.
- 2026-05-30T20:57:58Z – opencode – shell_pid=1684 – Started review via action command
- 2026-05-30T20:58:30Z – opencode – shell_pid=1684 – Review passed: User model extended with status, JWT tokens verified with status claims, IsApproved middleware applied, and Administrator organizer management and category CRUD endpoints successfully verified by integration tests.
