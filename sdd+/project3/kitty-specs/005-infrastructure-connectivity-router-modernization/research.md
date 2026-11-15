# Research: Infrastructure Connectivity & Router Modernization

## 1. Docker Compose Health Checks & Sequencing
- **Decision**: Use `pg_isready` for PostgreSQL and HTTP-based health checks for .NET services.
- **Rationale**: `pg_isready` is the standard tool for verifying Postgres availability without requiring a full connection. .NET services will implement a `/health` endpoint using the built-in `Microsoft.Extensions.Diagnostics.HealthChecks` library, which is the industry standard for ASP.NET Core.
- **Alternatives**: Simple TCP socket checks (too coarse, doesn't verify app readiness) or shell scripts (harder to maintain).

## 2. API Gateway Implementation
- **Decision**: Use a lightweight Nginx container as the API Gateway.
- **Rationale**: Nginx is highly performant, easy to configure via a single `nginx.conf`, and provides a clean separation between the host network and the internal Docker network. It allows us to map one host port (e.g., 8080) to multiple internal services based on path prefixes (e.g., `/api/identity` $\rightarrow$ `identity-service:80`).
- **Alternatives**: YARP (excellent but requires maintaining another .NET service), Ocelot (similar to YARP).

## 3. Vue Router 4 Navigation Guards
- **Decision**: Refactor guards to return `true`, `false`, or a route location object instead of relying solely on the `next()` callback.
- **Rationale**: Vue Router 4 shifted towards a more functional approach for guards. Removing the `next` callback in favor of return values eliminates the deprecation warnings and aligns with the current framework standards.
- **Alternatives**: Staying on the old pattern (leads to continued warnings and future breakage).

## 4. Resilient Axios Error Handling
- **Decision**: Implement a global response interceptor that checks for `ERR_NETWORK` or "Connection Refused" messages.
- **Rationale**: Connectivity errors in the browser often manifest as generic network errors. By intercepting these at the Axios level, we can trigger a standardized Toast Notification that informs the user the backend is unavailable, rather than letting the app fail silently or show a generic 500 error.
- **Alternatives**: Per-request try-catch blocks (repetitive and inconsistent).
