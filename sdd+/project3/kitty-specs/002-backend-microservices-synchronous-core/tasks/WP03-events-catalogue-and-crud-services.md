---
work_package_id: WP03
title: Events Catalogue & CRUD Services
lane: "done"
dependencies: [WP01, WP02]
base_branch: 002-backend-microservices-synchronous-core-WP02
base_commit: f4785766b4da915c40030c5517b00c6b894a7b72
created_at: '2026-06-01T16:00:58.368791+00:00'
subtasks:
- T008
- T009
- T010
phase: Phase 3 - Events Catalog Service
assignee: ''
agent: "opencode"
shell_pid: "7068"
review_status: "approved"
reviewed_by: "Tiago Ramos"
history:
- timestamp: '2026-06-01T15:16:23Z'
  lane: planned
  agent: system
  shell_pid: ''
  action: Prompt generated via /spec-kitty.tasks
requirement_refs:
- FR-005
---

# Work Package Prompt: WP03 – Events Catalogue & CRUD Services

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

- Create the `EventsService` Web API project with an isolated relational database.
- Establish an EF Core 10 database context and migrations for events.
- Implement the public-facing event querying API endpoint `/api/events`.
- Build secured event CRUD APIs with strict inputs validations.

## Context & Constraints

- Refers to `kitty-specs/002-backend-microservices-synchronous-core/spec.md` (defining public catalog and organiser dashboard endpoints).
- Requires a base branch with WP02:
  ```bash
  spec-kitty implement WP03 --base WP02
  ```

## Subtasks & Detailed Guidance

### Subtask T008 - Scaffold EventsService & EF Context
- **Purpose**: Setup EventsService Web API and relational datastore context.
- **Steps**:
  1. Create `/backend/src/Services/EventsService` Web API project.
  2. Create `Data/EventsDbContext.cs` exposing an `Events` entity.
  3. Generate and apply initial database migrations using `dotnet ef migrations add InitialEventsDb`.
- **Files**:
  - `/backend/src/Services/EventsService/EventsService.csproj`
  - `/backend/src/Services/EventsService/Data/EventsDbContext.cs`
- **Parallel?**: No. Database context must exist before event records can be saved.

### Subtask T009 - Implement public catalogue queries
- **Purpose**: Query available events with high performance.
- **Steps**:
  1. Create a public, anonymous endpoint `GET /api/events` in `EventsController`:
     - Return list of all events having `AvailableTickets > 0`.
     - Support search and titles filtering directly in SQL via LINQ.
- **Files**:
  - `/backend/src/Services/EventsService/Controllers/EventsController.cs`
- **Parallel?**: Yes. Can proceed once database context (T008) is complete.

### Subtask T010 - Build secured organiser event CRUD API
- **Purpose**: Enforce role-restrictions and strict inputs validations on event mutations.
- **Steps**:
  1. Create `POST /api/organiser/events` and `PUT /api/organiser/events/{id}` endpoints.
  2. Secure with `[Authorize(Roles = "Organiser,Admin")]` attributes.
  3. Enforce strict validations:
     - Event Title must be strictly less than or equal to 100 characters.
     - Event DateTime must be configured in the future.
     - Capacity must be strictly greater than or equal to 1.
- **Files**:
  - `/backend/src/Services/EventsService/Controllers/EventsController.cs`
- **Parallel?**: Yes. Can be built concurrently with T009.

## Risks & Mitigations

- **Validation mismatch**: Ensure past dates are blocked with a clear model-state validation response (HTTP 400 Bad Request) immediately.

## Review Guidance

- Verify that event titles exceeding 100 characters trigger HTTP 400 validation errors immediately.
- Confirm anonymous traffic cannot write, edit, or delete events.

## Activity Log

- 2026-06-01T15:16:23Z – system – lane=planned – Prompt generated via /spec-kitty.tasks
- 2026-06-01T16:01:01Z – opencode – shell_pid=20364 – lane=doing – Assigned agent via workflow command
- 2026-06-01T16:25:57Z – opencode – shell_pid=20364 – lane=for_review – Ready for review: EventsService with EF Core SQLite, anonymous GET /api/events queries with real-time filters, and role-restricted CRUD operations with strict field validations.
- 2026-06-01T16:26:19Z – opencode – shell_pid=7068 – lane=doing – Started review via workflow command
- 2026-06-01T16:26:30Z – opencode – shell_pid=7068 – lane=done – Review passed: EventsService database migrations, anonymous list search query API, and validated CRUD endpoints build successfully with zero errors.
