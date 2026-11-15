---
work_package_id: WP01
title: Repository Layout & Backend Setup
dependencies: []
requirement_refs:
- C-001
- C-003
planning_base_branch: feature/sdd-plus-project2
merge_target_branch: feature/sdd-plus-project2
branch_strategy: Planning artifacts for this feature were generated on feature/sdd-plus-project2. During /spec-kitty.implement this WP may branch from a dependency-specific base, but completed changes must merge back into feature/sdd-plus-project2 unless the human explicitly redirects the landing branch.
base_branch: kitty/mission-event-management-web-platform-01KSV4CE
base_commit: 55dce2684f834596a70785d30b89e1dcfbd608d4
created_at: '2026-05-30T00:35:15.803290+00:00'
subtasks:
- T001
- T002
- T003
- T004
- T005
phase: Phase 1 - Repository Layout & Backend Setup
assignee: ''
agent: "opencode"
shell_pid: "8628"
history:
- timestamp: '2026-05-30T00:31:40Z'
  agent: system
  action: Prompt generated via /spec-kitty.tasks
authoritative_surface: backend/
execution_mode: code_change
owned_files:
- backend/cmd/**
- backend/go.mod
- backend/go.sum
- backend/Dockerfile
- docker-compose.yml
tags: []
---

# Work Package Prompt: WP01 – Repository Layout & Backend Setup

## Objective
Establish the core multi-component directory layout, integrate the Project 1 Go code base, configure the docker-compose orchestrator, and verify standalone execution.

---

## Subtask Details

### Subtask T001: Copy Go REST API Codebase
*   **Purpose:** Migrate Project 1 core logic into the new nested `/backend` directory inside `project2`.
*   **Steps:**
    1. Copy all source files (cmd, db, internal, go.mod, go.sum) from sibling directory `../project1` to `backend/`.
    2. Confirm package paths and go.mod modules match correctly.
*   **Files:** `backend/`
*   **Validation:** Directory structure contains expected cmd and internal folders.

### Subtask T002: Create Backend Dockerfile
*   **Purpose:** Containerize the Go backend.
*   **Steps:**
    1. Create a multi-stage Dockerfile using `golang:1.25-alpine` as build stage and `alpine:latest` as final.
    2. Expose port 8080.
*   **Files:** `backend/Dockerfile`
*   **Validation:** Docker image compiles successfully.

### Subtask T003: Configure Docker Compose
*   **Purpose:** Set up PostgreSQL database and backend service mapping.
*   **Steps:**
    1. Create `docker-compose.yml` in project root with `postgres` and `backend` services.
    2. Share environment credentials.
*   **Files:** `docker-compose.yml`
*   **Validation:** Services boot cleanly.

### Subtask T004: Automate Schema Migrations
*   **Purpose:** Run DB migrations automatically on backend launch.
*   **Steps:**
    1. Integrate the migrate utility hook in Go entrypoint.
*   **Files:** `backend/cmd/server/main.go`
*   **Validation:** Database initialized automatically.

### Subtask T005: Local Backend Verification
*   **Purpose:** Validate base API endpoints are up and responding.
*   **Steps:**
    1. Run endpoint tests and hit Swagger healthcheck page.
*   **Files:** `backend/cmd/**`
*   **Validation:** Returns HTTP 200 on healthcheck.

## Activity Log

- 2026-05-30T00:35:16Z – opencode – shell_pid=12016 – Assigned agent via action command
- 2026-05-30T00:38:12Z – opencode – shell_pid=12016 – Backend infrastructure and Docker setup finalized
- 2026-05-30T00:38:33Z – opencode – shell_pid=8628 – Started review via action command
- 2026-05-30T00:38:47Z – opencode – shell_pid=8628 – Review passed: Backend codebase copied, Dockerfile and docker-compose.yml fully configured and verified with passing unit tests.
