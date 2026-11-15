# Mission Specification: Event Management REST API Foundations

**Mission ID**: `01KRSJ9HEM13RHANKZWSV5CMQK`
**Mission Type**: `software-dev`

## 1. Overview

This document specifies the requirements for the foundational components of a REST API for event management at ESTG/IPLeiria. The API will consolidate event publication, registration, and management for various user roles.

## 2. User Roles & Personas

| Role | Description |
| :--- | :--- |
| **Attendee** | Students, professors, and external visitors who browse and register for events. |
| **Organiser** | Staff, professors, or student representatives who create and manage events. |
| **Administrator** | Central platform administrators who oversee organisers, all events, and global settings. |

## 3. Functional Requirements

| ID | Requirement | Status |
| :--- | :--- | :--- |
| **FR-001** | Users must be able to authenticate using a username and password. | `pending` |
| **FR-002** | The system must distinguish between Attendee, Organiser, and Administrator roles upon authentication. | `pending` |
| **FR-003** | API endpoints for creating, updating, or deleting events must be restricted to the event's Organiser or an Administrator. | `pending` |
| **FR-004** | All users must be able to list and search for events. | `pending` |
| **FR-005** | Authenticated Attendees must be able to register for a published event. | `pending` |
| **FR-006** | Authenticated Attendees must be able to unregister from an event they are registered for. | `pending` |
| **FR-007** | Event registration must fail if the event's capacity has been reached. | `pending` |
| **FR-008** | Authenticated Organisers must be able to create, update, and delete their own events. | `pending` |
| **FR-009** | Events must have a title, description, location, start date/time, end date/time, organiser, capacity, and status. | `pending` |
| **FR-010** | Event status can be `draft`, `published`, or `cancelled`. | `pending` |
| **FR-011** | Administrators must have full CRUD access to all events, users, and registrations. | `pending` |

## 4. Non-Functional Requirements

| ID | Requirement | Metric | Status |
| :--- | :--- | :--- | :--- |
| **NFR-001** | API endpoints should respond to 95% of requests within 500ms under normal load. | `response_time < 500ms` | `pending` |
| **NFR-002**| The API must handle input validation and return meaningful HTTP status codes and error messages. | `clear_error_messages` | `pending` |

## 5. Constraints

| ID | Constraint |
| :--- | :--- |
| **C-001** | The system will use a PostgreSQL database. |
| **C-002** | The application will be deployed via Docker. |
| **C-003** | The API documentation will be generated using Swag. |

## 6. User Scenarios & Testing

### Scenario 1: Attendee registers for an event
- **Given** an Attendee is logged in
- **And** there is a published event with available capacity
- **When** the Attendee sends a request to register for the event
- **Then** the system confirms the registration
- **And** the event's registration count is incremented.

### Scenario 2: Organiser creates a new event
- **Given** an Organiser is logged in
- **When** the Organiser sends a request to create a new event with valid data
- **Then** the system creates the event with a `draft` status
- **And** associates the event with that Organiser.

## 7. Success Criteria

- The API provides all specified CRUD operations for Events, Users, and Registrations.
- Role-based access control is correctly implemented for all user types.
- The system successfully prevents registration for events that are at full capacity.
