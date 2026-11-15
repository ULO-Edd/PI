---
work_package_id: WP03
title: PrimeVue Component Styling (Unstyled Mode)
lane: "done"
dependencies: []
base_branch: main
base_commit: 6e659175864f05c655bd37c6b5d340969d544e4d
created_at: '2026-06-05T16:01:50.527717+00:00'
subtasks:
- T012
- T013
- T014
- T015
- T016
phase: Phase 2 - Component Implementation
assignee: ''
agent: "opencode"
shell_pid: "28016"
review_status: "approved"
reviewed_by: "Tiago Ramos"
history:
- timestamp: '2026-06-05T14:30:00Z'
  lane: planned
  agent: system
  shell_pid: ''
  action: Prompt generated via /spec-kitty.tasks
---

# Work Package Prompt: WP03 – PrimeVue Component Styling (Unstyled Mode)

## Objective
Create Tailwind-based presets for core PrimeVue components using the Pass Through (PT) property to achieve a consistent, modern, and streamlined UI.

## Context & Constraints
- **Library**: PrimeVue in **unstyled mode**.
- **Styling Tool**: Tailwind CSS via the `pt` (Pass Through) property.
- **Goal**: Ensure all components look cohesive and follow the design tokens defined in WP02.
- **Aesthetic**: Modern, streamlined, professional administrative interface.

## Subtasks & Detailed Guidance

### Subtask T012 – Primary Action Components
**Purpose**: Style the most frequently used interactive elements: Buttons and Links.

**Steps**:
1. Create a preset for the `Button` component in `src/assets/css/primevue-presets.js`.
2. Implement styles for `normal`, `hover`, `active`, `focus`, and `disabled` states using Tailwind.
3. Include styles for different severity levels (primary, secondary, success, danger, etc.) using the tokens from WP02.
4. Implement a basic `Link` component preset if applicable.

**Files**:
- `src/assets/css/primevue-presets.js`

**Validation**:
- [ ] Buttons change color on hover and have a clear focus ring.
- [ ] Severity levels (e.g., `severity="danger"`) are visually distinct.

### Subtask T013 – Input Elements
**Purpose**: Style form inputs to ensure a professional and consistent look.

**Steps**:
1. Create presets for `InputText`, `Select`, `Checkbox`, and `RadioButton`.
2. Ensure consistent padding, border-radius, and border-color tokens are applied.
3. Implement clear `focus` and `error` (if applicable) states.

**Files**:
- `src/assets/css/primevue-presets.js`

**Validation**:
- [ ] Inputs have consistent sizing and typography.
- [ ] Focus states are clearly visible and accessible.

### Subtask T014 – Layout and Container Components
**Purpose**: Style components that provide structure to the UI.

**Steps**:
1. Create presets for `Card`, `Panel`, and `Dialog`.
2. Ensure consistent use of shadows, border-radius, and padding tokens.
3. Implement responsive padding/margins for these containers.

**Files**:
- `src/assets/css/primevue-presets.js`

**Validation**:
- [ ] Containers (Cards, Panels) have consistent elevation (shadows) and corner rounding.
- [ ] Dialogs are properly centered and styled.

### Subtask T015 – Data Display Components
**Purpose**: Style components used for presenting information efficiently.

**Steps**:
1. Create presets for `Table`, `Tag`, and `Badge`.
2. For `Table`: Implement consistent cell padding, header styling, and row hover states.
3. For `Tag` and `Badge`: Implement color-coded semantic styles (success, warning, etc.).

**Files**:
- `src/assets/css/primevue-presets.js`

**Validation**:
- [ ] Tables are legible and professional.
- [ ] Tags and Badges use the correct semantic colors.

### Subtask T016 – Navigation and Feedback Components
**Purpose**: Style components that facilitate user flow and communication.

**Steps**:
1. Create presets for `Toast`, `Menu`, and `Breadcrumb`.
2. Ensure feedback elements (`Toast`) are prominent and clearly styled.
3. Ensure navigation elements (`Menu`, `Breadcrumb`) are clean and unobtrusive.

**Files**:
- `src/assets/css/primevue-presets.js`

**Validation**:
- [ ] Feedback notifications are highly visible.
- [ ] Navigation elements are intuitive and visually coherent.

## Definition of Done
- [ ] All five subtasks are completed.
- [ ] Core PrimeVue components are fully styled using Tailwind via the `pt` property.
- [ ] No default PrimeVue styling "leaks" into the application.

## Risks & Mitigations
- **Risk**: Complexity of the `pt` property for deeply nested components (e.g., Tables).
- **Mitigation**: Use the PrimeVue documentation to identify the exact internal element keys needed for styling.

**Requirements Refs**:
- FR5, FR6

## Activity Log

- 2026-06-05T16:01:57Z – opencode – shell_pid=26840 – lane=doing – Assigned agent via workflow command
- 2026-06-05T16:18:12Z – opencode – shell_pid=26840 – lane=for_review – Ready for review: PrimeVue components styled with Tailwind presets
- 2026-06-05T18:33:12Z – opencode – shell_pid=28016 – lane=doing – Started review via workflow command
- 2026-06-05T18:34:47Z – opencode – shell_pid=28016 – lane=done – Review passed: PrimeVue component presets are correctly implemented using Tailwind via the pt property, covering all required subtasks (T012-T016) and adhering to the design token system.
