# Feature Specification: Event Management RESTful API

## 1. Overview

This document specifies a RESTful API for managing events at ESTG / IPLeiria. The API will serve as the backend for programmatic clients, providing endpoints for event creation, registration, and management. It is designed to be consumed by HTTP clients like curl, Postman, or automated testing frameworks.

## 2. Actors

- **Attendee**: A standard user who can browse and register for events.
- **Organiser**: A user responsible for creating, managing, and updating specific events.
- **Admin**: A privileged user with full control over all events and users within the system.

## 3. User Scenarios & Testing

### 3.1. Admin Manages Users and Events
- **Scenario**: An admin needs to ensure the platform runs smoothly.
- **Steps**:
    1. An admin authenticates to get a JWT token.
    2. The admin creates a new user and assigns them the 'Organiser' role.
    3. The admin views all events on the platform, regardless of the organiser.
    4. The admin edits an event created by another organiser to correct a mistake.
    5. The admin deletes a cancelled event.
- **Acceptance Criteria**:
    - Admins can create, read, update, and delete any user or event.
    - Admins can assign roles to users.

### 3.2. Organiser Manages Their Events
- **Scenario**: An organiser wants to create and publish a new workshop.
- **Steps**:
    1. An organiser authenticates to get a JWT token.
    2. They create a new event with a title, description, date, and capacity, initially in 'draft' status.
    3. They update the event details.
    4. Once ready, they change the event status to 'published'.
    5. They view a list of attendees registered for their event.
- **Acceptance Criteria**:
    - Organisers can create events.
    - Organisers can only view, update, or delete events they own.
    - Organisers cannot modify events created by other organisers.

### 3.3. Attendee Registers for an Event
- **Scenario**: A student wants to register for an upcoming tech talk.
- **Steps**:
    1. An attendee authenticates to get a JWT token.
    2. They retrieve a list of all 'published' events.
    3. They filter events to find one of interest.
    4. They send a request to register for the event.
    5. Later, they decide to unregister from the event.
- **Acceptance Criteria**:
    - Any authenticated user can view published events.
    - Users can register for an event if its capacity has not been reached.
    - Users receive an error if they try to register for a full event.
    - Users can unregister from an event they are registered for.

## 4. Functional Requirements

### 4.1. Authentication
- **REQ-AUTH-01**: The system must secure endpoints using JWT tokens.
- **REQ-AUTH-02**: The system must provide endpoints for user registration and login, which return a JWT token on success.
- **REQ-AUTH-03**: The system shall differentiate between Attendee, Organiser, and Admin roles based on the JWT token.

### 4.2. Event Management (Admin)
- **REQ-ADMIN-01**: Admins shall be able to create, retrieve, update, and delete any event.
- **REQ-ADMIN-02**: Admins shall be able to create other admin users.

### 4.3. Event Management (Organiser)
- **REQ-ORG-01**: Organisers shall be able to create new events.
- **REQ-ORG-02**: Organisers shall only be able to retrieve, update, or delete events they have created.
- **REQ-ORG-03**: An event must have a `title`, `description`, `location`, `start_datetime`, `end_datetime`, `organiser_id`, `capacity`, and `status` (`draft`, `published`, `cancelled`).

### 4.4. Event Discovery and Registration (Attendee)
- **REQ-ATT-01**: Any authenticated user shall be able to list and filter all `published` events.
- **REQ-ATT-02**: Authenticated users shall be able to register for an event, subject to capacity constraints.
- **REQ-ATT-03**: Authenticated users shall be able to unregister from an event.

## 5. Key Entities

- **User**:
    - `user_id` (Primary Key)
    - `username`
    - `password_hash`
    - `role` (Attendee, Organiser, Admin)
- **Event**:
    - `event_id` (Primary Key)
    - `title`
    - `description`
    - `location`
    - `start_datetime`
    - `end_datetime`
    - `organiser_id` (Foreign Key to User)
    - `capacity`
    - `status`
- **Registration**:
    - `registration_id` (Primary Key)
    - `event_id` (Foreign Key to Event)
    - `user_id` (Foreign Key to User)

## 6. Success Criteria

- **SC-1**: 100% of API endpoints are protected by role-based access control.
- **SC-2**: Event registration count never exceeds the defined capacity.
- **SC-3**: The API documentation (OpenAPI/Swagger) is complete and accurately reflects all available endpoints and data models.
- **SC-4**: Organisers are unable to perform any modification on events they do not own.

## 7. Assumptions

- **A-1**: Authentication will be handled via JWT tokens.
- **A-2**: A relational database is required for persistent storage.
- **A-3**: The client is responsible for storing and managing the JWT token.
