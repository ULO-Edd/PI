---
work_package_id: WP09
title: Frontend Organizer Dashboard
dependencies:
- WP08
requirement_refs:
- FR-007
- FR-008
- FR-010
- FR-011
- FR-012
- NFR-001
- NFR-002
- NFR-003
- NFR-004
planning_base_branch: feature/sdd-plus-project2
merge_target_branch: feature/sdd-plus-project2
branch_strategy: Planning artifacts for this feature were generated on feature/sdd-plus-project2. During /spec-kitty.implement this WP may branch from a dependency-specific base, but completed changes must merge back into feature/sdd-plus-project2 unless the human explicitly redirects the landing branch.
subtasks:
- T038
- T039
- T040
- T041
phase: Phase 9 - Organizer Dashboard controls
assignee: ''
agent: "opencode"
shell_pid: "1788"
history:
- timestamp: '2026-05-30T00:31:40Z'
  agent: system
  action: Prompt generated via /spec-kitty.tasks
authoritative_surface: frontend/src/views/organizer/
execution_mode: code_change
owned_files:
- frontend/src/views/organizer/**
- frontend/src/components/organizer/**
tags: []
---

# Work Package Prompt: WP09 – Frontend Organizer Dashboard

## Objective
Develop organizer event CRUD panels, roster attendance toggles, stats metrics charts, and exported CSV triggers.

---

## Subtask Details

### Subtask T038: Organizer Event CRUD Forms
*   **Purpose:** Provide simple dashboard to create and publish events.
*   **Steps:**
    1. Implement forms for adding titles, descriptions, categories, capacities, dates, location, and metadata.
*   **Files:** `frontend/src/views/organizer/EventFormView.vue`, `frontend/src/components/organizer/DashboardNavbar.vue`
*   **Validation:** Event displays in lists immediately upon publication.

### Subtask T039: Interactive Attendance Roster
*   **Purpose:** Allow toggling attendance status on registrations.
*   **Steps:**
    1. Build table mapping attendees with present/no-show checkboxes.
*   **Files:** `frontend/src/views/organizer/RosterView.vue`
*   **Validation:** Toggles are instant and sync directly with Go backend.

### Subtask T040: Dashboard Charts & CSV Trigger
*   **Purpose:** Render statistics visual triggers and execute CSV download.
*   **Steps:**
    1. Build analytics layouts for fill rates and registrations trends.
    2. Add CSV download anchor mapping backend exporter.
*   **Files:** `frontend/src/views/organizer/StatsView.vue`
*   **Validation:** Clicking export downloads expected CSV roster.

### Subtask T041: Keyboard Focus on Tables
*   **Purpose:** Align dense stats displays with accessibility focus ring standards.
*   **Steps:**
    1. Map custom column tabs and row headers for screen readers.
*   **Files:** `frontend/src/components/organizer/AccessibleTable.vue`
*   **Validation:** Table rows and buttons support keyboard only selection.

## Activity Log

- 2026-05-31T10:36:56Z – opencode – shell_pid=7632 – Started implementation via action command
- 2026-05-31T21:15:01Z – opencode – shell_pid=7632 – Organizer dashboard with event CRUD, attendee management, CSV export, and accessible tables implemented and verified.
- 2026-05-31T21:15:13Z – opencode – shell_pid=1788 – Started review via action command
- 2026-05-31T21:16:45Z – opencode – shell_pid=1788 – Review passed: Organizer dashboard with event CRUD, attendee management, CSV export, and accessible tables implemented and verified.
