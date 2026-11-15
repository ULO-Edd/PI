# Quickstart: Local Development Environment

## Prerequisites
- Docker and Docker Compose installed.
- Node.js and npm installed (for frontend).

## Startup Sequence

### 1. Backend Infrastructure
From the project root, run:
```bash
docker-compose up -d
```
This command will:
1. Create the `admin-portal-network` bridge.
2. Start PostgreSQL and wait for `pg_isready`.
3. Start Identity, Events, and Registrations services.
4. Start the Nginx API Gateway on port 8080.

### 2. Verify Health
Check the status of all services:
```bash
docker-compose ps
```
All services should show `(healthy)`.

### 3. Frontend Startup
In the `frontend` directory, run:
```bash
npm run dev
```
The frontend is now configured to communicate with the backend via `http://localhost:8080`.

## Troubleshooting
- **Connection Refused**: Ensure Docker Compose is running and the gateway container is healthy.
- **Database Errors**: Check the logs of the postgres container: `docker-compose logs postgres`.
- **Clean Restart**: `docker-compose down -v` followed by `docker-compose up -d` to reset volumes.
