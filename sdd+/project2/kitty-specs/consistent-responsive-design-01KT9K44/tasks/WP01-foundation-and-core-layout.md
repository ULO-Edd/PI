---
work_package_id: WP01
title: Foundation and Core Layout
dependencies: []
requirement_refs:
- FR-001
- FR-002
- FR-003
- FR-006
planning_base_branch: feature/sdd-plus-project2
merge_target_branch: feature/sdd-plus-project2
branch_strategy: Planning artifacts for this feature were generated on feature/sdd-plus-project2. During /spec-kitty.implement this WP may branch from a dependency-specific base, but completed changes must merge back into feature/sdd-plus-project2 unless the human explicitly redirects the landing branch.
subtasks:
- T001
- T002
- T003
- T004
history: []
authoritative_surface: frontend/src/
execution_mode: code_change
owned_files:
- frontend/src/App.vue
- frontend/src/views/auth/HomeView.vue
- frontend/src/views/auth/ForbiddenView.vue
tags: []
agent: "opencode"
shell_pid: "18228"
---

# WP01: Foundation and Core Layout

## Objective
Establish responsive app-wide frame, navigation menus, default layout adjustments, landing screens, and standard view scaling using Tailwind CSS.

## Context
Currently, the application has non-responsive static elements. This package establishes the foundation for responsiveness, starting with the outer layout wrapper (`App.vue`), the main entry view (`HomeView.vue`), and the access control view (`ForbiddenView.vue`).

## Detailed Guidance

### Subtask T001: Implement responsive, accessible main layout and navigation in App.vue
- **Objective**: Redesign navigation bar and root container in `App.vue` to support responsive mobile drawer/collapsible navigation, scale padding appropriately across breakpoints, and ensure keyboard accessibility.
- **Detailed Steps**:
  1. Update navigation items: use flexbox with breakpoint-dependent visible toggles. Mobile devices should feature a touch-friendly hamburger menu which collapses/expands the main nav links.
  2. Maintain accessibility: ensure all button controls are clickable/tappable, provide appropriate ARIA attributes for visibility states, and verify text contrast.
  3. Ensure container constraints: set root boundaries using `max-w-7xl` with responsive paddings `px-4 sm:px-6 lg:px-8`.
- **Validation**:
  - Main header scales gracefully from mobile width 320px up to 1440px desktop.
  - Hamburger toggles navigation correctly without broken text wraps on mobile port.

### Subtask T002: Enhance HomeView.vue with a responsive landing layout
- **Objective**: Introduce dynamic grids, scaling headings, responsive call-to-actions, and consistent spacings in `HomeView.vue`.
- **Detailed Steps**:
  1. Redesign Hero section: adapt grid to single-column on mobile, two-column layouts on `md` and above breakpoints.
  2. Use scaling sizes for typography (e.g., `text-3xl sm:text-4xl md:text-5xl font-extrabold`).
  3. Form responsive container boundaries to avoid bleed-off on small screens.
- **Validation**:
  - Check text overlap at 375px viewport.
  - Call-to-action buttons scale to full width on mobile viewports.

### Subtask T003: Make ForbiddenView.vue clean and responsive
- **Objective**: Format access denial view to display elegantly on all screen aspect ratios.
- **Detailed Steps**:
  1. Center-align notice cards utilizing Tailwind flex column wrappers.
  2. Implement mobile-friendly navigation buttons back to homepage or login screen.
- **Validation**:
  - View fits cleanly into smallest target size without scrolling.

### Subtask T004: Validate responsiveness across breakpoints (sm, md, lg, xl) for foundational layouts
- **Objective**: Run overall layouts review, address minor overlapping text, spacing drifts, and build successfully.
- **Validation**:
  - Run linting and Vite build checks.

## Activity Log

- 2026-06-04T16:39:10Z – unknown – Moved to for_review
- 2026-06-04T16:40:30Z – opencode – shell_pid=18228 – Started review via action command
