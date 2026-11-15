---
work_package_id: WP03
title: Analytics Worker Service
lane: "done"
dependencies: []
base_branch: main
base_commit: 0d10dd9df5e173fc79737f30ae6ad60c1f270df5
created_at: '2026-06-02T15:58:43.676229+00:00'
subtasks: [T012, T013, T014, T015, T016]
requirement_refs:
- FR-003
- FR-005
- FR-006
- FR-007
shell_pid: "31232"
agent: "opencode"
reviewed_by: "Tiago Ramos"
review_status: "approved"
---

# WP03: Analytics Worker Service

## Objective
Implement the background worker that processes check-in events to update system analytics.

## Context
This service must be fully asynchronous. It should not block the check-in flow and must handle transient failures gracefully using retries before escalating to the DLQ.

## Detailed Guidance

### Subtask T012: Initialize `AnalyticsWorker` service
**Purpose**: Scaffold the worker project.
**Steps**:
1. Create `backend/src/Services/AnalyticsWorker/`:
   - Implement as an ASP.NET Core `BackgroundService` (inheriting from `BackgroundService`).
   - Configure `appsettings.json` for RabbitMQ connection strings.
**Validation**:
- [ ] Service starts and logs "Analytics Worker started" on boot.

### Subtask T013: Configure RabbitMQ connection and `analytics-queue` consumption
**Purpose**: Establish the bridge between the broker and the worker.
**Steps**:
1. Implement a `RabbitMQConnection` helper using `RabbitMQ.Client`.
2. In `ExecuteAsync`, create a `Model` (Channel) and start a consumer on `analytics-queue`.
3. Ensure `BasicQos` (prefetch count) is set (e.g., to 10) to prevent the worker from being overwhelmed.
**Validation**:
- [ ] Service successfully connects to RabbitMQ and waits for messages.

### Subtask T014: Implement `CheckInEventConsumer` logic
**Purpose**: Process the actual business logic of a check-in event.
**Steps**:
1. Create `CheckInEventConsumer` class:
   - Deserialize the AMQP body into the `CheckInEvent` shared model.
   - Perform the analytics update (e.g., increment attendee count in the analytics store).
   - Log the processed `CorrelationId` and `MessageId`.
**Validation**:
- [ ] Valid messages trigger the analytics update logic.
- [ ] Corrupt messages (JSON parsing failure) are identified.

### Subtask T015: Implement Polly retry policy (3 attempts, exponential backoff)
**Purpose**: Handle transient errors (e.g., DB timeouts) without failing the message immediately.
**Steps**:
1. Integrate `Polly` into the consumer pipeline.
2. Define a `RetryPolicy`:
   - Max 3 retries.
   - Exponential backoff: 2s, 4s, 8s.
   - Only retry on `TransientException` types.
**Validation**:
- [ ] Simulated transient failures result in 3 retries before giving up.
- [ ] Successful recovery on 2nd attempt stops further retries.

### Subtask T016: Implement DLX routing on final failure
**Purpose**: Ensure "poison" messages do not block the queue and are routed for admin review.
**Steps**:
1. Configure the RabbitMQ channel to use `BasicNack(requeue: false)` when the Polly retry limit is reached.
2. Verify that the queue argument `x-dead-letter-exchange` (set in WP01) routes the message to the `deadletter-queue`.
**Validation**:
- [ ] A message that fails 3 times is automatically moved to the `deadletter-queue`.

## Definition of Done
- [ ] Analytics Worker consumes from `analytics-queue`.
- [ ] Events are processed asynchronously.
- [ ] Polly retry policy is active and correctly handles transient errors.
- [ ] Permanent failures are successfully routed to the DLQ.

## Implementation Command
`spec-kitty implement WP03 --base WP01`

## Activity Log

- 2026-06-02T15:58:47Z – opencode – shell_pid=46876 – lane=doing – Assigned agent via workflow command
- 2026-06-02T16:00:22Z – opencode – shell_pid=46876 – lane=for_review – Ready for review: implemented analytics worker with RabbitMQ consumption, Polly exponential backoff, and DLX escalation
- 2026-06-02T17:16:38Z – opencode – shell_pid=31232 – lane=doing – Started review via workflow command
- 2026-06-02T17:16:54Z – opencode – shell_pid=31232 – lane=done – Review passed: Implementation successfully covers all subtasks (T012-T016), with correct RabbitMQ integration, Polly retry policy, and DLX routing.
