---
work_package_id: WP04
title: Backend Event Lifecycle & Attendance Management
dependencies:
- WP03
requirement_refs:
- FR-007
- FR-008
- FR-010
- FR-011
- FR-012
planning_base_branch: feature/sdd-plus-project2
merge_target_branch: feature/sdd-plus-project2
branch_strategy: Planning artifacts for this feature were generated on feature/sdd-plus-project2. During /spec-kitty.implement this WP may branch from a dependency-specific base, but completed changes must merge back into feature/sdd-plus-project2 unless the human explicitly redirects the landing branch.
subtasks:
- T016
- T017
- T018
- T019
phase: Phase 4 - Event Lifecycle & Roster Control
assignee: ''
agent: "opencode"
shell_pid: "9136"
history:
- timestamp: '2026-05-30T00:31:40Z'
  agent: system
  action: Prompt generated via /spec-kitty.tasks
authoritative_surface: backend/internal/events/
execution_mode: code_change
owned_files:
- backend/internal/events/**
tags: []
---

# Work Package Prompt: WP04 – Backend Event Lifecycle & Attendance Management

## Objective
Enable complete event CRUD, drafting, cancel states, attendee check-ins, and downloadable CSV list generation.

---

## Subtask Details

### Subtask T016: Organizer Event CRUD Logic
*   **Purpose:** Implement robust endpoints for creating, editing, and cancelling events.
*   **Steps:**
    1. Create PUT `/api/events/{id}/cancel` and POST/PUT operations ensuring only Approved organizers can publish.
*   **Files:** `backend/internal/events/event_controller.go`
*   **Validation:** Blocks publication attempts from pending organizers with 403.

### Subtask T017: Attendee Presence & Status Log
*   **Purpose:** Allow organizers to toggle presence/no-show on event registrations.
*   **Steps:**
    1. Write PUT `/api/events/{id}/attendees/{attendee_id}/presence` to save presence states.
*   **Files:** `backend/internal/events/attendance_controller.go`
*   **Validation:** Restricts editing actions to original event organizers.

### Subtask T018: CSV Attendance Roster Export
*   **Purpose:** Export an event registration list to CSV.
*   **Steps:**
    1. Write handler responding with standard content-disposition header and CSV payload.
*   **Files:** `backend/internal/events/csv_exporter.go`
*   **Validation:** Exported CSV includes correct headers (Name, Email, Registration Time, Attendance Status).

### Subtask T019: Backend Organizer Integration Tests
*   **Purpose:** Secure event routing with integration testing.
*   **Steps:**
    1. Implement functional endpoint testing suites for Organizer routes.
*   **Files:** `backend/internal/events/organizer_test.go`
*   **Validation:** Verification tests succeed with no-side effects.

## Activity Log

- 2026-05-30T20:59:21Z – opencode – shell_pid=17100 – Started implementation via action command
- 2026-05-30T21:04:38Z – opencode – shell_pid=17100 – Backend event lifecycle, attendance status, and CSV exports fully implemented and verified with integration tests.
- 2026-05-30T21:04:51Z – opencode – shell_pid=9136 – Started review via action command
- 2026-05-30T21:05:05Z – opencode – shell_pid=9136 – Review passed: Event CRUD, attendee rosters, presence updates, and CSV exports all implemented and verified.
