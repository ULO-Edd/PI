---
work_package_id: WP05
title: API Gateway & DLQ Admin Endpoints
lane: "done"
dependencies: []
base_branch: main
base_commit: 1a3d908567c492d11a489fd6cbdf1a4c2a34c378
created_at: '2026-06-02T16:12:09.583488+00:00'
subtasks: [T022, T023, T024]
requirement_refs:
- FR-011
- FR-012
shell_pid: "45020"
agent: "opencode"
reviewed_by: "Tiago Ramos"
review_status: "approved"
---

# WP05: API Gateway & DLQ Admin Endpoints

## Objective
Expose the Dead-Letter management functionality securely through the perimeter API Gateway.

## Context
The `DeadLetterService` is an internal microservice. Admins must access its management endpoints via the API Gateway, which handles authentication and role-based authorization.

## Detailed Guidance

### Subtask T022: Configure Gateway routing to `DeadLetterService`
**Purpose**: Route external requests to the internal DLQ service.
**Steps**:
1. Update the API Gateway routing configuration:
   - Define a route for `/api/admin/dlq/**`.
   - Point this route to the internal `DeadLetterService` URL.
   - Ensure the gateway forwards the original HTTP headers.
**Validation**:
- [ ] Request to `/api/admin/dlq` is correctly forwarded to the internal service.

### Subtask T023: Implement JWT role check for DLQ endpoints (Admin access)
**Purpose**: Prevent non-admins from inspecting or manipulating failed events.
**Steps**:
1. In the API Gateway's authorization middleware:
   - Apply a `[Authorize(Roles = "Admin")]` attribute or equivalent filter to the DLQ routes.
   - Verify the JWT contains the required `Admin` claim.
   - Return `403 Forbidden` if the user is authenticated but not an Admin.
**Validation**:
- [ ] Admin user can access endpoints.
- [ ] Attendee/Organiser user receives 403 Forbidden.
- [ ] Unauthenticated user receives 401 Unauthorized.

### Subtask T024: Integrate gateway with DLQ management endpoints
**Purpose**: Ensure end-to-end connectivity for all administrative actions.
**Steps**:
1. Test the full flow through the Gateway for:
   - GET `/api/admin/dlq` (List)
   - GET `/api/admin/dlq/{id}` (Inspect)
   - POST `/api/admin/dlq/{id}/replay` (Replay)
   - DELETE `/api/admin/dlq/{id}/purge` (Purge)
**Validation**:
- [ ] All admin operations are successful when called via the Gateway.

## Definition of Done
- [ ] API Gateway routes requests to `DeadLetterService`.
- [ ] Strict `Admin` role authorization is enforced for all DLQ routes.
- [ ] End-to-end connectivity for all 4 management endpoints is verified.

## Implementation Command
`spec-kitty implement WP05 --base WP02`

## Activity Log

- 2026-06-02T16:12:12Z – opencode – shell_pid=35668 – lane=doing – Assigned agent via workflow command
- 2026-06-02T16:13:48Z – opencode – shell_pid=35668 – lane=for_review – Ready for review: implemented API Gateway proxy to DeadLetterService with strict Admin role authorization
- 2026-06-02T17:19:57Z – opencode – shell_pid=45020 – lane=doing – Started review via workflow command
- 2026-06-02T17:20:15Z – opencode – shell_pid=45020 – lane=done – Review passed: Implementation successfully covers all subtasks (T022-T024), with correct API Gateway proxying, strict Admin role authorization, and endpoint integration.
