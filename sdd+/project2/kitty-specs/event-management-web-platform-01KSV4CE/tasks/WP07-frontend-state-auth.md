---
work_package_id: WP07
title: Frontend State Management & Authentication Views
dependencies:
- WP06
requirement_refs:
- FR-001
- FR-002
- NFR-002
- NFR-003
planning_base_branch: feature/sdd-plus-project2
merge_target_branch: feature/sdd-plus-project2
branch_strategy: Planning artifacts for this feature were generated on feature/sdd-plus-project2. During /spec-kitty.implement this WP may branch from a dependency-specific base, but completed changes must merge back into feature/sdd-plus-project2 unless the human explicitly redirects the landing branch.
subtasks:
- T029
- T030
- T031
- T032
- T033
phase: Phase 7 - State & Access Controls
assignee: ''
agent: "opencode"
shell_pid: "8932"
history:
- timestamp: '2026-05-30T00:31:40Z'
  agent: system
  action: Prompt generated via /spec-kitty.tasks
authoritative_surface: frontend/src/
execution_mode: code_change
owned_files:
- frontend/src/store/**
- frontend/src/router/**
- frontend/src/views/auth/**
- frontend/src/components/auth/**
tags: []
---

# Work Package Prompt: WP07 – Frontend State Management & Authentication Views

## Objective
Implement Pinia state, router guarding, accessible registration/login forms, and JWT status checks.

---

## Subtask Details

### Subtask T029: Implement Pinia Store Auth
*   **Purpose:** Centrally store JWT tokens, parse claims, and handle local storage.
*   **Steps:**
    1. Construct auth store tracking roles and verification status.
*   **Files:** `frontend/src/store/auth.js`
*   **Validation:** Store maintains states across soft page reloads.

### Subtask T030: Set up Navigation Route Guards
*   **Purpose:** Protect client views using dynamic router metadata checks.
*   **Steps:**
    1. Append navigation gates to verify authorization status before entry.
*   **Files:** `frontend/src/router/index.js`
*   **Validation:** Redirects unauthorized users straight back to login page.

### Subtask T031: Design Accessible Sign-up Forms
*   **Purpose:** Build accessible signup forms for Attendees and Organizers.
*   **Steps:**
    1. Code semantic inputs with clear aria attributes and keyboard controls.
*   **Files:** `frontend/src/views/auth/RegisterView.vue`
*   **Validation:** Forms conform to basic WCAG semantic validation.

### Subtask T032: Build Login Portal
*   **Purpose:** Implement login page handling redirects.
*   **Steps:**
    1. Build login fields and hook submission handlers up to Go API.
*   **Files:** `frontend/src/views/auth/LoginView.vue`
*   **Validation:** Correct credentials fetch token and save to store.

### Subtask T033: Form Validation Feedback
*   **Purpose:** Show accessible alerts and active error feedback.
*   **Steps:**
    1. Connect error elements using `aria-live` regions for instant feedback.
*   **Files:** `frontend/src/components/auth/Feedback.vue`
*   **Validation:** Screen readers announce errors upon submission failures.

## Activity Log

- 2026-05-31T10:23:11Z – opencode – shell_pid=13464 – Started implementation via action command
- 2026-05-31T10:25:40Z – opencode – shell_pid=13464 – Pinia store JWT client-side parsing, Vue Router view role guarding, and fully accessible registration and login forms implemented.
- 2026-05-31T10:26:12Z – opencode – shell_pid=8932 – Started review via action command
- 2026-05-31T10:26:53Z – opencode – shell_pid=8932 – Review passed: Pinia auth store, route guards, accessible registration/login forms, and validation feedback states fully implemented and validated with 100% Vite production build success.
