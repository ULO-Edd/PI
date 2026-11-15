# Quickstart Guide: Local Development Environment

This guide provides the necessary steps to get the University Admin Portal running on your local machine.

## 📋 Prerequisites

Ensure you have the following installed:
- **Docker & Docker Compose**
- **Node.js** (LTS version)
- **npm**

---

## 🚀 Getting Started

### 1. Start the Backend Infrastructure
The backend services and database are containerized. Use Docker Compose to bring them up in the correct sequence.

```bash
# Navigate to the local deployment directory
cd backend/deploy/local

# Start all services in detached mode and build images
docker compose up -d --build
```

**What this does**:
- Starts the **PostgreSQL** database.
- Launches the **Identity**, **Events**, and **Registrations** services.
- Configures an **Nginx API Gateway** on port `8080` to route your requests.

### 2. Start the Frontend Application
The frontend runs as a local Vite development server.

```bash
# Navigate to the frontend directory
cd frontend

# Install dependencies (first time only)
npm install

# Start the development server
npm run dev
```

The application will typically be available at `http://localhost:5173`.

---

## ✅ Verification

To verify that the system is healthy:

1. **Check Container Status**:
   ```bash
   docker compose ps
   ```
   All services should show a status of `(healthy)`.

2. **Test API Connectivity**:
   Open your browser or use `curl` to check the health of the Identity service via the gateway:
   ```bash
   curl http://localhost:8080/api/identity/health
   ```
   You should receive a `200 OK` response.

---

## 🛑 Stopping the Environment

To stop all backend services and remove the containers:

```bash
# From the backend/deploy/local directory
docker compose down
```

To completely wipe the database and start fresh:
```bash
docker compose down -v
```
