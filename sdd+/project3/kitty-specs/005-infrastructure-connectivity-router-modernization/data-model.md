# Data Model: Infrastructure Connectivity

## 1. Infrastructure Entities

### Service
- **Description**: A containerized backend component.
- **Attributes**:
  - `name`: Unique identifier (e.g., "identity-service").
  - `internal_port`: Port the service listens on inside the container (e.g., 80).
  - `health_endpoint`: Path used for readiness checks (e.g., `/health`).
  - `dependencies`: List of other services that must be healthy before this service starts.

### Network Bridge
- **Description**: The isolated virtual network for container communication.
- **Attributes**:
  - `name`: "admin-portal-network".
  - `driver`: "bridge".

### Port Mapping
- **Description**: The bridge between the host machine and the containerized services.
- **Attributes**:
  - `host_port`: The single entry point (e.g., 8080).
  - `target_service`: The gateway service that routes traffic.

### Health State
- **Description**: The current operational status of a service.
- **Values**: `STARTING`, `HEALTHY`, `UNHEALTHY`.
