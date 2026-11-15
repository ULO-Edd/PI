---
work_package_id: WP03
title: Frontend Stability & Modernization
lane: "done"
dependencies: [WP02]
base_branch: 005-infrastructure-connectivity-router-modernization-WP02
base_commit: a0c4504d139c45293bdef9426afd883e6973afad
created_at: '2026-06-08T13:20:17.627998+00:00'
subtasks: [T010, T011, T012, T013, T014]
requirement_refs:
- FR-005
- FR-006
shell_pid: "29516"
agent: "gemma-4-31b-it"
reviewed_by: "Tiago Ramos"
review_status: "approved"
---

# Work Package: Frontend Stability & Modernization

## Objective
Refactor navigation guards to eliminate framework warnings and implement resilient network error handling.

## Context & Constraints
- **Goal**: A clean browser console and user-friendly error states for network failures.
- **Standard**: Vue Router 4, Axios interceptors.
- **Tools**: Vue 3, Axios, PrimeVue Toasts.

## Subtasks & Detailed Guidance

### Subtask T010: Audit and refactor Vue Router navigation guards
**Purpose**: Eliminate deprecation warnings in the console.
**Steps**:
1. Identify all `beforeEach` and `beforeEnter` guards in the router configuration.
2. Replace the `next()` callback pattern with the functional return pattern:
   - Return `true` to allow navigation.
   - Return `false` to cancel.
   - Return a route location object (e.g., `{ name: 'Login' }`) to redirect.
3. Verify that all authentication and authorization logic remains intact.
**Validation**:
- [ ] Console is free of "navigation guard" deprecation warnings during full app traversal.

### Subtask T011: Implement a global Axios interceptor for network error detection
**Purpose**: Centralize the detection of connectivity failures.
**Steps**:
1. Create or update the global Axios instance.
2. Add a response interceptor that catches errors.
3. Specifically detect errors where `error.response` is undefined and `error.message` contains "Network Error" or is a `ERR_CONNECTION_REFUSED` type.
**Validation**:
- [ ] Interceptor correctly identifies network outages vs. 4xx/5xx API errors.

### Subtask T012: Implement a Toast Notification system for connectivity alerts
**Purpose**: Provide immediate, non-intrusive feedback to the user.
**Steps**:
1. Use PrimeVue's `Toast` component.
2. Create a `useToast` service or Pinia store to trigger notifications globally.
3. Define a specific "Network Error" toast style (e.g., severity: 'error', summary: 'Connection Error', detail: 'Backend services are unavailable. Please check if Docker is running.').
**Validation**:
- [ ] Toasts appear correctly when triggered manually.

### Subtask T013: Integrate Axios interceptor with Toast system
**Purpose**: Automatically notify the user of connectivity issues.
**Steps**:
1. Import the Toast service into the Axios interceptor.
2. Trigger the "Network Error" toast whenever the interceptor catches a connectivity failure.
**Validation**:
- [ ] Stopping the Docker containers and making a request triggers the connectivity toast.

### Subtask T014: Verify full end-to-end flow
**Purpose**: Ensure all components work together as a stable system.
**Steps**:
1. Run `docker-compose up -d` and `npm run dev`.
2. Verify the app loads and navigates without warnings.
3. Verify API requests work.
4. Stop the backend and verify the connectivity toast appears.
**Validation**:
- [ ] End-to-end flow matches the success criteria in the spec.

## Definition of Done
- [ ] No navigation guard warnings in console.
- [ ] Network failures trigger a clear Toast notification.
- [ ] Application remains responsive during outages.

## Risks & Mitigations
- **Risk**: Guard refactoring might introduce infinite redirect loops.
- **Mitigation**: Thoroughly test all route transitions, especially between public and private areas.

## Activity Log

- 2026-06-08T13:20:25Z – gemma-4-31b-it – shell_pid=13456 – lane=doing – Assigned agent via workflow command
- 2026-06-08T13:24:46Z – gemma-4-31b-it – shell_pid=13456 – lane=for_review – Ready for review: Refactored Vue Router guards to remove deprecation warnings, implemented global Axios interceptor for network errors, and added a Toast Notification system.
- 2026-06-08T13:32:18Z – gemma-4-31b-it – shell_pid=29516 – lane=doing – Started review via workflow command
- 2026-06-08T13:33:14Z – gemma-4-31b-it – shell_pid=29516 – lane=done – Review passed: Refactored Vue Router guards to remove deprecation warnings, implemented global Axios interceptor for network errors, and added a Toast Notification system.
