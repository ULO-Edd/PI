---
work_package_id: WP01
title: Project & Database Setup
lane: "done"
dependencies: []
base_branch: main
base_commit: a3327819048a33e89822ff7e9f5cfe48202b84fb
created_at: '2026-05-13T10:57:53.267527+00:00'
subtasks: [T001, T002, T003, T004, T005]
shell_pid: "23104"
agent: "opencode"
reviewed_by: "Tiago Ramos"
review_status: "approved"
---

## Objective

Initialize the project structure, configure the database, and define the core data models and schemas. This work package lays the foundation for the entire application.

## Implementation Command

```bash
spec-kitty implement WP01
```

## Subtasks

### T001: Set up Project Directory Structure
**Purpose**: Create the standardized directory layout for the application.
**Steps**:
1.  Create the directories as specified in the `plan.md` file:
    ```
    src/
    ├── api/
    │   └── routers/
    ├── core/
    ├── db/
    ├── schemas/
    └── services/
    tests/
    ├── test_api/
    └── test_services/
    ```
2.  Create `__init__.py` files in each directory to mark them as Python packages.

### T002: Create Configuration File
**Purpose**: Centralize application settings.
**Steps**:
1.  Create `src/core/config.py`.
2.  Use Pydantic's `BaseSettings` to define and load configuration from environment variables (e.g., `DATABASE_URL`, `SECRET_KEY`).

### T003: Define SQLAlchemy ORM Models
**Purpose**: Create the database table structures using SQLAlchemy's ORM.
**Steps**:
1.  Create `src/db/models.py`.
2.  Define the `User`, `Event`, and `Registration` models with columns and relationships as specified in `data-model.md`.

### T004: Implement Database Connection
**Purpose**: Manage database sessions for the application.
**Steps**:
1.  Create `src/db/database.py`.
2.  Set up the SQLAlchemy engine and `sessionmaker`.
3.  Create a dependency function (`get_db`) to provide a database session to API endpoints.

### T005: Define Pydantic Schemas
**Purpose**: Create data shapes for API validation and serialization.
**Steps**:
1.  Create `src/schemas/user.py`, `src/schemas/event.py`, and `src/schemas/token.py`.
2.  Define Pydantic models for creating and reading users and events (e.g., `UserCreate`, `User`, `EventCreate`, `Event`).
3.  Define a schema for the JWT token (`Token`).

## Definition of Done
- The project directory structure is in place.
- Configuration is loaded from environment variables.
- SQLAlchemy models are defined for all entities.
- A database session dependency is available.
- Pydantic schemas for all API data shapes are created.

## Activity Log

- 2026-05-13T10:57:55Z – opencode – shell_pid=1332 – lane=doing – Assigned agent via workflow command
- 2026-05-13T11:21:19Z – opencode – shell_pid=1332 – lane=for_review – Ready for review
- 2026-05-13T11:22:24Z – opencode – shell_pid=23104 – lane=doing – Started review via workflow command
- 2026-05-13T11:22:45Z – opencode – shell_pid=23104 – lane=done – Review passed
