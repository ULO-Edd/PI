# Feature Specification: Backend Microservices Synchronous Core

**Feature Branch**: `002-backend-microservices-synchronous-core`  
**Created**: 2026-06-01  
**Status**: Draft  
**Input**: User description: "Generate the specification for the 'Backend Microservices (The Synchronous Core)'."

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Secure and Managed API Access (Priority: P1)

As an API consumer or client application, I want to route all traffic through a secure API gateway that enforces authentication and request rate limits, so that downstream core services are protected from abuse and unauthorized access.

**Why this priority**: Gateway-level security and rate limiting are foundational requirements. They protect downstream services from performance degradation and ensure only authenticated clients can reach private data.

**Independent Test**: Can be fully tested by sending unauthenticated requests (which should be blocked with an HTTP 401), authenticated requests within the rate limits (which should pass), and high-frequency requests that exceed the rate limits (which should be blocked with an HTTP 429).

**Acceptance Scenarios**:

1. **Given** an unauthenticated request to a secured dashboard endpoint, **When** the request arrives at the API Gateway, **Then** the request is blocked and returns an HTTP 401 Unauthorized status.
2. **Given** an authenticated user who has exceeded their allocated request frequency limit, **When** they make an additional request, **Then** the API Gateway returns an HTTP 429 Too Many Requests status with a header specifying when they can retry.

---

### User Story 2 - User Identity and Role Management (Priority: P2)

As a service operator, I want to manage user accounts and explicitly enforce role-based access control (Admin, Organiser, Attendee) so that operations like creating events and viewing attendee lists are limited to authorized users.

**Why this priority**: Operational security relies on role-based isolation. Attendees should not be able to delete events, and organisers should only manage their own events.

**Independent Test**: Can be fully tested by creating users with Admin, Organiser, and Attendee roles, and verifying that restricted endpoints (e.g., event creation) are only accessible to the correct roles.

**Acceptance Scenarios**:

1. **Given** an attendee user token, **When** they request to create a new event, **Then** the gateway blocks the request and returns an HTTP 403 Forbidden status.
2. **Given** an organiser user token, **When** they request to create a new event, **Then** the identity check validates their role and allows the event to be created successfully.

---

### User Story 3 - Synchronous Capacity-Validated Ticket Booking (Priority: P1)

As an attendee, I want to book a ticket to an event and receive instant confirmation, knowing that the system has synchronously validated that the event has available capacity and reserved my spot without any double-booking.

**Why this priority**: Essential for transactional integrity. Capacity validation must be synchronous and reliable to prevent overbooking popular events.

**Independent Test**: Can be fully tested by booking a ticket for an event with 1 available spot (which succeeds), and then immediately attempting another book request for the same spot (which must fail with a capacity-exhausted error).

**Acceptance Scenarios**:

1. **Given** an event has 5 available tickets, **When** a registration request is submitted, **Then** the registrations service synchronously requests the events service to deduct 1 ticket, deducts the capacity, and returns an HTTP 201 Created status with ticket details.
2. **Given** an event has 0 available tickets, **When** a registration request is submitted, **Then** the capacity validation fails, the reservation is rejected, no registration record is created, and the system returns an HTTP 400 Bad Request status.

---

### Edge Cases

- **Network Timeout During Capacity Reservation**: If the registrations service successfully contacts the events service to deduct capacity but the network drops before the registrations database transaction completes, the system must issue an automatic rollback call to restore the ticket capacity in the events database.
- **Concurrent Booking for Last Ticket**: If multiple users attempt to book the last remaining ticket at the exact same millisecond, the events service must execute an atomic deduction to ensure only one booking succeeds and the rest are rejected with a capacity error.
- **Mismatched Authentication Token**: If a request contains a valid token but the token signature is unrecognized by the API gateway, the gateway must drop the request instantly.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The API Gateway MUST act as the single entry point for all synchronous HTTP traffic, routing requests to the appropriate downstream microservice.
- **FR-002**: The API Gateway MUST validate JWT authentication tokens before routing requests to secured downstream services.
- **FR-003**: The API Gateway MUST apply rate-limiting configurations based on the client's identity and user role.
- **FR-004**: The Identity Service MUST manage user roles, enforcing permissions for three specific roles: Admin, Organiser, and Attendee.
- **FR-005**: The Events Service MUST manage event metadata, tracking overall capacity and remaining available ticket counts.
- **FR-006**: The Registrations Service MUST handle ticket bookings and registrations, and it MUST maintain a separate, isolated database from the Events Service.
- **FR-007**: No cross-database joins or direct queries between the Events database and Registrations database are permitted.
- **FR-008**: The Registrations Service MUST synchronously query the Events Service via secure HTTP to validate capacity and deduct inventory during a booking request.
- **FR-009**: If the database transaction in the Registrations Service fails, the system MUST automatically issue a rollback command to the Events Service to restore the deducted capacity.
- **FR-010**: All service endpoints MUST publish and adhere to valid OpenAPI 3.0 schemas.

### Key Entities *(include if feature involves data)*

- **User Identity**: Represents a registered system user. Key attributes: unique identifier, email, password hash, status, and role (Admin, Organiser, Attendee).
- **Event Catalog**: Represents an upcoming event with ticket inventory. Key attributes: unique identifier, title, overall capacity, allocated capacity, remaining capacity, status, and organiser reference.
- **Registration Booking**: Represents a successful ticket purchase. Key attributes: unique identifier, event reference, attendee reference, purchase timestamp, ticket status, and gateway correlation ID.

### Assumptions and Dependencies

- **Assumption 1**: Downstream microservices communicate over a highly secure, private network segment.
- **Assumption 2**: Client applications handle the HTTP status codes (such as 401, 403, 429) gracefully, displaying appropriate alerts to the end-user.
- **Dependency 1**: The client or external identity provider is capable of issuing standard, cryptographically verifiable JWTs.
- **Dependency 2**: The database engines used by both the Events and Registrations services support ACID-compliant transactions and atomic counter updates.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: Centralized routing entry points respond to incoming HTTP requests in under 200 milliseconds at p95 under standard operating load.
- **SC-002**: Synchronous capacity verification and reservation requests between domain services complete in under 500 milliseconds.
- **SC-003**: 100% of invalid or unauthenticated requests are blocked at the perimeter and return a standard HTTP 401 or 403 response.
- **SC-004**: The system isolates data access perfectly, resulting in 0 cross-database queries or direct schema references.
- **SC-005**: Rate limits are applied dynamically, allowing public attendees 100 requests per minute and registered organisers up to 1000 requests per minute.
- **SC-006**: In the event of registration transaction failures, ticket inventory allocations are rolled back in the inventory database within 1.0 seconds.
