# Task Breakdown: Event Management RESTful API

This document outlines the work packages (WPs) required to implement the Event Management RESTful API.

## Phase 1: Foundational Setup

### WP01: Project & Database Setup
- **Summary**: Initialize the project structure, configure the database, and define the core data models. This WP is the foundation for all subsequent work.
- **Included Subtasks**: T001, T002, T003, T004, T005
- **Implementation Sketch**: Create the directory structure, set up FastAPI configuration, define SQLAlchemy models, and establish the initial database schema.
- **Dependencies**: None
- **Estimated Prompt Size**: ~450 lines
**Requirement Refs**: None (foundational)

## Phase 2: Core API Implementation

### WP02: Authentication Endpoints
- **Summary**: Implement user registration and authentication using JWT. This WP secures the API.
- **Included Subtasks**: T006, T007, T008, T009
- **Implementation Sketch**: Create endpoints for user creation and token generation. Implement password hashing and JWT logic.
- **Dependencies**: WP01
- **Estimated Prompt Size**: ~400 lines
**Requirement Refs**: REQ-AUTH-01, REQ-AUTH-02, REQ-AUTH-03

### WP03: Event Management Endpoints
- **Summary**: Build the CRUD endpoints for managing events. This work package enables Organiser and Admin roles to manage event data.
- **Included Subtasks**: T010, T011, T012, T013
- **Implementation Sketch**: Develop the API routers for creating, reading, updating, and deleting events, including role-based access control.
- **Dependencies**: WP02
- **Estimated Prompt Size**: ~500 lines
**Requirement Refs**: REQ-ADMIN-01, REQ-ORG-01, REQ-ORG-02, REQ-ORG-03

### WP04: Event Registration Endpoints
- **Summary**: Implement the functionality for attendees to register and unregister for events.
- **Included Subtasks**: T014, T015, T016, T017
- **Implementation Sketch**: Create endpoints for event registration, ensuring that capacity constraints are respected. Add logic for unregistering.
- **Dependencies**: WP02
- **Estimated Prompt Size**: ~400 lines
**Requirement Refs**: REQ-ATT-01, REQ-ATT-02, REQ-ATT-03

## Subtask List

- **T001**: [P] Set up the main project directory structure as defined in `plan.md`.
- **T002**: [P] Create the configuration file (`src/core/config.py`) to manage environment variables.
- **T003**: [P] Define the SQLAlchemy ORM models in `src/db/models.py` for User, Event, and Registration.
- **T004**: [P] Implement the database connection and session management in `src/db/database.py`.
- **T005**: [P] Define Pydantic schemas for data validation and serialization in the `src/schemas/` directory.
- **T006**: Implement password hashing and verification logic in `src/core/security.py`.
- **T007**: Implement JWT token creation and decoding functions in `src/core/security.py`.
- **T008**: Create the user registration endpoint (`POST /users/`) in `src/api/routers/users.py`.
- **T009**: Create the authentication endpoint (`POST /token`) to return a JWT token in `src/api/routers/auth.py`.
- **T010**: Implement the endpoint to create events (`POST /events/`) with access control for Organisers and Admins.
- **T011**: Implement the endpoints to retrieve events (`GET /events/` and `GET /events/{event_id}/`).
- **T012**: Implement the endpoint to update events (`PUT /events/{event_id}/`) with ownership checks for Organisers.
- **T013**: Implement the endpoint to delete events (`DELETE /events/{event_id}/`) with ownership checks for Organisers.
- **T014**: Create the endpoint for users to register for an event (`POST /events/{event_id}/register`).
- **T015**: Implement logic to check event capacity before allowing registration.
- **T016**: Create the endpoint for users to unregister from an event (`DELETE /events/{event_id}/unregister`).
- **T017**: Implement dependency injection for authentication and database sessions in API routers.
