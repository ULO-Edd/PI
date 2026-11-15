# Mission Specification: Fix Event and User Registration Core

**Mission ID**: `01KS8SHC2ETY63EWAVDBNA8PP0`
**Mission Type**: `software-dev`

## 1. Overview

This document specifies the requirements to fix missing and failing implementations from the initial event management system. Specifically, it addresses database schema migration issues that prevent table creation, and implements robust endpoints for user registration, login, event management (CRUD operations with searching/filtering), and event registration/unregistration with capacity controls.

## 2. User Roles & Personas

| Role | Description |
| :--- | :--- |
| **Attendee** | Students, professors, and external visitors who browse events and register/unregister for them. |
| **Organiser** | Staff, professors, or student representatives who create, update, and manage their own events. |
| **Administrator** | Central platform administrators who oversee organisers, all events, and global settings with full CRUD access. |

## 3. Functional Requirements

| ID | Requirement | Status |
| :--- | :--- | :--- |
| **FR-001** | Split and correct the SQL migration scripts so `000001_initial_schema.up.sql` only creates the `users`, `events`, and `registrations` tables, and `000001_initial_schema.down.sql` only drops them. | `pending` |
| **FR-002** | The user registration (`POST /register`) endpoint must persist user accounts in the `users` table, hashing passwords and handling duplicate username conflicts with a `409 Conflict` or `400 Bad Request`. | `pending` |
| **FR-003** | Endpoints must support creating (`POST /events`), updating (`PUT /events/{eventId}`), and deleting (`DELETE /events/{eventId}`) events, restricted to the associated Organiser or Administrator. | `pending` |
| **FR-004** | Endpoints must support listing (`GET /events`), searching/filtering (e.g., by location, organizer, status, or date), and retrieving (`GET /events/{eventId}`) individual events for all users. | `pending` |
| **FR-005** | Authenticated Attendees must be able to register (`POST /events/{eventId}/register`) for published events. | `pending` |
| **FR-006** | Authenticated Attendees must be able to unregister (`DELETE /events/{eventId}/register`) from events they are currently registered for. | `pending` |
| **FR-007** | Event registration must respect the event's capacity. Attempting to register for a full event must return `409 Conflict`. | `pending` |

## 4. Non-Functional Requirements

| ID | Requirement | Metric | Status |
| :--- | :--- | :--- | :--- |
| **NFR-001** | The API must return appropriate HTTP status codes (201 for creation, 200/204 for success, 400 for bad input, 401/403 for access control, 404 for missing resources, 409 for conflicts). | `http_status_correctness >= 100%` | `pending` |
| **NFR-002** | API endpoints must validate user inputs (e.g., ensuring passwords have a minimum length, roles match valid types, capacity is non-negative). | `input_validation_enforced` | `pending` |
| **NFR-003** | API endpoints should respond within 500ms for 95% of requests under standard testing load. | `latency_p95 < 500ms` | `pending` |

## 5. Constraints

| ID | Constraint |
| :--- | :--- |
| **C-001** | The application must continue using PostgreSQL as its persistent relational database. |
| **C-002** | The application environment and services must compile and run using Docker Compose. |
| **C-003** | All code changes must match the existing Go modular design and folder structure. |

## 6. User Scenarios & Testing

### Scenario 1: Corrected Schema Migration
- **Given** an uninitialized PostgreSQL database
- **When** the API service boots up and connects to the database
- **Then** the schema migration is executed successfully
- **And** the `users`, `events`, and `registrations` tables remain created and accessible in the database.

### Scenario 2: Successful User Registration
- **Given** a user registration payload with a unique username, password, and role "Attendees"
- **When** a `POST /register` request is sent
- **Then** the server hashes the password and persists the user
- **And** returns `201 Created` with the newly created user object (excluding the raw password).

### Scenario 3: Event Registration Capacity Gate
- **Given** a published event with a capacity of 1 and already has 1 registration
- **And** a different authenticated Attendee attempts to register
- **When** they send a `POST /events/{eventId}/register` request
- **Then** the server blocks the registration and returns `409 Conflict`.

## 7. Success Criteria

- Database migrations execute cleanly, and the tables are correctly persisted across restarts.
- Endpoints for `POST /register` and `POST /login` function correctly and securely.
- Full CRUD capabilities for events, including searching and filtering, are implemented and restricted based on the user's role.
- Attendee registration/unregistration endpoints are operational and respect event capacity constraints.

## 8. Assumptions & Dependencies

- **A-001**: Docker and Docker Compose are installed on the target environment.
- **A-002**: The base JWT authentication middleware library is functional and correctly configured.
- **D-001**: PostgreSQL database container must be healthy and reachable by the API container at startup.
