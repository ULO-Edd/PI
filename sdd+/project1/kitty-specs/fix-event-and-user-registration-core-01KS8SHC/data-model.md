# Data Model: Fix Event and User Registration Core

**Mission ID**: `01KS8SHC2ETY63EWAVDBNA8PP0`

This document outlines the entity definitions, database schemas, validation rules, and state transitions for the Event Management REST API.

---

## 1. Entity Definitions

### 1.1 User Entity
Represents an authenticated user within the system.

- **id** (UUID, Primary Key): Unique identifier.
- **username** (VARCHAR, Unique, Indexed): User's unique handle.
- **password_hash** (VARCHAR): Securely hashed password.
- **role** (VARCHAR): One of `attendee`, `organiser`, or `administrator`.
- **created_at** (TIMESTAMPTZ): Creation timestamp.

### 1.2 Event Entity
Represents an event created by an Organiser.

- **id** (UUID, Primary Key): Unique identifier.
- **title** (VARCHAR): Event title.
- **description** (TEXT): Detailed description of the event.
- **location** (VARCHAR): Event venue or link.
- **start_time** (TIMESTAMPTZ): Start timestamp.
- **end_time** (TIMESTAMPTZ): End timestamp.
- **organiser_id** (UUID, Foreign Key → `users.id`): Creator/organiser of the event.
- **capacity** (INTEGER): Maximum number of registered attendees allowed.
- **status** (VARCHAR): One of `draft`, `published`, or `cancelled`.
- **created_at** (TIMESTAMPTZ): Creation timestamp.

### 1.3 Registration Entity
Represents a registration mapping between an Attendee (User) and an Event.

- **id** (UUID, Primary Key): Unique identifier.
- **event_id** (UUID, Foreign Key → `events.id`): Registered event.
- **attendee_id** (UUID, Foreign Key → `users.id`): Registered attendee.
- **created_at** (TIMESTAMPTZ): Registration timestamp.

---

## 2. Validation & Rules

### User Validation
- **Username**: Must be non-empty, alphanumeric, and at least 3 characters. Must be unique in the database.
- **Password**: Must be at least 6 characters before hashing.
- **Role**: Must be checked against canonical values: `attendee`, `organiser`, `administrator` (case-insensitive during request parsing, normalized to lowercase).

### Event Validation
- **Title**: Must be non-empty and less than 255 characters.
- **Times**: `end_time` must be strictly after `start_time`.
- **Capacity**: Must be a positive integer (> 0).
- **Status**: Must be `draft`, `published`, or `cancelled`.

### Registration Validation
- **Status Gate**: Registrations are only allowed on events with status `published`. Registrations on `draft` or `cancelled` events must fail with `400 Bad Request`.
- **Duplicate Gate**: An attendee cannot register for the same event multiple times (`409 Conflict`).
- **Capacity Gate**: The total count of active registrations for an event must be less than its defined `capacity` (`409 Conflict`).

---

## 3. State Transitions

### Event Status Transitions
- **Draft**: Initial state of any newly created event.
- **Published**: Transitioned from `Draft`. This state allows Attendees to register.
- **Cancelled**: Transitioned from `Draft` or `Published`. Registrations are blocked and existing registrations should be considered void.

```
       [ Create ]
           │
           ▼
     ┌───────────┐
     │   Draft   │◄───────┐
     └─────┬─────┘        │ (Update)
           │              │
     [ Publish ]          │
           │              │
           ▼              │
     ┌───────────┐────────┘
     │ Published │
     └─────┬─────┘
           │
      [ Cancel ]
           │
           ▼
     ┌───────────┐
     │ Cancelled │
     └───────────┘
```
