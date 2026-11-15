# Data Model

This document defines the database schema for the core entities of the Event Management API.

### `users` table

| Column | Type | Constraints | Description |
| :--- | :--- | :--- | :--- |
| `id` | `UUID` | Primary Key, Not Null | Unique identifier for the user. |
| `username` | `VARCHAR(255)` | Not Null, Unique | User's chosen username. |
| `password_hash` | `VARCHAR(255)` | Not Null | Hashed password. |
| `role` | `VARCHAR(50)` | Not Null | User role (`attendee`, `organiser`, `administrator`). |
| `created_at` | `TIMESTAMPZ` | Not Null, Default `NOW()` | Timestamp of user creation. |

### `events` table

| Column | Type | Constraints | Description |
| :--- | :--- | :--- | :--- |
| `id` | `UUID` | Primary Key, Not Null | Unique identifier for the event. |
| `title` | `VARCHAR(255)` | Not Null | Title of the event. |
| `description` | `TEXT` | | Detailed description of the event. |
| `location` | `VARCHAR(255)` | | Physical or virtual location of the event. |
| `start_time` | `TIMESTAMPZ` | Not Null | Start date and time of the event. |
| `end_time` | `TIMESTAMPZ` | Not Null | End date and time of the event. |
| `organiser_id`| `UUID` | Foreign Key (`users.id`) | The user who organised the event. |
| `capacity` | `INTEGER` | Not Null, Default 0 | Maximum number of attendees. |
| `status` | `VARCHAR(50)` | Not Null, Default `draft` | Status of the event (`draft`, `published`, `cancelled`). |
| `created_at` | `TIMESTAMPZ` | Not Null, Default `NOW()` | Timestamp of event creation. |

### `registrations` table

| Column | Type | Constraints | Description |
| :--- | :--- | :--- | :--- |
| `id` | `UUID` | Primary Key, Not Null | Unique identifier for the registration. |
| `event_id` | `UUID` | Foreign Key (`events.id`), Not Null | The event being registered for. |
| `attendee_id` | `UUID` | Foreign Key (`users.id`), Not Null | The user registering for the event. |
| `created_at` | `TIMESTAMPZ` | Not Null, Default `NOW()` | Timestamp of registration. |
