# API Gateway Routing Contract

## Base URL: `http://localhost:8080`

## Endpoints

### 1. Identity Service
- **Path**: `/api/identity/*` $\rightarrow$ Proxied to `identity-service:80`
- **Purpose**: Authentication and user management.

### 2. Events Service
- **Path**: `/api/events/*` $\rightarrow$ Proxied to `events-service:80`
- **Purpose**: Event catalogue and management.

### 3. Registrations Service
- **Path**: `/api/registrations/*` $\rightarrow$ Proxied to `registrations-service:80`
- **Purpose**: Event registration and ticket handling.

### 4. System Health
- **Path**: `/health` $\rightarrow$ Proxied to `gateway-service:80/health`
- **Response**: `200 OK` if the gateway is up.
