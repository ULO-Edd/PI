# Planning Research: Fix Event and User Registration Core

**Mission ID**: `01KS8SHC2ETY63EWAVDBNA8PP0`

This document details the research findings, technical investigations, and architectural decisions made to address database schema migrations and endpoints stubs for user and event registration.

## 1. Investigation of Migration Failure

### Findings
- The project implements schema migrations using the `github.com/golang-migrate/migrate/v4` package in Go.
- The migration file `000001_initial_schema.up.sql` contains BOTH the table creations AND the table destructions (`DROP TABLE`) at the end of the same file under `-- +migrate Down`.
- Since `golang-migrate` does not parse `-- +migrate Up/Down` annotations (which belong to `rubenv/sql-migrate`), it executes the entire script when running `Up`.
- Consequently, the tables are successfully created, and then **immediately dropped** in the same execution run, leaving the database completely empty while marking version `1` as successfully applied in `schema_migrations`.

### Resolution
- Split the migration into two separate, dedicated files:
  - `000001_initial_schema.up.sql`: Contains only `CREATE TABLE` statements.
  - `000001_initial_schema.down.sql`: Contains only `DROP TABLE` statements.
- Ensure the database volume can be reset during development using `docker compose down -v` to clear the cached, dirty schema version.

---

## 2. Domain-Level Validation Design

### Findings
- Currently, validations are mostly stubbed or absent.
- Based on alignment, all constraints (username uniqueness, password strength, role validation, registration limits, and event capacities) will be handled in Go service layers.

### Resolution
- Define custom domain errors:
  - `ErrUsernameTaken = errors.New("username is already taken")`
  - `ErrInvalidRole = errors.New("role must be attendee, organiser, or administrator")`
  - `ErrEventFull = errors.New("event has reached its maximum capacity")`
  - `ErrAlreadyRegistered = errors.New("user is already registered for this event")`
  - `ErrNotRegistered = errors.New("user is not registered for this event")`
  - `ErrUnauthorizedAction = errors.New("user is not authorized to perform this action")`
- Map these domain errors inside the API handlers to standard HTTP status codes (e.g. `409 Conflict`, `400 Bad Request`, `403 Forbidden`).

---

## 3. Event CRUD & Search/Filter Endpoints

### Findings
- Endpoints for creating, updating, and deleting events are defined but lack full implementation or verification of the actual organizer.
- Listing events requires pagination, location filtering, search terms, and status filtering (`draft`, `published`, `cancelled`).

### Resolution
- Add SQL queries or repository functions to support optional filters in `GET /events`:
  - `status` (e.g. `published`)
  - `location` (partial match)
  - `organiser_id` (UUID match)
- Restrict writing endpoints (`POST`, `PUT`, `DELETE` events) so they verify that the authenticated user's ID matches the `organiser_id` of the event (or that the user is an `admin`).
