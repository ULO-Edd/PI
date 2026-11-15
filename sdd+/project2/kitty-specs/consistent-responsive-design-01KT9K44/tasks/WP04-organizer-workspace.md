---
work_package_id: WP04
title: Organizer Workspace
dependencies:
- WP01
requirement_refs:
- FR-001
- FR-002
- FR-003
- FR-005
- FR-006
planning_base_branch: feature/sdd-plus-project2
merge_target_branch: feature/sdd-plus-project2
branch_strategy: Planning artifacts for this feature were generated on feature/sdd-plus-project2. During /spec-kitty.implement this WP may branch from a dependency-specific base, but completed changes must merge back into feature/sdd-plus-project2 unless the human explicitly redirects the landing branch.
subtasks:
- T011
- T012
- T013
history: []
authoritative_surface: frontend/src/
execution_mode: code_change
owned_files:
- frontend/src/components/organizer/AccessibleTable.vue
- frontend/src/views/organizer/OrganizerDashboardView.vue
- frontend/src/views/organizer/StatsView.vue
- frontend/src/views/organizer/EventAttendeesView.vue
tags: []
---

# WP04: Organizer Workspace

## Objective
Build adaptive grids, responsive charts representation, metrics boxes, and tables for the Organizer dashboard and views.

## Context
Organizers need to track metrics, inspect attendees, and manage their events. Tabular data and stats dashboards represent unique challenges on mobile screen aspect ratios.

## Detailed Guidance

### Subtask T011: Implement responsive table layouts in AccessibleTable.vue and OrganizerDashboardView.vue
- **Objective**: Optimize data-dense tables for screens with smaller width, without losing accessible features.
- **Detailed Steps**:
  1. Wrap `AccessibleTable.vue` in a responsive overflow wrapper (`overflow-x-auto`).
  2. Implement responsive column structures, toggling less vital columns on smaller screen thresholds (`hidden md:table-cell`).
  3. Ensure active action buttons (edit, publish, attendees list) match appropriate padding and responsive tap targets.
- **Validation**:
  - Resizing tables shows scroll indicators if columns exceed available layout width.
  - Interactive rows are touch accessible.

### Subtask T012: Ensure StatsView.vue charts and metric grids adapt to smaller viewports
- **Objective**: Arrange charts and high-level stats cards grid to flow gracefully.
- **Detailed Steps**:
  1. Build stats-counter widgets using flex grids (`grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4`).
  2. Style chart wrappers with absolute parent sizes or viewport aspect rules, ensuring SVGs/charts automatically adjust width.
- **Validation**:
  - Counter metrics box grid wraps cleanly at small widths.

### Subtask T013: Redesign EventAttendeesView.vue for clean display on tablet and mobile screens
- **Objective**: Provide structured details for attendee lists on mobile screens.
- **Detailed Steps**:
  1. Simplify list representations of attendees and registration validation states for screens below `sm` size.
- **Validation**:
  - Verify attendance lists read beautifully on touch screens.

## Activity Log

- 2026-06-04T16:39:19Z – unknown – Moved to for_review
