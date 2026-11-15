---
work_package_id: WP01
title: Frontend Foundation and Security
lane: "done"
dependencies: []
base_branch: main
base_commit: 1d3858c6f3508aab528838dacdcba8d19d43ecf9
created_at: '2026-06-01T14:07:57.946080+00:00'
subtasks:
- T001
- T002
- T003
- T004
phase: Phase 1 - Scaffolding & Setup
assignee: ''
agent: "opencode"
shell_pid: "38884"
review_status: "approved"
reviewed_by: "Tiago Ramos"
history:
- timestamp: '2026-06-01T13:35:24Z'
  lane: planned
  agent: system
  shell_pid: ''
  action: Prompt generated via /spec-kitty.tasks
requirement_refs:
- FR-003
---

# Work Package Prompt: WP01 - Frontend Foundation and Security

## ⚠️ IMPORTANT: Review Feedback Status

**Read this first if you are implementing this task!**

- **Has review feedback?**: Check the `review_status` field above. If it says `has_feedback`, scroll to the **Review Feedback** section immediately (right below this notice).
- **You must address all feedback** before your work is complete. Feedback items are your implementation TODO list.
- **Mark as acknowledged**: When you understand the feedback and begin addressing it, update `review_status: acknowledged` in the frontmatter.
- **Report progress**: As you address each feedback item, update the Activity Log explaining what you changed.

---

## Review Feedback

*[This section is empty initially. Reviewers will populate it if the work is returned from review. If you see feedback here, treat each item as a must-do before completion.]*

---

## Markdown Formatting
Wrap HTML/XML tags in backticks: `` `<div>` ``, `` `<script>` ``
Use language identifiers in code blocks: ````javascript`, ````bash`

---

## Objectives & Success Criteria

- Establish the core SPA Vue 3, Pinia, PrimeVue, and Tailwind CSS development structure under the `/frontend` subfolder.
- Construct the base HTTP service wrapper with cookie credentials enabling HTTP-only secure cookie session sharing.
- Configure Vue Router with strict route authentication checking and automatic redirection guards.
- Implement the authentication Pinia store coordinating secure logins and role permissions.

## Context & Constraints

- Refers to `.kittify/memory/constitution.md` (defining PrimeVue and Tailwind CSS integration).
- Refers to `kitty-specs/001-frontend-application-client/plan.md`.
- Honors secure cookie-based session management, preventing sensitive tokens from being read by client-side scripts.

## Subtasks & Detailed Guidance

### Subtask T001 - Scaffold Frontend workspace structure
- **Purpose**: Create a clean, typed development environment under `/frontend` root.
- **Steps**:
  1. Initialize Vite project with Vue 3, TypeScript, and ESLint configurations.
  2. Configure `tailwind.config.js` with basic theme setups.
  3. Initialize PrimeVue config in `src/main.ts` using the tailwind-compatible theme.
  4. Ensure package.json defines npm scripts for `dev`, `lint`, and `test:unit`.
- **Files**:
  - `/frontend/package.json`
  - `/frontend/vite.config.ts`
  - `/frontend/tailwind.config.js`
  - `/frontend/src/main.ts`
- **Parallel?**: No. Must complete before any other task begins.

### Subtask T002 - Implement base HTTP client wrapper
- **Purpose**: Build custom HTTP client that supports credential sharing across subdomains.
- **Steps**:
  1. Create `frontend/src/services/api.ts` exposing a custom Axios instance.
  2. Configure instance defaults: `withCredentials: true` (critical for HTTP-only cookies).
  3. Expose standard methods: `get`, `post`, `put`, `delete`.
  4. Implement an interceptor that handles HTTP 401 Unauthorized globally, triggering state logout in the Pinia auth store.
- **Files**:
  - `/frontend/src/services/api.ts`
- **Parallel?**: Yes. Can proceed once scaffolding (T001) is complete.

### Subtask T003 - Configure Vue Router with navigation guards
- **Purpose**: Map core pages and protect private paths (Dashboard, Check-In).
- **Steps**:
  1. Configure `/frontend/src/router/index.ts` with definitions for:
     - `LandingPage` (public, `/`)
     - `OrganiserDashboard` (private, `/dashboard`)
     - `LiveCheckIn` (private, `/check-in`)
  2. Mark private paths using `meta: { requiresAuth: true }`.
  3. Implement global `beforeEach` navigation guard:
     - Verify auth state using the Auth Pinia store.
     - Redirection rules: If route requires auth and state is unauthenticated, redirect to Landing Page.
- **Files**:
  - `/frontend/src/router/index.ts`
- **Parallel?**: Yes. Can run alongside T002 once T001 is finished.

### Subtask T004 - Implement Auth Pinia Store
- **Purpose**: Maintain global security session state and role definitions.
- **Steps**:
  1. Create `/frontend/src/stores/auth.ts` containing the following state:
     - `isAuthenticated: boolean`
     - `user: { id: string, email: string, role: 'Admin' | 'Organiser' | 'Attendee' } | null`
  2. Expose actions for `login(credentials)`, `logout()`, and `checkSession()`.
  3. Implement a reactive role getter `isOrganiser` that returns true if the user's role is Admin or Organiser.
- **Files**:
  - `/frontend/src/stores/auth.ts`
- **Parallel?**: Yes. Can run concurrently with T002 and T003.

## Risks & Mitigations

- **Cookie Sharing issues**: Ensure the development Vite proxy is correctly configured to forward cookies to the backend gateway port. Configure proxy options with `changeOrigin: true`.

## Review Guidance

- Reviewer must verify that unauthenticated requests to `/dashboard` trigger router redirects immediately.
- Ensure the Axios client is configured with `withCredentials: true` to prevent sessions from dropping.

## Activity Log

- 2026-06-01T13:35:24Z – system – lane=planned – Prompt generated via /spec-kitty.tasks
- 2026-06-01T14:08:02Z – opencode – shell_pid=37684 – lane=doing – Assigned agent via workflow command
- 2026-06-01T14:13:18Z – opencode – shell_pid=37684 – lane=for_review – Ready for review: Scaffolding, router guards, Axios wrappers, and Pinia auth stores implemented and successfully compiled.
- 2026-06-01T14:51:53Z – opencode – shell_pid=38884 – lane=doing – Started review via workflow command
- 2026-06-01T14:52:07Z – opencode – shell_pid=38884 – lane=done – Review passed: Scaffolding, Axios base configuration, secure route guards, and Pinia stores are fully verified and successfully built.
