---
work_package_id: WP03
title: Event Management Endpoints
lane: "done"
dependencies: [WP02]
base_branch: 001-event-management-restful-api-WP02
base_commit: ca777fcb2e15abab42d39d3f59f90375b31d976d
created_at: '2026-05-13T11:03:37.094062+00:00'
subtasks: [T010, T011, T012, T013]
shell_pid: "21640"
agent: "opencode"
reviewed_by: "Tiago Ramos"
review_status: "approved"
---

## Objective

Build the CRUD endpoints for managing events, with role-based access control for Organisers and Admins.

## Implementation Command

```bash
spec-kitty implement WP03 --base WP02
```

## Subtasks

### T010: Implement Event Creation Endpoint
**Purpose**: Allow Organisers and Admins to create new events.
**Steps**:
1.  Create `src/api/routers/events.py`.
2.  Implement a `POST /events/` endpoint that requires authentication.
3.  The endpoint should only be accessible to users with the 'organiser' or 'admin' role.
4.  The `organiser_id` should be set to the current authenticated user's ID.

### T011: Implement Event Retrieval Endpoints
**Purpose**: Allow users to view events.
**Steps**:
1.  In `src/api/routers/events.py`, implement a `GET /events/` endpoint to list all `published` events. This should be a public endpoint.
2.  Implement a `GET /events/{event_id}/` endpoint to retrieve a single event by its ID.

### T012: Implement Event Update Endpoint
**Purpose**: Allow the owner of an event or an Admin to update it.
**Steps**:
1.  In `src/api/routers/events.py`, implement a `PUT /events/{event_id}/` endpoint.
2.  Add logic to ensure that only the user who created the event (the organiser) or an admin can perform the update.

### T013: Implement Event Deletion Endpoint
**Purpose**: Allow the owner of an event or an Admin to delete it.
**Steps**:
1.  In `src/api/routers/events.py`, implement a `DELETE /events/{event_id}/` endpoint.
2.  Add logic to ensure that only the event's organiser or an admin can delete it.

## Definition of Done
- The `POST /events/` endpoint is functional and restricted to Organisers and Admins.
- The `GET /events/` and `GET /events/{event_id}/` endpoints are functional.
- The `PUT /events/{event_id}/` endpoint is functional with correct ownership checks.
- The `DELETE /events/{event_id}/` endpoint is functional with correct ownership checks.

## Activity Log

- 2026-05-13T11:03:40Z – opencode – shell_pid=3860 – lane=doing – Assigned agent via workflow command
- 2026-05-13T11:04:41Z – opencode – shell_pid=3860 – lane=for_review – Ready for review
- 2026-05-13T11:13:44Z – opencode – shell_pid=21640 – lane=doing – Started review via workflow command
- 2026-05-13T11:13:53Z – opencode – shell_pid=21640 – lane=done – Review passed
