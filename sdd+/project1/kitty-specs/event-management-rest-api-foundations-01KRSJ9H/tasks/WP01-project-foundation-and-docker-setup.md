---
work_package_id: WP01
title: Project Foundation & Docker Setup
dependencies: []
requirement_refs:
- FR-004
- FR-009
- FR-010
planning_base_branch: feature/sdd-plus-project1
merge_target_branch: feature/sdd-plus-project1
branch_strategy: Planning artifacts for this feature were generated on feature/sdd-plus-project1. During /spec-kitty.implement this WP may branch from a dependency-specific base, but completed changes must merge back into feature/sdd-plus-project1 unless the human explicitly redirects the landing branch.
base_branch: kitty/mission-event-management-rest-api-foundations-01KRSJ9H
base_commit: 3897d34972f5a7a31f97b7273d69df7d7e8d731b
created_at: '2026-05-22T13:27:57.381748+00:00'
subtasks:
- T001
- T002
- T003
shell_pid: "4452"
history: []
authoritative_surface: cmd/
execution_mode: code_change
owned_files:
- cmd/api/main.go
- internal/config/
- go.mod
- go.sum
- Dockerfile
- docker-compose.yml
tags: []
agent: "opencode"
---

## 1. Objective

Initialize the Go project, set up the basic server structure, and configure Docker for a reproducible local development environment.

## 2. Detailed Guidance

### Subtask T001: Set up initial Go project structure
- Create the standard Go project layout: `/cmd/api`, `/internal/config`, `/internal/api`, `/internal/domain`, `/internal/platform`.
- Run `go mod init <your_module_path>` to initialize the Go module.

### Subtask T002: Implement basic server setup and configuration loading
- In `cmd/api/main.go`, create a basic HTTP server.
- In `internal/config`, implement logic to load configuration from environment variables (e.g., for database connection strings, server port).

### Subtask T003: Set up Docker and `docker-compose.yml`
- Create a `Dockerfile` for the Go application. It should be a multi-stage build to create a small, optimized final image.
- Create a `docker-compose.yml` file that defines two services: `api` (your Go app) and `db` (PostgreSQL).
- The `api` service should build from the `Dockerfile`.
- The `db` service should use the official `postgres` image and configure a volume for data persistence.

## 3. Definition of Done
- The project structure is created.
- `docker-compose up` successfully builds and starts the Go API and the PostgreSQL database.
- The API server starts and listens on the configured port.

## Activity Log

- 2026-05-22T13:27:59Z – opencode – shell_pid=15380 – Assigned agent via action command
- 2026-05-22T13:30:21Z – opencode – shell_pid=15380 – Ready for review
- 2026-05-22T13:32:21Z – opencode – shell_pid=4452 – Started review via action command
- 2026-05-22T13:32:37Z – opencode – shell_pid=4452 – Review passed
