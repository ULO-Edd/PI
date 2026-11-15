# Data Model: Infrastructure Event-Driven Backbone

This document specifies the key entities, database schemas, and state transitions for the event-driven backbone's persistent and in-flight structures.

---

## 1. Event Message (In-Flight AMQP Payload)

Represented as a JSON-serialized payload published to RabbitMQ.

### Fields

| Field Name | Type | Description |
|---|---|---|
| `messageId` | `string (UUID)` | Unique identifier for the message instance. |
| `correlationId` | `string (UUID)` | Preserved across retries and manual replays to ensure end-to-end event lineage and tracing. |
| `eventType` | `string` | Type of event (e.g., `check_in`, `notification_request`). |
| `timestamp` | `string (ISO 8601)`| UTC timestamp when the event was originally created. |
| `payload` | `string (JSON)` | Event-specific data context. |
| `retryAttemptCount`| `int` | Progress of retries (initialized to 0, incremented during retries up to 3). |

---

## 2. Dead-Letter Entry (Persistent Failure Record)

Stored in the isolated SQL database of the `DeadLetterService`. These records persist failed messages that surpassed retry thresholds or were rejected as toxic, allowing administrative action.

### Database Schema (EF Core Representation)

| Column Name | Database Type | Nullable | Key | Description |
|---|---|---|---|---|
| `Id` | `INTEGER` | No | PK | Auto-incrementing internal primary key. |
| `MessageId` | `TEXT (UUID)` | No | - | Original message unique identifier. |
| `CorrelationId` | `TEXT (UUID)` | No | Index | Preserved correlation ID for lineage. |
| `EventType` | `TEXT` | No | - | Event type of the failed message. |
| `OriginalQueueName`| `TEXT` | No | - | Name of the queue where the message failed (e.g., `analytics-checkin-queue`). |
| `TargetConsumer` | `TEXT` | No | - | Name of the consumer class/service that failed. |
| `FailureReason` | `TEXT` | No | - | Summary error message of the final failure. |
| `ErrorStackTrace` | `TEXT` | Yes | - | Full call stack trace of the failure for analysis. |
| `Payload` | `TEXT (JSON)` | No | - | Entire serialized payload of the original event message. |
| `FailureTimestamp` | `TEXT` | No | - | UTC timestamp when the record was written to DLQ. |
| `RetryHistory` | `TEXT (JSON)` | No | - | Array of timestamps/errors from the 3 failed backoff attempts. |
| `Status` | `TEXT` | No | - | Lifecycle state: `PendingAction`, `Replayed`, `Purged`. |

### State Transitions

Failed event message moves through the following states in the DeadLetter database:

```
        +---------------+
        | PendingAction |
        +---------------+
          /           \
         /             \ (Admin manual 'purge')
        /               \
       v                 v
  +----------+     +----------+
  | Replayed |     |  Purged  |
  +----------+     +----------+
```

1. **PendingAction**: The initial state when a failed message is stored.
2. **Replayed**: Set when an authorized Admin triggers a manual 'replay' endpoint. The event is re-published to its `OriginalQueueName` with its original `CorrelationId`, and the DB status updates to record the successful action.
3. **Purged**: Set when an Admin deletes the record. The record can be physically or logically deleted from the system.

---

## 3. Deployment Manifest Specification

This specifies the structures of orchestration configurations for services.

### Properties

- **Target Service Name**: Maps to container identifiers (e.g., `analytics-worker`, `notification-worker`, `dead-letter-service`).
- **Environment Variable Mappings**: Loaded dynamically during boot (such as `ConnectionStrings:DbConnection`, `MessageBroker:Host`, `MessageBroker:Username`).
- **Secret Mounts**: Encrypted credentials (e.g., certificates or access tokens) mounted as volumes or loaded via container env secrets.
- **Health Checks**: Standard TCP or HTTP readiness and liveness probes (e.g., `/health` or `amqp://broker` check).
