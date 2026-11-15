# Task Breakdown: Event Management REST API Foundations



## Subtask Index

| ID | Description | Work Package | Parallel |
| :--- | :--- | :--- | :--- |
| T001 | Set up initial Go project structure (`/cmd`, `/internal`, etc.). | WP01 | | [D] |
| T002 | Implement basic server setup and configuration loading. | WP01 | | [D] |
| T003 | Set up Docker and `docker-compose.yml` for the Go app and PostgreSQL. | WP01 | | [D] |
| T004 | Implement database connection and migration setup. | WP02 | | [D] |
| T005 | Create database schema for `users`, `events`, and `registrations`. | WP02 | | [D] |
| T006 | Create domain models/entities for User, Event, and Registration. | WP02 | | [D] |
| T019 | Write unit tests for the User service. | WP07 | [D] |
| T020 | Write unit tests for the Event service. | WP07 | [D] |
| T021 | Write unit tests for the Registration service. | WP07 | [D] |

---



## Work Package WP01 — Project Foundation & Docker Setup
- **Summary**: Initialize the Go project, set up the basic server structure, and configure Docker for local development.
- **Priority**: High
- **Dependencies**: None
- **Included Subtasks**:
  - [x] T001 Set up initial Go project structure (`/cmd`, `/internal`, etc.).
  - [x] T002 Implement basic server setup and configuration loading.
  - [x] T003 Set up Docker and `docker-compose.yml` for the Go app and PostgreSQL.

## Work Package WP02 — Database & Domain Models
- **Summary**: Establish the database connection, set up migrations, and create the core domain entities.
- **Priority**: High
- **Dependencies**: WP01
- **Included Subtasks**:
  - [x] T004 Implement database connection and migration setup.
  - [x] T005 Create database schema for `users`, `events`, and `registrations`.
  - [x] T006 Create domain models/entities for User, Event, and Registration.

## Work Package WP07 — Unit Testing
- **Summary**: Implement unit tests for all services to meet the 70% coverage requirement.
- **Priority**: Low
- **Dependencies**: WP02
- **Included Subtasks**:
  - [x] T019 Write unit tests for the User service.
  - [x] T020 Write unit tests for the Event service.
  - [x] T021 Write unit tests for the Registration service.
