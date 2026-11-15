---
work_package_id: WP03
title: Attendee Portal
dependencies:
- WP01
requirement_refs:
- FR-001
- FR-002
- FR-003
- FR-004
- FR-005
- FR-006
planning_base_branch: feature/sdd-plus-project2
merge_target_branch: feature/sdd-plus-project2
branch_strategy: Planning artifacts for this feature were generated on feature/sdd-plus-project2. During /spec-kitty.implement this WP may branch from a dependency-specific base, but completed changes must merge back into feature/sdd-plus-project2 unless the human explicitly redirects the landing branch.
subtasks:
- T008
- T009
- T010
history: []
authoritative_surface: frontend/src/
execution_mode: code_change
owned_files:
- frontend/src/views/attendee/DiscoveryView.vue
- frontend/src/views/attendee/EventDetailView.vue
- frontend/src/views/attendee/DashboardView.vue
- frontend/src/components/attendee/EventCard.vue
tags: []
---

# WP03: Attendee Portal

## Objective
Bring full responsive flow to the Attendee dashboard, event discovery grids, event detail cards, and registration display tables.

## Context
Attendees browse events on mobile and register. Ensuring event search, filter components, cards, details, and personal dashboards scale cleanly is high priority for WCAG 2.1 AA accessibility and overall experience.

## Detailed Guidance

### Subtask T008: Adapt DiscoveryView.vue (and EventCard.vue) with responsive search/filter layouts and card grid
- **Objective**: Scale card layout grid, filters, and search forms according to screen width.
- **Detailed Steps**:
  1. Redesign discovery controls: wrap search query inputs and category filter dropdown in responsive block-to-row structures (`flex flex-col md:flex-row`).
  2. Implement responsive grid styling on event cards section: use `grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6`.
  3. Ensure individual `EventCard` components adapt cleanly, scaling image ratios and text sizes.
- **Validation**:
  - Verify cards grid transition seamlessly between single-column (mobile), double-column (tablet), and multi-column (desktop).
  - Search fields do not overflow container borders.

### Subtask T009: Make EventDetailView.vue responsive (single column mobile, two columns desktop)
- **Objective**: Refactor event details into a flexible layout that structures details efficiently on smaller screens.
- **Detailed Steps**:
  1. Apply two-column grid on desktop/tablets (`lg:grid-cols-3` with content spanning 2 columns, and info panel spanning 1 column) which collapses to one single column on mobile.
  2. Implement responsive typography, image banners, and action registration buttons.
- **Validation**:
  - Image scales correctly without distortion on all viewports.
  - Form/registration actions panels fit correctly.

### Subtask T010: Ensure attendee DashboardView.vue lists items gracefully on mobile and desktop
- **Objective**: Make attendee registrations, list states, and historical details adapt elegantly.
- **Detailed Steps**:
  1. For tabular event registries, employ mobile card layouts or horizontal overflow tables to prevent page breakage on phones.
  2. Maintain responsive CTA buttons for registration status management.
- **Validation**:
  - Verify no horizontal scrolls exist on the registration list dashboard across standard screens.

## Activity Log

- 2026-06-04T16:39:16Z – unknown – Moved to for_review
