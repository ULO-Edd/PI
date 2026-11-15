---
work_package_id: WP02
title: Identity Service and Role Contracts
lane: "done"
dependencies: [WP01]
base_branch: 002-backend-microservices-synchronous-core-WP01
base_commit: 69d72721268c2993c1b145aa5105bc43cd59f5b8
created_at: '2026-06-01T15:47:15.553445+00:00'
subtasks:
- T005
- T006
- T007
phase: Phase 2 - Core Identity & Authorization
assignee: ''
agent: "opencode"
shell_pid: "39360"
review_status: "approved"
reviewed_by: "Tiago Ramos"
history:
- timestamp: '2026-06-01T15:16:23Z'
  lane: planned
  agent: system
  shell_pid: ''
  action: Prompt generated via /spec-kitty.tasks
requirement_refs:
- FR-004
---

# Work Package Prompt: WP02 – Identity Service and Role Contracts

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

- Create the `IdentityService` Web API project with an isolated relational database.
- Establish an EF Core 10 database context and migrations for users.
- Implement user password registration using cryptographic hashing (bcrypt).
- Build the `/api/auth/login` endpoint to issue cryptographically signed JWT tokens containing role claims.

## Context & Constraints

- Refers to `kitty-specs/002-backend-microservices-synchronous-core/spec.md` (defining user roles and secure login).
- Requires a base branch with WP01:
  ```bash
  spec-kitty implement WP02 --base WP01
  ```

## Subtasks & Detailed Guidance

### Subtask T005 - Scaffold IdentityService & EF Context
- **Purpose**: Setup IdentityService Web API and relational datastore context.
- **Steps**:
  1. Create `/backend/src/Services/IdentityService` Web API project.
  2. Install `Microsoft.EntityFrameworkCore.Design` and SQLite/PostgreSQL relational providers.
  3. Create `Data/IdentityDbContext.cs` exposing a `DbSet<User>` table.
  4. Generate and apply initial database migrations using `dotnet ef migrations add InitialIdentityDb`.
- **Files**:
  - `/backend/src/Services/IdentityService/IdentityService.csproj`
  - `/backend/src/Services/IdentityService/Data/IdentityDbContext.cs`
- **Parallel?**: No. Database context must exist before user registrations.

### Subtask T006 - Implement credentials registration with bcrypt
- **Purpose**: Securely hash and store user credentials during registration.
- **Steps**:
  1. Add `BCrypt.Net-Next` package to `IdentityService`.
  2. Implement an `/api/auth/register` endpoint or registration controller:
     - Input: `{ email, password, role }`
     - Validate unique email address.
     - Hash the incoming password: `BCrypt.HashPassword(password)`.
     - Save user entity to database.
- **Files**:
  - `/backend/src/Services/IdentityService/Controllers/AuthController.cs`
- **Parallel?**: Yes. Can proceed once DB context is complete.

### Subtask T007 - Build login claims endpoint
- **Purpose**: Authenticate credentials and return signed JWT role tokens.
- **Steps**:
  1. Create `/api/auth/login` endpoint:
     - Input: `{ email, password }`
     - Fetch user, verify hash: `BCrypt.Verify(password, user.PasswordHash)`.
  2. Generate a JWT token containing:
     - `ClaimTypes.NameIdentifier` (User ID)
     - `ClaimTypes.Email` (User Email)
     - `ClaimTypes.Role` (User Role claim: Admin, Organiser, or Attendee)
  3. Sign token using SHA256 and return secure session cookies or response payloads.
- **Files**:
  - `/backend/src/Services/IdentityService/Controllers/AuthController.cs`
- **Parallel?**: Yes. Can run concurrently with T006.

## Risks & Mitigations

- **Timing Attacks**: Ensure that invalid login emails take approximately the same time to process as invalid passwords by executing dummy hashes when no user is located.
- **Weak Role Validations**: Limit roles strictly to the verified set (Admin, Organiser, Attendee). Enforce this constraint directly in registration schema validations.

## Review Guidance

- Verify that passwords are never stored in plain-text inside the Users table.
- Decrypt a test login token and verify it contains the correct Role claim.

## Activity Log

- 2026-06-01T15:16:23Z – system – lane=planned – Prompt generated via /spec-kitty.tasks
- 2026-06-01T15:47:18Z – opencode – shell_pid=40932 – lane=doing – Assigned agent via workflow command
- 2026-06-01T15:58:11Z – opencode – shell_pid=40932 – lane=for_review – Ready for review: IdentityService with EF Core SQLite, user registration with BCrypt hashing, and JWT login session cookies/endpoints.
- 2026-06-01T15:58:20Z – opencode – shell_pid=39360 – lane=doing – Started review via workflow command
- 2026-06-01T15:59:03Z – opencode – shell_pid=39360 – lane=done – Review passed: IdentityService database migrations, credentials BCrypt hashing, and token issuance compile and build successfully with zero errors.
