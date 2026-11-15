---
work_package_id: WP02
title: Design Token System Implementation
lane: "done"
dependencies: []
base_branch: main
base_commit: 4ad6497f45fb552a330d3e4c9771b7f18780c7d2
created_at: '2026-06-05T15:42:06.809919+00:00'
subtasks:
- T007
- T008
- T009
- T010
- T011
phase: Phase 1 - Design
assignee: ''
agent: "opencode"
shell_pid: "19256"
review_status: "has_feedback"
reviewed_by: "Tiago Ramos"
review_feedback_file: "C:\Users\tiago\AppData\Local\Temp\spec-kitty-review-feedback-WP02.md"
history:
- timestamp: '2026-06-05T14:30:00Z'
  lane: planned
  agent: system
  shell_pid: ''
  action: Prompt generated via /spec-kitty.tasks
---

# Work Package Prompt: WP02 – Design Token System Implementation

## Objective
Implement the centralized design token system for colors, typography, spacing, and shadows, ensuring they are accessible via Tailwind CSS.

## Context & Constraints
- **Source of Truth**: `contracts/tailwind-config.json`.
- **Implementation**: All tokens must be integrated into `tailwind.config.js`.
- **Goal**: Enable a "single source of truth" for the entire design system.

## Subtasks & Detailed Guidance

### Subtask T007 – Implement Color Palette
**Purpose**: Define the primary, secondary, and semantic colors in Tailwind.

**Steps**:
1. Map the color tokens from `contracts/tailwind-config.json` into the `theme.extend.colors` section of `tailwind.config.js`.
2. Ensure semantic colors (success, warning, danger, info) are correctly defined.

**Files**:
- `tailwind.config.js`

**Validation**:
- [ ] Tailwind classes like `bg-primary-500` or `text-success-500` work correctly.

### Subtask T008 – Implement Typography Scale
**Purpose**: Establish a consistent typographic hierarchy.

**Steps**:
1. Define font families (e.g., Sans, Mono) in `tailwind.config.js`.
2. Implement a typographic scale using font size and line height utilities (e.g., `text-sm`, `text-base`, `text-lg`).
3. Ensure font sizes are consistent with the design requirements.

**Files**:
- `tailwind.config.js`

**Validation**:
- [ ] Typography utilities (e.g., `text-xl`, `font-sans`) are functional.

### Subtask T009 – Implement Spacing and Sizing System
**Purpose**: Create a predictable spacing grid.

**Steps**:
1. Add custom spacing tokens (e.g., `xs`, `sm`, `md`, `lg`, `xl`) to `tailwind.config.js`.
2. Verify these can be used with margin (`m-`), padding (`p-`), and width/height (`w-`, `h-`) utilities.

**Files**:
- `tailwind.config.js`

**Validation**:
- [ ] Utilities like `p-md` or `m-xl` are functional.

### Subtask T010 – Implement Shadow and Border Radius Tokens
**Purpose**: Define the elevation and corner radius system.

**Steps**:
1. Add custom `boxShadow` and `borderRadius` tokens to `tailwind.config.js`.
2. Ensure they match the intended "modern and streamlined" aesthetic.

**Files**:
- `tailwind.config.js`

**Validation**:
- [ ] Classes like `shadow-md` (custom) or `rounded-lg` (custom) work correctly.

### Subtask T011 – Create Design Token Sandbox
**Purpose**: Provide a visual way to verify all implemented tokens.

**Steps**:
1. Create a new Vue component `src/components/tokens/TokenSandbox.vue`.
2. Display all colors, font sizes, spacing values, and shadows as visual examples (e.g. color swatches, text samples).
3. Mount this component in the main `App.vue` temporarily for verification.

**Files**:
- `src/components/tokens/TokenSandbox.vue`
- `src/App.vue`

**Validation**:
- [ ] The sandbox component correctly displays all tokens defined in `tailwind.config.js`.

## Definition of Done
- [ ] All five subtasks are completed.
- [ ] All design tokens are accessible via Tailwind utility classes.
- [ ] The Design Token Sandbox provides a 100% accurate visual representation of the system.

## Risks & Mitigations
- **Risk**: Inconsistent token naming or values.
- **Mitigation**: Cross-reference strictly with `contracts/tailwind-config.json`.

- FR1

## Review Feedback

**Reviewed by**: Tiago Ramos
**Status**: ❌ Changes Requested
**Date**: 2026-06-05
**Feedback file**: `C:\Users\tiago\AppData\Local\Temp\spec-kitty-review-feedback-WP02.md`

**Issue 1**: `frontend/tailwind.config.js` contains a syntax error in the `fontFamily.sans` array. The string `'Segoe UI Emoji"` uses a mix of single and double quotes, which will cause the Tailwind build to fail. 

Please fix the quotes in the `fontFamily` configuration to be consistent (e.g., use all single quotes).


## Activity Log

- 2026-06-05T15:42:13Z – opencode – shell_pid=26344 – lane=doing – Assigned agent via workflow command
- 2026-06-05T16:00:42Z – opencode – shell_pid=26344 – lane=for_review – Ready for review: Design token system and sandbox implemented successfully
- 2026-06-05T16:55:35Z – opencode – shell_pid=19256 – lane=doing – Started review via workflow command
- 2026-06-05T16:57:22Z – opencode – shell_pid=19256 – lane=planned – Moved to planned
- 2026-06-05T17:14:30Z – opencode – shell_pid=19256 – lane=planned – Moved to planned
- 2026-06-05T17:14:45Z – opencode – shell_pid=19256 – lane=doing – Moved to doing
- 2026-06-05T17:15:25Z – opencode – shell_pid=19256 – lane=planned – Moved to planned
- 2026-06-05T17:15:36Z – opencode – shell_pid=19256 – lane=doing – Moved to doing
- 2026-06-05T17:15:50Z – opencode – shell_pid=19256 – lane=planned – Moved to planned
- 2026-06-05T17:28:33Z – opencode – shell_pid=19256 – lane=doing – Moved to doing
- 2026-06-05T17:29:02Z – opencode – shell_pid=19256 – lane=for_review – Moved to for_review
- 2026-06-05T17:29:25Z – opencode – shell_pid=19256 – lane=planned – Moved to planned
- 2026-06-05T17:29:39Z – opencode – shell_pid=19256 – lane=doing – Moved to doing
- 2026-06-05T17:29:50Z – opencode – shell_pid=19256 – lane=for_review – Moved to for_review
- 2026-06-05T17:30:08Z – opencode – shell_pid=19256 – lane=done – Testing count
- 2026-06-05T17:30:52Z – opencode – shell_pid=19256 – lane=planned – Moved to planned
- 2026-06-05T17:31:08Z – opencode – shell_pid=19256 – lane=doing – Moved to doing
- 2026-06-05T17:31:29Z – opencode – shell_pid=19256 – lane=for_review – Moved to for_review
- 2026-06-05T17:31:52Z – opencode – shell_pid=19256 – lane=done – Testing
- 2026-06-05T17:32:07Z – opencode – shell_pid=19256 – lane=planned – Moved to planned
- 2026-06-05T17:32:21Z – opencode – shell_pid=19256 – lane=doing – Moved to doing
- 2026-06-05T17:32:39Z – opencode – shell_pid=19256 – lane=for_review – Moved to for_review
- 2026-06-05T17:33:40Z – opencode – shell_pid=19256 – lane=planned – Moved to planned
- 2026-06-08T12:00:30Z – opencode – shell_pid=19256 – lane=done – Force moved to done per user request
