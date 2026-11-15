---
work_package_id: WP04
title: Responsive Layout & Accessibility Polish
lane: "done"
dependencies: []
base_branch: main
base_commit: a4eb107cbd475897c6051f2065fcc7d97e073215
created_at: '2026-06-05T16:27:47.376188+00:00'
subtasks:
- T017
- T018
- T019
- T020
phase: Phase 3 - Quality & Polish
assignee: ''
agent: "gemma-4-31b-it"
shell_pid: "27512"
review_status: "has_feedback"
reviewed_by: "Tiago Ramos"
review_feedback_file: "C:\Users\tiago\OneDrive\Documents\dev\SDD-Project-3-SecondTry\feedback-wp04.md"
history:
- timestamp: '2026-06-05T14:30:00Z'
  lane: planned
  agent: system
  shell_pid: ''
  action: Prompt generated via /spec-kitty.tasks
---

# Work Package Prompt: WP04 – Responsive Layout & Accessibility Polish

## Objective
Ensure the styling system is fully responsive across all target devices and strictly adheres to WCAG 2.1 AA accessibility standards.

## Context & Constraints
- **Goal**: A seamless experience for university administrators on desktop, tablet, and mobile.
- **Standard**: WCAG 2.1 AA compliance.
- **Tools**: Tailwind breakpoints, Lighthouse, Axe-core.

## Subtasks & Detailed Guidance

### Subtask T017 – Responsive Grid Strategy Implementation
**Purpose**: Implement the mobile-first responsive grid strategy in standard application layouts.

**Steps**:
1. Create a sample layout (e.g., `DashboardLayout.vue`) using Tailwind's responsive grid/flex utilities.
2. Ensure the layout adapts from a single-column mobile view to a multi-column desktop view.
3. Verify that content does not overflow or become illegible on small viewports.

**Files**:
- `src/layouts/DashboardLayout.vue` (new)
- `src/views/Dashboard.vue` (new)

**Validation**:
- [ ] Layout transitions smoothly between `sm`, `md`, `lg`, and `xl` breakpoints.
- [ ] No horizontal scrolling on mobile.

### Subtask T018 – Accessible Interactive States
**Purpose**: Standardize focus and interactive states to ensure keyboard navigability.

**Steps**:
1. Implement a global focus ring style in `src/assets/css/main.css` using Tailwind's `ring` utilities.
2. Ensure all interactive components (Buttons, Inputs, Links) from WP03 use this standardized focus style.
3. Verify that focus states are visually distinct and meet contrast requirements.

**Files**:
- `src/assets/css/main.css`
- Component presets in `src/assets/css/primevue-presets.js`

**Validation**:
- [ ] Keyboard navigation (Tab key) shows a clear, high-contrast focus indicator on all interactive elements.

### Subtask T019 – Contrast & High-Contrast Mode Support
**Purpose**: Implement color contrast verification and high-contrast mode support.

**Steps**:
1. Audit all color tokens in `tailwind.config.js` to ensure they meet the 4.5:1 ratio.
2. Implement support for the `prefers-contrast: more` media query using Tailwind's custom variants or standard CSS.
3. Ensure the UI remains functional and legible when high-contrast mode is active.

**Files**:
- `tailwind.config.js`
- `src/assets/css/main.css`

**Validation**:
- [ ] Colors meet minimum contrast requirements.
- [ ] UI is legible in high-contrast mode.

### Subtask T020 – Automated Accessibility Audits
**Purpose**: Validate the entire styling system using automated accessibility testing tools.

**Steps**:
1. Run a Lighthouse accessibility audit on the application in a browser.
2. Use Axe-core (or a similar tool) to scan core components (Buttons, Inputs, Tables).
3. Fix any critical accessibility violations identified.

**Tools**: Lighthouse, Axe-core.

**Validation**:
- [ ] Lighthouse accessibility score is 90 or higher.
- [ ] No critical/serious violations found by Axe-core.

## Definition of Done
- [ ] All four subtasks are completed.
- [ ] The application is fully responsive and accessible.
- [ ] Accessibility audits pass the defined criteria.

## Risks & Mitigations
- **Risk**: Fixing accessibility issues might compromise the "modern and streamlined" aesthetic.
- **Mitigation**: Use subtle but clear focus rings and high-contrast color variants that still align with the brand identity.

- FR2, FR4

## Review Feedback

**Reviewed by**: Tiago Ramos
**Status**: ❌ Changes Requested
**Date**: 2026-06-08
**Feedback file**: `C:\Users\tiago\OneDrive\Documents\dev\SDD-Project-3-SecondTry\feedback-wp04.md`

**Issue 1**: Violation of Work Package Isolation Rules. The implementation commit modified `kitty-specs/004-university-admin-portal-styling-spec/tasks.md` and several other WP files (WP01, WP02, WP03), resetting their statuses and clearing history. You must only modify the status of WP04.

**How to fix**: Revert all changes to files within the `kitty-specs/` directory. Only the implementation files in `frontend/` should be committed.


## Activity Log

- 2026-06-05T16:27:53Z – opencode – shell_pid=10632 – lane=doing – Assigned agent via workflow command
- 2026-06-05T16:46:16Z – opencode – shell_pid=10632 – lane=for_review – Ready for review: Responsive layout and accessibility polish implemented.
- 2026-06-05T18:36:08Z – opencode – shell_pid=6880 – lane=doing – Started review via workflow command
- 2026-06-08T11:30:42Z – opencode – shell_pid=6880 – lane=for_review – Ready for review: Implemented responsive grid strategy, standardized accessible focus states, corrected color contrast for WCAG 2.1 AA compliance, and enhanced high-contrast mode support.
- 2026-06-08T11:48:22Z – gemma-4-31b-it – shell_pid=27512 – lane=doing – Started review via workflow command
- 2026-06-08T11:50:15Z – gemma-4-31b-it – shell_pid=27512 – lane=planned – Moved to planned
- 2026-06-08T12:00:34Z – gemma-4-31b-it – shell_pid=27512 – lane=done – Force moved to done per user request
