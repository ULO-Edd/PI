# Relational Data Model: Event Management Web Platform
*Path: [kitty-specs/event-management-web-platform-01KSV4CE/data-model.md](data-model.md)*

This document defines the relational database schema, entity structures, attributes, and constraints required to support the platform.

---

## 1. Entity-Relationship Outline

```
  ┌───────────────┐          ┌────────────────┐
  │     USERS     │1       * │     EVENTS     │
  │ (Organisers/  ├─────────►│(Draft/Pub/Can) │
  │  Attendees)   │          └──────┬─────────┘
  └──────┬────────┘                 │1
         │1                         │
         │                          │*
         │                          ▼
         │                  ┌────────────────┐
         │                  │  REGISTRATIONS │
         │1                *│ (Registered/   │
         └─────────────────►│ Present/NoShow)│
                            └────────────────┘
```

---

## 2. Table Schemas

### Table: `users`
Represents all system actors.

| Field Name | Data Type | Constraints | Description |
|---|---|---|---|
| `id` | `UUID` | Primary Key, Default: `gen_random_uuid()` | Unique user identifier. |
| `name` | `VARCHAR(100)` | NOT NULL | User's full name. |
| `email` | `VARCHAR(255)` | NOT NULL, UNIQUE | User's email address (for login and notifications). |
| `password_hash` | `VARCHAR(255)` | NOT NULL | Bcrypt hashed password. |
| `role` | `VARCHAR(20)` | NOT NULL | One of: `Attendee`, `Organizer`, `Administrator`. |
| `status` | `VARCHAR(20)` | NOT NULL, Default: `Pending` | One of: `Pending` (for organizers), `Approved`, `Suspended`. |
| `created_at` | `TIMESTAMP` | NOT NULL, Default: `NOW()` | Timestamp when created. |
| `updated_at` | `TIMESTAMP` | NOT NULL, Default: `NOW()` | Timestamp when updated. |

### Table: `categories`
Represents event categorizations managed by Administrators.

| Field Name | Data Type | Constraints | Description |
|---|---|---|---|
| `id` | `UUID` | Primary Key, Default: `gen_random_uuid()` | Unique category identifier. |
| `name` | `VARCHAR(100)` | NOT NULL, UNIQUE | Human-readable category name. |
| `description` | `TEXT` | NULL | Description of category content. |
| `created_at` | `TIMESTAMP` | NOT NULL, Default: `NOW()` | Timestamp when created. |

### Table: `events`
Represents an event published on the platform.

| Field Name | Data Type | Constraints | Description |
|---|---|---|---|
| `id` | `UUID` | Primary Key, Default: `gen_random_uuid()` | Unique event identifier. |
| `title` | `VARCHAR(200)` | NOT NULL | Event title. |
| `description` | `TEXT` | NOT NULL | Event markdown description. |
| `organizer_id` | `UUID` | Foreign Key REFERENCES `users(id)` ON DELETE CASCADE | Reference to the event creator. |
| `category_id` | `UUID` | Foreign Key REFERENCES `categories(id)` ON DELETE SET NULL | Reference to the event category. |
| `max_capacity` | `INTEGER` | NOT NULL, Check: `max_capacity > 0` | Maximum available spots. |
| `start_time` | `TIMESTAMP` | NOT NULL, Check: `start_time < end_time` | Event start time. |
| `end_time` | `TIMESTAMP` | NOT NULL | Event end time. |
| `status` | `VARCHAR(20)` | NOT NULL, Default: `Draft` | One of: `Draft`, `Published`, `Cancelled`. |
| `contact_email` | `VARCHAR(255)` | NOT NULL | Contact email for queries. |
| `location` | `VARCHAR(255)` | NOT NULL | Virtual meeting link or physical room number. |
| `created_at` | `TIMESTAMP` | NOT NULL, Default: `NOW()` | Timestamp when created. |
| `updated_at` | `TIMESTAMP` | NOT NULL, Default: `NOW()` | Timestamp when updated. |

### Table: `registrations`
Represents an Attendee's signup for an Event.

| Field Name | Data Type | Constraints | Description |
|---|---|---|---|
| `id` | `UUID` | Primary Key, Default: `gen_random_uuid()` | Unique registration identifier. |
| `event_id` | `UUID` | Foreign Key REFERENCES `events(id)` ON DELETE CASCADE | Target event. |
| `attendee_id` | `UUID` | Foreign Key REFERENCES `users(id)` ON DELETE CASCADE | Registering user. |
| `status` | `VARCHAR(20)` | NOT NULL, Default: `Registered` | One of: `Registered`, `Present`, `No-Show`, `Cancelled`. |
| `created_at` | `TIMESTAMP` | NOT NULL, Default: `NOW()` | Registration timestamp. |
| `updated_at` | `TIMESTAMP` | NOT NULL, Default: `NOW()` | Last status modification. |

---

## 3. Core Database Constraints & Rules
1.  **Unique Registrations**: There is a composite UNIQUE constraint on `registrations(event_id, attendee_id)` when status is not `Cancelled` to prevent attendees from registering multiple times for the same event.
2.  **Date Validations**: Events start date/time must be in the future during creation, and `start_time` must precede `end_time`.
3.  **No Organizer Publishing Before Approval**: A database trigger or business level constraint in Go will enforce that `events.status = 'Published'` is rejected if the associated `organizer_id` has `users.status != 'Approved'`.
