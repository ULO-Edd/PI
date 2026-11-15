# Quickstart Runbook: Event Management Web Platform
*Path: [kitty-specs/event-management-web-platform-01KSV4CE/quickstart.md](quickstart.md)*

This guide details how to build, run, and verify the multi-component web platform locally.

---

## 1. Prerequisites
Ensure you have the following installed on your machine:
*   [Docker Desktop](https://www.docker.com/products/docker-desktop/) with Docker Compose.
*   [Go 1.25+](https://go.dev/dl/) (optional, for local backend compilation).
*   [Node.js v18+](https://nodejs.org/) (optional, for local frontend testing).

---

## 2. Fast-Bootstrap & Launch

To spin up the entire application stack including the PostgreSQL database, Go REST API, Vue.js Frontend, and a developer mail catcher (e.g. Mailpit), run:

```bash
# Build and launch all services in the background
docker compose up --build -d
```

---

## 3. Local Port Reference

Once running, the following services are available:

*   **Frontend Vue Web App:** [http://localhost:5173](http://localhost:5173) (or dynamic dev port)
*   **Backend Go REST API:** [http://localhost:8080](http://localhost:8080)
*   **Interactive API Documentation (Swagger):** [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
*   **Developer Mail Catcher Dashboard:** [http://localhost:8025](http://localhost:8025) (previews sent HTML notifications)

---

## 4. Verification & Testing

### Running Backend Tests
To execute Go backend unit tests:
```bash
docker compose exec backend go test ./...
```

### Running Frontend Tests
To execute Vue component frontend unit tests:
```bash
docker compose exec frontend npm run test:unit
```

### Checking Container Health and Logs
To check logs for the services:
```bash
docker compose logs -f backend
docker compose logs -f frontend
```
To stop the workspace:
```bash
docker compose down
```
