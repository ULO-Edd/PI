---
work_package_id: WP01
title: Foundation Setup & Message Broker
lane: "done"
dependencies: []
base_branch: main
base_commit: 99e7813d5cad7ac9c1d500782632c6c452340b51
created_at: '2026-06-02T15:40:05.033569+00:00'
subtasks: [T001, T002, T003, T004, T005]
requirement_refs:
- FR-001
- FR-002
shell_pid: "36268"
agent: "opencode"
reviewed_by: "Tiago Ramos"
review_status: "approved"
---

# WP01: Foundation Setup & Message Broker

## Objective
Establish the RabbitMQ messaging backbone and shared contracts required for all event-driven services.

## Context
This is the critical path for the entire feature. All other services (Workers, DeadLetterService) depend on the exchange and queue topology defined here.

## Detailed Guidance

### Subtask T001: Setup RabbitMQ Docker Compose definition
**Purpose**: Provide a reproducible local development environment for the message broker.
**Steps**:
1. Create `backend/deploy/local/docker-compose.yml`:
   - Use image `rabbitmq:3-management` to enable the Management UI.
   - Map ports `5672` (AMQP) and `15672` (UI).
   - Set environment variables for default user/pass (guest/guest) and hostname.
   - Ensure a persistent volume is mapped for `/var/lib/rabbitmq` to prevent data loss on restart.
**Validation**:
- [ ] `docker-compose up -d` starts the container without errors.
- [ ] RabbitMQ Management UI is accessible at `http://localhost:15672`.

### Subtask T002: Define `checkin-exchange` in RabbitMQ setup
**Purpose**: Establish the central entry point for all check-in events.
**Steps**:
1. Implement an initialization script or use a RabbitMQ definition JSON file:
   - Create an exchange named `checkin-exchange`.
   - Set the type to `topic` to allow flexible routing via routing keys.
   - Ensure the exchange is `durable: true` to survive broker restarts.
**Validation**:
- [ ] Exchange `checkin-exchange` is visible in the Management UI under the "Exchanges" tab.

### Subtask T003: Configure RabbitMQ queues
**Purpose**: Create dedicated queues for each consumer to ensure decoupled processing.
**Steps**:
1. Configure the following queues as `durable: true`:
   - `analytics-queue`: Bound to `checkin-exchange` with routing key `checkin.event`.
   - `notification-queue`: Bound to `checkin-exchange` with routing key `checkin.event`.
   - `deadletter-queue`: Bound to a dedicated `dead-letter-exchange` (DLX).
2. Set `x-dead-letter-exchange` argument on `analytics-queue` and `notification-queue` to route failed messages automatically to the DLX.
**Validation**:
- [ ] All three queues are created and correctly bound in the Management UI.
- [ ] Messages published to `checkin-exchange` with key `checkin.event` appear in both worker queues.

### Subtask T004: Implement base `Shared` library for event models
**Purpose**: Ensure type safety and consistency for event payloads across all services.
**Steps**:
1. Create `backend/src/Shared/Events/`:
   - Define `BaseEvent` abstract class with common properties: `MessageId` (UUID), `CorrelationId` (UUID), `Timestamp` (DateTime), and `EventType` (string).
   - Create `CheckInEvent` inheriting from `BaseEvent` with specific payload fields (e.g., `AttendeeId`, `EventId`, `CheckInTime`).
2. Use `System.Text.Json` for serialization.
**Validation**:
- [ ] Shared library compiles.
- [ ] Event models can be serialized/deserialized without data loss.

### Subtask T005: Implement `CorrelationId` middleware/helper
**Purpose**: Guarantee end-to-end traceability for every event through the system.
**Steps**:
1. Implement a `CorrelationContext` helper in `backend/src/Shared/`:
   - Method to extract `CorrelationId` from incoming AMQP message headers.
   - Method to inject `CorrelationId` into outgoing messages.
   - Log enrichment: Ensure all logs include the current `CorrelationId`.
**Validation**:
- [ ] A published event's `CorrelationId` is correctly logged by the consumer.
- [ ] Replayed messages retain the original `CorrelationId`.

## Definition of Done
- [ ] RabbitMQ is running locally via Docker Compose.
- [ ] `checkin-exchange` and all three queues are configured and bound.
- [ ] Shared event library is implemented and usable.
- [ ] Correlation ID helper is functional and integrated into logging.

## Implementation Command
`spec-kitty implement WP01`

## Activity Log

- 2026-06-02T15:40:08Z – opencode – shell_pid=25332 – lane=doing – Assigned agent via workflow command
- 2026-06-02T15:43:37Z – opencode – shell_pid=25332 – lane=for_review – Ready for review: implemented RabbitMQ infrastructure, event models, and correlation helper
- 2026-06-02T16:31:17Z – opencode – shell_pid=36268 – lane=doing – Started review via workflow command
- 2026-06-02T16:41:15Z – opencode – shell_pid=36268 – lane=done – Review passed: RabbitMQ topology correctly configured with DLX, shared event models implemented, and correlation ID helper functional.
