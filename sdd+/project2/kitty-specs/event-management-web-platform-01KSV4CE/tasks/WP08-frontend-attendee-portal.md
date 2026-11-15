---
work_package_id: WP08
title: Frontend Attendee Event Portal
dependencies:
- WP07
requirement_refs:
- FR-004
- FR-005
- FR-006
- NFR-001
- NFR-003
planning_base_branch: feature/sdd-plus-project2
merge_target_branch: feature/sdd-plus-project2
branch_strategy: Planning artifacts for this feature were generated on feature/sdd-plus-project2. During /spec-kitty.implement this WP may branch from a dependency-specific base, but completed changes must merge back into feature/sdd-plus-project2 unless the human explicitly redirects the landing branch.
subtasks:
- T034
- T035
- T036
- T037
phase: Phase 8 - Attendee Layout & Discover Portal
assignee: ''
agent: "opencode"
shell_pid: "9460"
history:
- timestamp: '2026-05-30T00:31:40Z'
  agent: system
  action: Prompt generated via /spec-kitty.tasks
authoritative_surface: frontend/src/views/attendee/
execution_mode: code_change
owned_files:
- frontend/src/views/attendee/**
- frontend/src/components/attendee/**
tags: []
---

# Work Package Prompt: WP08 – Frontend Attendee Event Portal

## Objective
Design accessible search filters, detail grids, seat counter cards, registration triggers, and personal timeline dashboards.

---

## Subtask Details

### Subtask T034: Build Searchable Events Discovery Grid
*   **Purpose:** Render event lists categorized and filterable.
*   **Steps:**
    1. Develop accessible cards showing title, dates, organizer info, and category.
*   **Files:** `frontend/src/views/attendee/DiscoveryView.vue`, `frontend/src/components/attendee/EventCard.vue`
*   **Validation:** Search filter returns exact matches on active queries.

### Subtask T035: Develop Event Info Detail View
*   **Purpose:** Display event details and seat remaining counters.
*   **Steps:**
    1. Design page capturing rich event metadata and capacity.
*   **Files:** `frontend/src/views/attendee/EventDetailView.vue`
*   **Validation:** Disable registration triggers when events hit capacity.

### Subtask T036: Attendee Personal Dashboard
*   **Purpose:** Track active and past registered events.
*   **Steps:**
    1. Design dashboard layout filtering user sign-ups.
*   **Files:** `frontend/src/views/attendee/DashboardView.vue`
*   **Validation:** Timeline shows correct historical listing.

### Subtask T037: Visual Registration status notifications
*   **Purpose:** Confirm registration triggers visually.
*   **Steps:**
    1. Add accessible status banners confirming registration.
*   **Files:** `frontend/src/components/attendee/StatusBanner.vue`
*   **Validation:** Flash notification pops up and is readable by screen-readers.

## Activity Log

- 2026-05-31T10:32:42Z – opencode – shell_pid=12812 – Started implementation via action command
- 2026-05-31T10:33:58Z – opencode – shell_pid=12812 – Attendee event discovery grid with filters, detailed info views, and registration tracking implemented and verified.
- 2026-05-31T10:34:08Z – opencode – shell_pid=9460 – Started review via action command
- 2026-05-31T10:34:17Z – opencode – shell_pid=9460 – Review passed: Searchable discovery grid, detail views with registration logic, and tracking dashboard implemented and verified with production builds.
