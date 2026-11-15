---
work_package_id: WP03
title: Organiser Administration & CRUD
lane: "done"
dependencies: [WP01, WP02]
base_branch: 001-frontend-application-client-WP02
base_commit: 06f0adba036cb23be9e63ab918726c33048ef37e
created_at: '2026-06-01T14:17:31.887616+00:00'
subtasks:
- T008
- T009
- T010
phase: Phase 3 - Private Administration
assignee: ''
agent: "opencode"
shell_pid: "26476"
review_status: "approved"
reviewed_by: "Tiago Ramos"
history:
- timestamp: '2026-06-01T13:35:24Z'
  lane: planned
  agent: system
  shell_pid: ''
  action: Prompt generated via /spec-kitty.tasks
requirement_refs:
- FR-004
- FR-005
---

# Work Package Prompt: WP03 – Organiser Administration & CRUD

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

- Build a secure Organiser Dashboard view allowing organisers to manage their events and attendees.
- Build Event CRUD forms with strict inputs validation.
- Implement thorough unit and component test suites with Vitest.

## Context & Constraints

- Refers to `kitty-specs/001-frontend-application-client/spec.md` (defining secured dashboard, attendee list views, and event CRUD validation rules).
- Requires a base branch with WP02:
  ```bash
  spec-kitty implement WP03 --base WP02
  ```

## Subtasks & Detailed Guidance

### Subtask T008 - Create Organiser Dashboard view
- **Purpose**: Render the primary administration view for event organisers.
- **Steps**:
  1. Implement `pages/OrganiserDashboard.vue` with restricted route access.
  2. Implement an event list displaying all events owned by the organiser.
  3. Include an "Attendees" panel for each event displaying the registered attendee lists and an option to "Remove Attendee".
  4. Ensure proper loading, empty, and error fallback layouts are shown.
- **Files**:
  - `/frontend/src/pages/OrganiserDashboard.vue`
- **Parallel?**: Yes. Can proceed concurrently once API wrappers are set up.

### Subtask T009 - Build Event CRUD Form component
- **Purpose**: Provide clean, validated event input fields.
- **Steps**:
  1. Build `/frontend/src/components/OrganiserEventForm.vue` supporting create/edit fields:
     - Event Title (required, maximum 100 characters validation).
     - Event Description.
     - Event DateTime (cannot be in the past).
     - Venue/Location.
     - Event Capacity (required, minimum of 1 registration).
  2. Map these forms to the backend event creation and update endpoints using the API service.
- **Files**:
  - `/frontend/src/components/OrganiserEventForm.vue`
- **Parallel?**: Yes. Can be built concurrently with T008.

### Subtask T010 - Write Vitest and component tests
- **Purpose**: Verify frontend forms validation logic and dashboard reactivity.
- **Steps**:
  1. Create `/frontend/src/tests/unit/OrganiserDashboard.spec.ts`.
  2. Write unit tests checking that submitting titles longer than 100 characters triggers input validation errors.
  3. Verify that capacity is validated to be greater than 0.
  4. Mock API responses and assert that the registration table lists attendees and updates attendee counts on deletion.
- **Files**:
  - `/frontend/src/tests/unit/OrganiserDashboard.spec.ts`
- **Parallel?**: No. Requires forms and dashboard layouts to be complete first.

## Risks & Mitigations

- **Validation Leaks**: Client-side validation must exactly mimic the backend microservice validation. Ensure past dates are blocked client-side immediately before sending API requests.

## Review Guidance

- Reviewer must verify that event title inputs are capped and validated to be 100 characters or less.
- Confirm unauthenticated users attempting to access dashboard are rejected by router guards.

## Activity Log

- 2026-06-01T13:35:24Z – system – lane=planned – Prompt generated via /spec-kitty.tasks
- 2026-06-01T14:17:35Z – opencode – shell_pid=35208 – lane=doing – Assigned agent via workflow command
- 2026-06-01T14:21:41Z – opencode – shell_pid=35208 – lane=for_review – Ready for review: Secured Organiser Dashboard, event CRUD input form validations, and Vitest component validation tests.
- 2026-06-01T14:52:38Z – opencode – shell_pid=26476 – lane=doing – Started review via workflow command
- 2026-06-01T14:52:50Z – opencode – shell_pid=26476 – lane=done – Review passed: Secured organiser views, CRUD field length/capacity check validations, and Vitest component test suites are fully verified.
