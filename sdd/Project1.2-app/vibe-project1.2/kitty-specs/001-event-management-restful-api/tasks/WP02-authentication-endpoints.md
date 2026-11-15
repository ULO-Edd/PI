---
work_package_id: WP02
title: Authentication Endpoints
lane: "done"
dependencies: [WP01]
base_branch: 001-event-management-restful-api-WP01
base_commit: 4ec3713ed42044b75535212f16b03de1bf376864
created_at: '2026-05-13T11:01:16.960658+00:00'
subtasks: [T006, T007, T008, T009]
shell_pid: "6804"
agent: "opencode"
reviewed_by: "Tiago Ramos"
review_status: "approved"
---

## Objective

Implement user registration and authentication using JWT. This work package provides the security foundation for the API.

## Implementation Command

```bash
spec-kitty implement WP02 --base WP01
```

## Subtasks

### T006: Implement Password Hashing
**Purpose**: Securely store user passwords.
**Steps**:
1.  Create `src/core/security.py`.
2.  Add functions to hash a plain-text password using `passlib` and to verify a password against a hash.

### T007: Implement JWT Token Logic
**Purpose**: Create and manage JWT tokens for authenticated sessions.
**Steps**:
1.  In `src/core/security.py`, add functions to:
    -   Create a JWT access token containing the user's identity.
    -   Decode and validate a JWT token.
2.  Use the `python-jose` library for JWT operations.

### T008: Create User Registration Endpoint
**Purpose**: Allow new users to register.
**Steps**:
1.  Create `src/api/routers/users.py`.
2.  Implement a `POST /users/` endpoint that takes a `UserCreate` schema.
3.  Hash the user's password before storing it in the database.
4.  Return the newly created user's data.

### T009: Create Authentication Endpoint
**Purpose**: Allow users to log in and receive a JWT token.
**Steps**:
1.  Create `src/api/routers/auth.py`.
2.  Implement a `POST /token` endpoint that accepts a username and password via an OAuth2-compatible form.
3.  Verify the credentials, and if they are valid, create and return a JWT token.

## Definition of Done
- Password hashing and verification functions are implemented.
- JWT creation and decoding functions are implemented.
- The user registration endpoint is functional and correctly stores new users.
- The login endpoint is functional and returns a valid JWT token upon successful authentication.

## Activity Log

- 2026-05-13T11:01:20Z – opencode – shell_pid=12736 – lane=doing – Assigned agent via workflow command
- 2026-05-13T11:02:06Z – opencode – shell_pid=12736 – lane=for_review – Ready for review
- 2026-05-13T11:13:11Z – opencode – shell_pid=6804 – lane=doing – Started review via workflow command
- 2026-05-13T11:13:34Z – opencode – shell_pid=6804 – lane=done – Review passed
