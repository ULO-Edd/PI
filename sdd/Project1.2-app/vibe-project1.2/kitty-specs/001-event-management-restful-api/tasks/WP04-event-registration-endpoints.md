---
work_package_id: WP04
title: Event Registration Endpoints
lane: "done"
dependencies: [WP02]
base_branch: 001-event-management-restful-api-WP02
base_commit: ca777fcb2e15abab42d39d3f59f90375b31d976d
created_at: '2026-05-13T11:05:44.993377+00:00'
subtasks: [T014, T015, T016, T017]
shell_pid: "22296"
agent: "opencode"
reviewed_by: "Tiago Ramos"
review_status: "approved"
---

## Objective

Implement the functionality for attendees to register and unregister for events, including capacity checks and required authentication.

## Implementation Command

```bash
spec-kitty implement WP04 --base WP02
```

## Subtasks

### T014: Create Event Registration Endpoint
**Purpose**: Allow authenticated users to register for an event.
**Steps**:
1.  In `src/api/routers/events.py`, create a `POST /events/{event_id}/register` endpoint.
2.  This endpoint must require authentication.
3.  Add the authenticated user to the event's list of registered attendees.

### T015: Implement Capacity Check
**Purpose**: Prevent registration for events that are full.
**Steps**:
1.  Before creating a new registration, query the database to count the number of existing registrations for the event.
2.  If the count is equal to or greater than the event's `capacity`, return an HTTP 400 error.

### T016: Create Event Unregistration Endpoint
**Purpose**: Allow authenticated users to unregister from an event.
**Steps**:
1.  In `src/api/routers/events.py`, create a `DELETE /events/{event_id}/unregister` endpoint.
2.  This endpoint must require authentication.
3.  Remove the authenticated user from the event's list of registered attendees.

### T017: Implement Dependency Injection
**Purpose**: Provide dependencies like the database session and current user to the endpoint functions.
**Steps**:
1.  Create `src/api/dependencies.py`.
2.  Implement a dependency function `get_current_user` that decodes the JWT token from the request and retrieves the user from the database.
3.  Apply the `get_db` and `get_current_user` dependencies to all protected endpoints.

## Definition of Done
- The `POST /events/{event_id}/register` endpoint is functional and respects event capacity.
- The `DELETE /events/{event_id}/unregister` endpoint is functional.
- User authentication is correctly enforced on all registration endpoints.
- Dependencies for database sessions and user authentication are correctly injected.

## Activity Log

- 2026-05-13T11:05:48Z – opencode – shell_pid=13424 – lane=doing – Assigned agent via workflow command
- 2026-05-13T11:06:25Z – opencode – shell_pid=13424 – lane=for_review – Ready for review
- 2026-05-13T11:14:02Z – opencode – shell_pid=22296 – lane=doing – Started review via workflow command
- 2026-05-13T11:14:12Z – opencode – shell_pid=22296 – lane=done – Review passed
