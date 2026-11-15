# Quickstart Guide: Backend Microservices Synchronous Core

This guide explains how to restore, build, migrate databases, and spin up the synchronous backend microservices ecosystem locally.

## Prerequisites

- **.NET 10 SDK** or higher  
- **Docker** (optional, for local container runs)

## Setup and Restore

1. Navigate to the backend directory:
   ```bash
   cd backend
   ```

2. Restore all project dependencies and tooling:
   ```bash
   dotnet restore
   ```

3. Compile the complete project:
   ```bash
   dotnet build
   ```

## Database Migrations (EF Core 10)

Each microservice maintains an isolated relational database. Apply migrations to all three databases using the EF Core CLI:

```bash
# Apply Identity Database Migrations
dotnet ef database update --project src/Services/IdentityService --context IdentityDbContext

# Apply Events Database Migrations
dotnet ef database update --project src/Services/EventsService --context EventsDbContext

# Apply Registrations Database Migrations
dotnet ef database update --project src/Services/RegistrationsService --context RegistrationsDbContext
```

## Running the Ecosystem Locally

Spin up the API Gateway and the core services in separate terminal sessions or use a Docker Compose profile:

```bash
# Run the secure API Gateway
dotnet run --project src/ApiGateway

# Run individual microservices
dotnet run --project src/Services/IdentityService
dotnet run --project src/Services/EventsService
dotnet run --project src/Services/RegistrationsService
```

## Running Backend Verification Tests

Verify backend validations and inter-service HTTP communications:

```bash
dotnet test tests/UnitTests
dotnet test tests/IntegrationTests
```
