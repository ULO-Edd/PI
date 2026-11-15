---
work_package_id: WP04
title: Notification Worker Service
lane: "done"
dependencies: []
base_branch: main
base_commit: 829c60d2e2b0ff506f0ad0e9ff64796e8b76f9bb
created_at: '2026-06-02T16:09:21.887745+00:00'
subtasks: [T017, T018, T019, T020, T021]
requirement_refs:
- FR-004
- FR-005
- FR-006
- FR-007
shell_pid: "26012"
agent: "opencode"
reviewed_by: "Tiago Ramos"
review_status: "approved"
---

# WP04: Notification Worker Service

## Objective
Implement the background worker that processes check-in events to trigger attendee notifications.

## Context
Notifications are high-latency operations (external API calls). This worker must run independently of analytics and handle failures without blocking other notifications.

## Detailed Guidance

### Subtask T017: Initialize `NotificationWorker` service
**Purpose**: Scaffold the notification worker project.
**Steps**:
1. Create `backend/src/Services/NotificationWorker/`:
   - Implement as an ASP.NET Core `BackgroundService`.
   - Configure `appsettings.json` for RabbitMQ and Notification API keys.
**Validation**:
- [ ] Service starts and logs "Notification Worker started" on boot.

### Subtask T018: Configure RabbitMQ connection and `notification-queue` consumption
**Purpose**: Establish the bridge between the broker and the notification worker.
**Steps**:
1. Use the shared `RabbitMQConnection` helper.
2. Start a consumer on `notification-queue`.
3. Set `BasicQos` prefetch count to a low value (e.g., 1) if the downstream notification API is strictly rate-limited.
**Validation**:
- [ ] Service successfully connects to RabbitMQ and waits for messages.

### Subtask T019: Implement `CheckInEventConsumer` logic
**Purpose**: Trigger the actual notification process.
**Steps**:
1. Create `NotificationConsumer` class:
   - Deserialize `CheckInEvent`.
   - Call the notification service (e.g., Email/Push API) using the `AttendeeId` from the event.
   - Log the processing status and `CorrelationId`.
**Validation**:
- [ ] Valid events trigger a notification API call.

### Subtask T020: Implement Polly retry policy (3 attempts, exponential backoff)
**Purpose**: Handle intermittent notification API timeouts.
**Steps**:
1. Implement a Polly `RetryPolicy` wrapping the API call:
   - Max 3 attempts.
   - Exponential backoff (2s, 4s, 8s).
**Validation**:
- [ ] Network timeouts trigger retries.
- [ ] Success on retry stops further attempts.

### Subtask T021: Implement DLX routing on final failure
**Purpose**: Ensure failed notifications are preserved for admin manual replay.
**Steps**:
1. Use `BasicNack(requeue: false)` upon the 3rd failed attempt.
2. Verify routing to the `deadletter-queue` via the broker's DLX configuration.
**Validation**:
- [ ] A permanently failing notification is moved to the DLQ.

## Definition of Done
- [ ] Notification Worker consumes from `notification-queue`.
- [ ] Notifications are triggered asynchronously.
- [ ] Polly retry policy handles transient API errors.
- [ ] Permanent failures are routed to the DLQ.

## Implementation Command
`spec-kitty implement WP04 --base WP01`

## Activity Log

- 2026-06-02T16:09:26Z – opencode – shell_pid=45132 – lane=doing – Assigned agent via workflow command
- 2026-06-02T16:11:11Z – opencode – shell_pid=45132 – lane=for_review – Ready for review: implemented notification worker with RabbitMQ consumption, Polly exponential backoff, and DLX escalation
- 2026-06-02T17:17:23Z – opencode – shell_pid=26012 – lane=doing – Started review via workflow command
- 2026-06-02T17:17:47Z – opencode – shell_pid=26012 – lane=done – Review passed: Implementation successfully covers all subtasks (T017-T021), with correct RabbitMQ integration, Polly retry policy, and DLX routing for the notification service.
