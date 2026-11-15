# Data Model: Frontend Application Client

This document defines the key frontend data models and state shapes utilized in Pinia stores and UI views.

## 1. Event Model

Represents event details fetched from the Catalogue service and displayed to the user.

### Fields

- `id` (string, UUID): Unique identifier for the event.
- `title` (string): Title of the event (required, max 100 characters).
- `description` (string): Detailed event description.
- `dateTime` (string, ISO-8601): Event start timestamp.
- `location` (string): Venue name or online access link.
- `capacity` (number): Maximum number of registrations allowed.
- `availableTickets` (number): Remaining seats.
- `organiserId` (string, UUID): Reference to the creator organiser.

---

## 2. Attendee Model

Represents a registered attendee displayed in the dashboard and live roster lists.

### Fields

- `id` (string, UUID): Unique registration ID.
- `eventId` (string, UUID): Reference to the registered event.
- `fullName` (string): Attendee's full name.
- `email` (string): Attendee's email address (must match email format).
- `registrationTimestamp` (string, ISO-8601): When they registered.
- `checkInStatus` (string, enum): Active check-in state.
  - Transitions: `Pending` -> `CheckedIn`
- `checkInTimestamp` (string, ISO-8601, nullable): When they checked in.

---

## 3. CheckInScan Model (Offline Queue)

State shape of individual scanned entries queued locally in the Pinia check-in store.

### Fields

- `scanId` (string, UUID): Local tracking identifier.
- `eventId` (string, UUID): Target event identifier.
- `attendeeId` (string, UUID): Target attendee identifier.
- `scannedAt` (string, ISO-8601): Scan timestamp (client clock).
- `syncStatus` (string, enum): Queue sync state.
  - State transitions: `Queued` -> `Syncing` -> `Synced`
- `correlationId` (string): Correlation ID for tracing throughout the transaction.

---

## 4. AuthState Store Shape

State structure of the security and authentication store.

### Fields

- `isAuthenticated` (boolean): Current login state.
- `user` (object, nullable): Details of the authenticated user.
  - `id` (string, UUID): User identifier.
  - `email` (string): User email.
  - `role` (string, enum): Access level (`Admin`, `Organiser`, `Attendee`).
- `sessionExpiresAt` (string, ISO-8601): Expiry indicator of the current session.
