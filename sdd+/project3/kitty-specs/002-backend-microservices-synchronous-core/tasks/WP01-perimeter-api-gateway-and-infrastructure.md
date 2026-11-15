---
work_package_id: WP01
title: Perimeter API Gateway and Infrastructure
lane: "done"
dependencies: []
base_branch: main
base_commit: 228d0c7c78a5b66bacda226ad6de44f21bb8f440
created_at: '2026-06-01T15:21:36.402933+00:00'
subtasks:
- T001
- T002
- T003
- T004
phase: Phase 1 - Solution Scaffolding & Perimeter Setup
assignee: ''
agent: "opencode"
shell_pid: "38940"
review_status: "approved"
reviewed_by: "Tiago Ramos"
history:
- timestamp: '2026-06-01T15:16:23Z'
  lane: planned
  agent: system
  shell_pid: ''
  action: Prompt generated via /spec-kitty.tasks
requirement_refs:
- FR-001
- FR-002
- FR-003
---

# Work Package Prompt: WP01 - Perimeter API Gateway and Infrastructure

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

- Scaffold the .NET 10 solution and backend project layout under the `/backend` folder.
- Configure the centralized ApiGateway path routing and proxy rules.
- Set up JWT validation guards on secured routes.
- Implement role-based rate-limiting controls to prevent service abuse.

## Context & Constraints

- Refers to `.kittify/memory/constitution.md` (defining .NET 10).
- Refers to `kitty-specs/002-backend-microservices-synchronous-core/plan.md`.
- Enforces secure perimeter security before forwarding request traffic downstreams.

## Subtasks & Detailed Guidance

### Subtask T001 - Scaffold solution structure
- **Purpose**: Create a clean .NET 10 solution workspace with correct folder structures.
- **Steps**:
  1. Create folder `/backend` at the repository root.
  2. In `/backend`, run `dotnet new sln -n SDDProject3Backend`.
  3. Create an empty Web API project for `ApiGateway` and class library for `Shared` contracts.
  4. Register all projects within the solution file.
- **Files**:
  - `/backend/SDDProject3Backend.sln`
  - `/backend/src/ApiGateway/ApiGateway.csproj`
- **Parallel?**: No. Solution must exist first.

### Subtask T002 - Configure API Gateway routing
- **Purpose**: Map ingress routing paths to downstream microservices.
- **Steps**:
  1. Set up route forwarding in `ApiGateway/Program.cs` or using basic YARP/Ocelot routing maps.
  2. Define forward rules mapping:
     - `/api/auth/*` -> `http://localhost:5001` (Identity Service)
     - `/api/events/*` -> `http://localhost:5002` (Events Service)
     - `/api/registrations/*` -> `http://localhost:5003` (Registrations Service)
  3. Configure launch settings with fixed host port `5000`.
- **Files**:
  - `/backend/src/ApiGateway/Program.cs`
  - `/backend/src/ApiGateway/appsettings.json`
- **Parallel?**: Yes. Can proceed once scaffolding (T001) is complete.

### Subtask T003 - Configure JWT validation guards
- **Purpose**: Intercept incoming secure requests and validate token authenticity.
- **Steps**:
  1. Add `Microsoft.AspNetCore.Authentication.JwtBearer` (v10) package to the gateway.
  2. In `Program.cs`, add JwtBearer authentication middleware.
  3. Configure parameters: `ValidateIssuer`, `ValidateAudience`, `ValidateLifetime`, and set `IssuerSigningKey` matching environment secret values.
- **Files**:
  - `/backend/src/ApiGateway/Program.cs`
- **Parallel?**: Yes. Can run alongside T002.

### Subtask T004 - Implement role-based rate limits
- **Purpose**: Throttles inbound requests dynamically based on user identity or client IP.
- **Steps**:
  1. In `Program.cs`, configure the native ASP.NET Core `AddRateLimiter` middleware.
  2. Implement two distinct token bucket or sliding window policies:
     - `AttendeePolicy`: Capped at 100 requests per minute per IP or authenticated ID.
     - `OrganiserPolicy`: Capped at 1000 requests per minute.
  3. Apply the rate limiter globally or map explicitly to respective forward paths.
- **Files**:
  - `/backend/src/ApiGateway/Program.cs`
- **Parallel?**: Yes. Can run concurrently once base routing is completed.

## Risks & Mitigations

- **Signature Key exposure**: Never hardcode signature keys in `appsettings.json`. Load security keys securely via environment configurations or user-secrets during local runs.

## Review Guidance

- Verify that unauthenticated requests to protected paths return HTTP 401 Unauthorized at the gateway level.
- Assert rate-limiting triggers HTTP 429 Too Many Requests once thresholds are exceeded.

## Activity Log

- 2026-06-01T15:16:23Z – system – lane=planned – Prompt generated via /spec-kitty.tasks
- 2026-06-01T15:21:39Z – opencode – shell_pid=35648 – lane=doing – Assigned agent via workflow command
- 2026-06-01T15:43:50Z – opencode – shell_pid=35648 – lane=for_review – Ready for review: ApiGateway solution, path routing proxy rules, Jwt authentication validation guards, and dynamic role-based partitioned rate-limiting.
- 2026-06-01T15:45:02Z – opencode – shell_pid=38940 – lane=doing – Started review via workflow command
- 2026-06-01T15:45:13Z – opencode – shell_pid=38940 – lane=done – Review passed: Scaffolding, perimeter routing proxies, Jwt validation, and partitioned rate limits build successfully with zero errors.
