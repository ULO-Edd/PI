# Research Document: Backend Microservices Synchronous Core

This research document analyzes key technical decisions, architectures, and alternatives for the backend microservices.

## 1. Database Isolation and Migration Management

- **Decision**: Use separate relational database instances per microservice (Identity, Events, Registrations), managed independently via Entity Framework Core (EF Core 10) Code-First migrations.
- **Rationale**: Isolating databases per service guarantees strict bounded context isolation. Services cannot access each other's data stores directly, which prevents unintended schema dependencies and enables autonomous deployment and database scaling. EF Core Code-First migrations provide an excellent, strictly typed history of database changes.
- **Alternatives Considered**:
  - **Shared Database with Separate Tables**: Rejected because shared datastores violate the non-negotiable isolated relational databases rule in the project constitution.
  - **Dapper with Raw SQL Migrations**: Rejected. While high-performing, EF Core 10 offers extremely mature LINQ-to-SQL translation, change-tracking, and robust automated migrations which reduce administrative overhead significantly.

## 2. Synchronous Capacity Validation Strategy

- **Decision**: Option A (The Registrations service performs a synchronous HTTP call to the Events service to verify and deduct ticket inventory, with an automatic transactional rollback on failure).
- **Rationale**: Synchronous HTTP calls ensure strong transactional consistency at the moment of ticket booking. Before a booking is saved, capacity is reserved in the Events service. If the registration fails to save in the Registrations database, an automatic rollback call is triggered to restore capacity, preventing overbooking.
- **Alternatives Considered**:
  - **Read-replica Sync (Option B)**: Rejected. Under high-traffic event bursts (e.g., 500+ concurrent scans), eventually consistent replicas introduce high risks of overbooking tickets.
  - **Event-Driven Capacity Checking**: Rejected for synchronous flows. Booking confirmation must be immediate and guaranteed. Asynchronous checks would make the booking experience complex and slow for attendees.

## 3. Perimeter Gateway Security and Traffic Controls

- **Decision**: Centralized API Gateway implementing JwtBearer authentication guards, path routing, and role-based rate-limiting.
- **Rationale**: Centralizing routing and perimeter security at the gateway ensures all requests are validated before reaching downstream services. It reduces the authentication overhead on individual services. Rate limits can be dynamically throttled based on client roles extracted from decrypted JWT claims.
- **Alternatives Considered**:
  - **Direct Client-to-Service Communication**: Rejected. Exposing all microservice ports directly to the internet increases the attack surface, complicates CORS configuration, and duplicates JWT validation logic across every microservice.
