# Data Models: Event Management RESTful API

This document defines the database schema based on the entities identified in the feature specification.

## User

Represents a user of the application.

- **`id`**: Integer, Primary Key, Auto-increment
- **`username`**: String, Unique, Not Null
- **`hashed_password`**: String, Not Null
- **`role`**: String, Not Null (Enum: 'attendee', 'organiser', 'admin')
- **`events`**: Relationship to `Event` (one-to-many)
- **`registrations`**: Relationship to `Registration` (one-to-many)

## Event

Represents an event created by an organiser.

- **`id`**: Integer, Primary Key, Auto-increment
- **`title`**: String, Not Null
- **`description`**: String, Not Null
- **`location`**: String, Not Null
- **`start_datetime`**: DateTime, Not Null
- **`end_datetime`**: DateTime, Not Null
- **`capacity`**: Integer, Not Null
- **`status`**: String, Not Null (Enum: 'draft', 'published', 'cancelled')
- **`organiser_id`**: Integer, Foreign Key to `User.id`
- **`registrations`**: Relationship to `Registration` (one-to-many)

## Registration

Represents a user's registration for an event.

- **`id`**: Integer, Primary Key, Auto-increment
- **`user_id`**: Integer, Foreign Key to `User.id`
- **`event_id`**: Integer, Foreign Key to `Event.id`
- **`timestamp`**: DateTime, Not Null, Default: current time
