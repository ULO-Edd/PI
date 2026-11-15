---
work_package_id: WP05
title: Synchronous Capacity Booking & Rollbacks
lane: "done"
dependencies: [WP04]
base_branch: 002-backend-microservices-synchronous-core-WP04
base_commit: 15aec63d6812f274d96d019f7a02ebe35082dd97
created_at: '2026-06-01T16:39:23.518761+00:00'
subtasks:
- T014
- T015
- T016
- T017
phase: Phase 5 - Transaction Coordinator & Resilience Testing
assignee: ''
agent: "opencode"
shell_pid: "31460"
review_status: "approved"
reviewed_by: "Tiago Ramos"
history:
- timestamp: '2026-06-01T15:16:23Z'
  lane: planned
  agent: system
  shell_pid: ''
  action: Prompt generated via /spec-kitty.tasks
requirement_refs:
- FR-009
- FR-010
---

# Work Package Prompt: WP05 – Synchronous Capacity Booking & Rollbacks

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
Use language identifiers in code blocks: ````csharp`, ````bash`

---

## Objectives & Success Criteria

- Implement the ticket booking transaction endpoint `/api/registrations/book` in `RegistrationsService`.
- Build the atomic inventory rollback endpoint `/api/events/{id}/release` in `EventsService`.
- Write xUnit unit tests verifying controller authorizations and capacity checks.
- Build Playwright .NET API integration tests for concurrency, transaction states, and automatic rollbacks.

## Context & Constraints

- Refers to `kitty-specs/002-backend-microservices-synchronous-core/spec.md` (defining automatic rollbacks, transaction consistency, and p95 latency bounds).
- Requires a base branch with WP04:
  ```bash
  spec-kitty implement WP05 --base WP04
  ```

## Subtasks & Detailed Guidance

### Subtask T014 - Build booking coordinator API
- **Purpose**: Coordinate the distributed booking transaction synchronously.
- **Steps**:
  1. In `RegistrationsService`, implement `POST /api/registrations/book`:
     - Input: `{ eventId, attendeeId, correlationId }`.
  2. Steps in handler:
     - Synchronously call the Events service `/reserve` endpoint using the `EventsClient` (from WP04).
     - If reserve returns HTTP 400 (capacity exhausted): return HTTP 400 immediately.
     - If reserve succeeds, attempt to save the new registration entity in the local Registrations database.
     - If saving fails or throws an exception (database error): immediately call the Events service `/release` endpoint to roll back the reserved capacity, and return HTTP 400.
     - If database save succeeds, commit and return HTTP 201 Created.
- **Files**:
  - `/backend/src/Services/RegistrationsService/Controllers/RegistrationsController.cs`
- **Parallel?**: No. Coordinates the core transaction.

### Subtask T015 - Build capacity rollback endpoint
- **Purpose**: Provide a reliable mechanism to restore inventory upon registration failures.
- **Steps**:
  1. In `EventsService/EventsController.cs`, create `POST /api/events/{id}/release` endpoint.
  2. Implement atomic SQL restoration:
     - `UPDATE Events SET AvailableTickets = AvailableTickets + 1 WHERE Id = @Id AND AvailableTickets < Capacity`
  3. Propagate correlation ID in HTTP headers to trace the rollback action.
- **Files**:
  - `/backend/src/Services/EventsService/Controllers/EventsController.cs`
- **Parallel?**: Yes. Can be built concurrently with T014.

### Subtask T016 - Write xUnit unit tests
- **Purpose**: Assert controller action outcomes, field validations, and error states.
- **Steps**:
  1. Build `/backend/tests/UnitTests/` assembly.
  2. Write unit tests checking:
     - Invalid registration requests (missing fields) return HTTP 400.
     - Role authorizations block non-organisers from creating events.
     - Transaction handler executes rollback calls upon mocked database errors.
- **Files**:
  - `/backend/tests/UnitTests/RegistrationsControllerTests.cs`
- **Parallel?**: Yes. Can be prepared concurrently.

### Subtask T017 - Write Playwright E2E rollback checks
- **Purpose**: Execute end-to-end integration tests simulating transaction errors and concurrency.
- **Steps**:
  1. Build `/backend/tests/IntegrationTests/` assembly.
  2. Write Playwright .NET integration tests:
     - Simulate high concurrent booking requests and assert that capacity is never exceeded.
     - Simulate a database failure in RegistrationsService and assert that the capacity in EventsService is successfully rolled back.
     - Check Gateway-level rate limiter configurations.
- **Files**:
  - `/backend/tests/IntegrationTests/RollbackResilienceTests.cs`
- **Parallel?**: No. Serves as the final integration check.

## Risks & Mitigations

- **Zombie Reservations**: If a network error occurs during a rollback/release call, the Events service might preserve a reservation that Registrations abandoned. Mitigate by using correlation IDs and logging warning exceptions for manual operations alerting.

## Review Guidance

- Verify that any Registrations database exception triggers the rollback call successfully.
- Assert that concurrent bookings on the last ticket allow exactly one user to succeed and reject the rest.

## Activity Log

- 2026-06-01T15:16:23Z – system – lane=planned – Prompt generated via /spec-kitty.tasks
- 2026-06-01T16:39:27Z – opencode – shell_pid=17728 – lane=doing – Assigned agent via workflow command
- 2026-06-01T16:43:07Z – opencode – shell_pid=17728 – lane=for_review – Ready for review: Registrations transaction booking, Events atomic rollback releases, xUnit test coverage, and automated integration tests.
- 2026-06-01T16:43:16Z – opencode – shell_pid=31460 – lane=doing – Started review via workflow command
- 2026-06-01T16:43:28Z – opencode – shell_pid=31460 – lane=done – Review passed: Transaction book coordinators, atomic capacity release rollbacks, and full xUnit/Resilience testing suites are fully verified.
