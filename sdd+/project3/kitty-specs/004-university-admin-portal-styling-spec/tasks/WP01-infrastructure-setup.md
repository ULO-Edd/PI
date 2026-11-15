---
work_package_id: WP01
title: Infrastructure & Configuration Setup
lane: "done"
dependencies: []
base_branch: main
base_commit: 6e0a725751f4b0926258508c30e4b39682b9ae74
created_at: '2026-06-05T15:30:17.344632+00:00'
subtasks:
- T001
- T002
- T003
- T004
- T005
- T006
phase: Phase 1 - Foundation
assignee: ''
agent: "opencode"
shell_pid: "25140"
review_status: "approved"
reviewed_by: "Tiago Ramos"
history:
- timestamp: '2026-06-05T14:30:00Z'
  lane: planned
  agent: system
  shell_pid: ''
  action: Prompt generated via /spec-kitty.tasks
---

# Work Package Prompt: WP01 – Infrastructure & Configuration Setup

## Objective
Establish the foundational development environment with Tailwind CSS, PostCSS, and PrimeVue unstyled mode.

## Context & Constraints
- **Stack**: Vue 3, Pinia, Tailwind CSS, PostCSS.
- **UI Library**: PrimeVue must be configured in `unstyled: true` mode.
- **Design Source**: The design tokens are defined in `contracts/tailwind-config.json`.
- **Accessibility**: The configuration must support WCAG 2.1 AA standards.

## Subtasks & Detailed Guidance

### Subtask T001 – Install Tailwind CSS and PostCSS dependencies
**Purpose**: Install the necessary packages to enable Tailwind CSS and PostCSS processing.

**Steps**:
1. Use `npm install` or `yarn add` to add `tailwindcss`, `postcss`, and `autoprefixer` as dev dependencies.
2. Verify installation by checking `package.json`.

**Files**:
- `package.json`

**Validation**:
- [ ] `tailwindcss`, `postcss`, and `autoprefixer` appear in `devDependencies`.

### Subtask T002 – Configure `tailwind.config.js`
**Purpose**: Integrate the project's design tokens into the Tailwind configuration.

**Steps**:
1. Create/update `tailwind.config.js` in the project root.
2. Import/copy the color, spacing, and typography tokens from `contracts/tailwind-config.json`.
3. Ensure the `content` array includes all relevant Vue files (`./src/**/*.{vue,js,ts,jsx,tsx}`).

**Files**:
- `tailwind.config.js`

**Validation**:
- [ ] `tailwind.config.js` contains the `theme.extend` object from `contracts/tailwind-config.json`.
- [ ] Tailwind utilities for the new colors (e.g., `text-primary-500`) are available.

### Subtask T003 – Configure PostCSS
**Purpose**: Ensure PostCSS is properly configured to process Tailwind and autoprefix CSS.

**Steps**:
1. Create or update `postcss.config.js` in the project root.
2. Add `tailwindcss` and `autoprefixer` to the plugins array.

**Files**:
- `postcss.config.js`

**Validation**:
- [ ] `postcss.config.js` is present and correctly configured.

### Subtask T004 – Initialize Global CSS
**Purpose**: Set up the entry point for Tailwind's utility classes.

**Steps**:
1. Create or update `src/assets/css/main.css`.
2. Add the `@tailwind base;`, `@tailwind components;`, and `@tailwind utilities;` directives.
3. Import this CSS file into the main entry point (e.g., `src/main.js`).

**Files**:
- `src/assets/css/main.css`
- `src/main.js`

**Validation**:
- [ ] Tailwind utility classes work in the application.

### Subtask T005 – Configure PrimeVue Unstyled Mode
**Purpose**: Enable unstyled mode for PrimeVue to allow complete Tailwind styling control.

**Steps**:
1. In `src/main.js` (or your plugin setup file), configure PrimeVue with `unstyled: true`.
2. Ensure `PrimeVue` is properly installed as a plugin.

**Files**:
- `src/main.js`

**Validation**:
- [ ] PrimeVue components render without default PrimeVue styling.

### Subtask T006 – Set up Base CSS for PrimeVue Overrides
**Purpose**: Create a dedicated file for CSS-based overrides and PrimeVue preset integration.

**Steps**:
1. Create `src/assets/css/primevue-presets.css`.
2. Import this file into `src/assets/css/main.css` after the Tailwind directives.

**Files**:
- `src/assets/css/primevue-presets.css`
- `src/assets/css/main.css`

**Validation**:
- [ ] `primevue-presets.css` is imported and available in the build pipeline.

## Definition of Done
- [ ] All six subtasks are completed.
- [ ] The application compiles without CSS errors.
- [ ] Tailwind utility classes and PrimeVue components render without default/conflicting styles.

## Risks & Mitigations
- **Risk**: Version mismatch between Tailwind and PrimeVue.
- **Mitigation**: Use `unstyled: true` and carefully manage the CSS cascade in `main.css`.

- FR1, FR6, FR7

## Activity Log

- 2026-06-05T15:30:22Z – opencode – shell_pid=25324 – lane=doing – Assigned agent via workflow command
- 2026-06-05T15:35:43Z – opencode – shell_pid=25324 – lane=for_review – Ready for review: WP01 infrastructure setup complete
- 2026-06-05T16:52:16Z – opencode – shell_pid=25140 – lane=doing – Started review via workflow command
- 2026-06-05T16:54:53Z – opencode – shell_pid=25140 – lane=done – Review passed: Infrastructure setup complete (Tailwind v4, PostCSS, and PrimeVue unstyled mode configured correctly)
- 2026-06-05T17:16:06Z – opencode – shell_pid=25140 – lane=done – Review passed
- 2026-06-05T17:30:21Z – opencode – shell_pid=25140 – lane=doing – Moved to doing
- 2026-06-05T17:30:32Z – opencode – shell_pid=25140 – lane=done – Back to done
