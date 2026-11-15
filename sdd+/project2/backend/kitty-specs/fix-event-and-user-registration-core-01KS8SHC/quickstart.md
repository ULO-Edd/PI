# Quickstart Guide: Fix Event and User Registration Core

**Mission ID**: `01KS8SHC2ETY63EWAVDBNA8PP0`

Follow these steps to run, interact with, and verify the Event Management API.

## 1. Prerequisites
- Docker & Docker Compose installed.
- Go 1.25+ installed locally (for running tests without Docker).

---

## 2. Running the Application

To build and start the database and API services, run:

```bash
# Force a clean rebuild and clear any orphaned volumes
docker compose down -v
docker compose up --build
```

You should see:
- `db` container starting up and readiness logs.
- `api` container showing migration success messages: `Migrations ran successfully` and starting on port `:8080`.

---

## 3. Interactive Verification with Swagger

Once the containers are running:
1. Open your browser and navigate to: `http://localhost:8080/swagger/index.html`.
2. This displays the OpenAPI documentation generated using `swag`.
3. You can execute requests directly through the Swagger UI.

---

## 4. Running Tests

To run the unit and integration tests locally:

```bash
# Run all tests in the repository
go test ./... -v
```
