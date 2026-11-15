# Quickstart: Infrastructure Event-Driven Backbone

This guide outlines how to bootstrap the event-driven infrastructure, run consumer services locally, and manage the Dead-Letter Queue (DLQ).

## Prerequisites

- Docker Desktop (or engine) running.
- .NET 10 SDK installed.
- Access to the platform's shared infrastructure configuration.

## 1. Local Infrastructure Setup

Use Docker Compose to spin up the message broker and local storage dependencies.

```bash
cd backend/deploy/local
docker-compose up -d rabbitmq failure-db
```

Verify RabbitMQ is healthy at `http://localhost:15672` (guest/guest).

## 2. Running Services

Ensure the backend services can connect to the broker.

```bash
# Run workers
cd backend/src/Services/AnalyticsWorker
dotnet run &

cd backend/src/Services/NotificationWorker
dotnet run &

# Run DeadLetter Service
cd backend/src/Services/DeadLetterService
dotnet run &
```

## 3. Verifying Event Flow

1. **Publish a test event** to RabbitMQ:
   - Use the RabbitMQ Management UI to publish a `check_in` event to the `checkin-exchange` with routing key `checkin.event`.

2. **Verify Consumers**:
   - Check logs for `AnalyticsWorker` and `NotificationWorker` to confirm they received the event.

3. **Verify DLQ Handling**:
   - Publish a message designed to fail (e.g., set payload to trigger validation error).
   - Observe 3 retries in worker logs.
   - Confirm the message appears in the `deadletter-db`.
   - Verify the `DeadLetterService` API reports the failure.
