---
work_package_id: WP04
title: Registrations Service & Inter-Service Client
lane: "done"
dependencies: [WP01, WP03]
base_branch: 002-backend-microservices-synchronous-core-WP03
base_commit: d503fc2d2dd38592ccff9931e39de2a418032fae
created_at: '2026-06-01T16:27:50.068622+00:00'
subtasks:
- T011
- T012
- T013
phase: Phase 4 - Registrations and Inter-Service Setup
assignee: ''
agent: "opencode"
shell_pid: "41856"
review_status: "approved"
reviewed_by: "Tiago Ramos"
history:
- timestamp: '2026-06-01T15:16:23Z'
  lane: planned
  agent: system
  shell_pid: ''
  action: Prompt generated via /spec-kitty.tasks
requirement_refs:
- FR-006
- FR-007
- FR-008
---

# Work Package Prompt: WP04 – Registrations Service & Inter-Service Client

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

- Create the `RegistrationsService` Web API project with an isolated database.
- Create EventsService atomic capacity reservation API endpoint `/api/events/{id}/reserve`.
- Build the synchronous inter-service HTTP capacity-check client wrapper in RegistrationsService.

## Context & Constraints

- Refers to `kitty-specs/002-backend-microservices-synchronous-core/spec.md` (defining isolated datastores, zero cross-database joins, and synchronous inter-service capacity checks).
- Requires a base branch with WP03:
  ```bash
  spec-kitty implement WP04 --base WP03
  ```

## Subtasks & Detailed Guidance

### Subtask T011 - Scaffold RegistrationsService & EF Context
- **Purpose**: Setup RegistrationsService Web API and relational datastore context.
- **Steps**:
  1. Create `/backend/src/Services/RegistrationsService` Web API project.
  2. Create `Data/RegistrationsDbContext.cs` exposing a `Registrations` entity.
  3. Generate and apply initial database migrations using `dotnet ef migrations add InitialRegistrationsDb`.
- **Files**:
  - `/backend/src/Services/RegistrationsService/RegistrationsService.csproj`
  - `/backend/src/Services/RegistrationsService/Data/RegistrationsDbContext.cs`
- **Parallel?**: Yes. Database context setup can proceed concurrently.

### Subtask T012 - Build Events capacity reserve API
- **Purpose**: Expose an atomic capacity verification and deduction API.
- **Steps**:
  1. In `EventsService/EventsController.cs`, create `POST /api/events/{id}/reserve` endpoint.
  2. Implement atomic SQL deduction in the database (to prevent overbooking race conditions):
     - `UPDATE Events SET AvailableTickets = AvailableTickets - 1 WHERE Id = @Id AND AvailableTickets >= 1`
  3. If deduction affects 1 row: return HTTP 200 OK.
  4. If deduction affects 0 rows: return HTTP 400 Bad Request (Capacity exhausted).
- **Files**:
  - `/backend/src/Services/EventsService/Controllers/EventsController.cs`
- **Parallel?**: Yes. Can be built concurrently with T011.

### Subtask T013 - Build inter-service HTTP sync client
- **Purpose**: Synchronously query EventsService capacity status from RegistrationsService.
- **Steps**:
  1. In `RegistrationsService`, configure `HttpClient` factory targeting `EventsService`.
  2. Create `Services/EventsClient.cs` wrapper:
     - Method `ReserveCapacityAsync(eventId, correlationId)`.
     - Perform POST request to `http://localhost:5002/api/events/{eventId}/reserve`.
     - Propagate structured correlation ID in HTTP headers.
- **Files**:
  - `/backend/src/Services/RegistrationsService/Services/EventsClient.cs`
- **Parallel?**: Yes. Can proceed once T011 and T012 are done.

## Risks & Mitigations

- **Concurrecy Race condition**: Direct subtraction (`AvailableTickets = AvailableTickets - 1`) must include checks in the `WHERE` clause (`AND AvailableTickets >= 1`) to prevent negative ticket counts under race conditions.

## Review Guidance

- Verify that no direct database connection, references, or cross-database joins exist between `EventsService` and `RegistrationsService`.
- Assert capacity reservations decrease the tickets available count atomically in the Events database.

## Activity Log

- 2026-06-01T15:16:23Z – system – lane=planned – Prompt generated via /spec-kitty.tasks
- 2026-06-01T16:27:52Z – opencode – shell_pid=35984 – lane=doing – Assigned agent via workflow command
- 2026-06-01T16:37:11Z – opencode – shell_pid=35984 – lane=for_review – Ready for review: RegistrationsService database context, EventsService atomic ticket reservation endpoint, and secure HTTP inter-service capacity-checking clients.
- 2026-06-01T16:37:32Z – opencode – shell_pid=41856 – lane=doing – Started review via workflow command
- 2026-06-01T16:38:02Z – opencode – shell_pid=41856 – lane=done – Review passed: RegistrationsService database migrations, EventsService atomic ticket reservation endpoint, and secure HTTP inter-service client wrapper build successfully with zero errors.
