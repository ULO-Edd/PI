---
work_package_id: WP02
title: Dead-Letter Service & DB Persistence
lane: "done"
dependencies: []
base_branch: main
base_commit: 4add4f960c1b94caabec281a2d7a4f0c213de870
created_at: '2026-06-02T15:49:16.842137+00:00'
subtasks: [T006, T007, T008, T009, T010, T011]
requirement_refs:
- FR-007
- FR-011
- FR-012
- FR-013
shell_pid: "10460"
agent: "opencode"
reviewed_by: "Tiago Ramos"
review_status: "approved"
---

# WP02: Dead-Letter Service & DB Persistence

## Objective
Implement the service that persists failed events in a relational database for administrative oversight.

## Context
Standard RabbitMQ DLQs provide basic storage, but an external database is required for rich querying, reporting, and manual administrative actions (replay/purge) as specified in FR-011 and FR-012.

## Detailed Guidance

### Subtask T006: Initialize EF Core SQLite DB context for `DeadLetterEntry`
**Purpose**: Create the persistent store for failed event messages.
**Steps**:
1. Create `backend/src/Services/DeadLetterService/Data/`:
   - Define `DeadLetterDbContext` inheriting from `DbContext`.
   - Create `DeadLetterEntry` entity with fields: `Id` (PK), `MessageId`, `CorrelationId`, `EventType`, `OriginalQueueName`, `TargetConsumer`, `FailureReason`, `ErrorStackTrace`, `Payload` (JSON), `FailureTimestamp`, `RetryHistory` (JSON), and `Status` (enum: Pending, Replayed, Purged).
2. Configure SQLite as the provider.
3. Run initial migration to create the database.
**Validation**:
- [ ] Database file is created.
- [ ] Schema matches the `data-model.md` definition.

### Subtask T007: Implement `DeadLetterService` API
**Purpose**: Create the business logic for interacting with the failure store.
**Steps**:
1. Create `backend/src/Services/DeadLetterService/Services/FailureManager.cs`:
   - `SaveFailure(entry)`: Persists a failed message.
   - `GetFailures(filter, paging)`: Retrieves lists of failures.
   - `MarkAsReplayed(id)`: Updates status to `Replayed`.
   - `MarkAsPurged(id)`: Updates status to `Purged`.
2. Implement the logic to re-publish a message to RabbitMQ using the `OriginalQueueName` while preserving the `CorrelationId`.
**Validation**:
- [ ] Logic correctly updates DB status and interacts with RabbitMQ.

### Subtask T008: Create API endpoint: GET `/api/admin/dlq`
**Purpose**: Allow Admins to list all failed events.
**Steps**:
1. Implement handler in `DeadLetterController`:
   - Support query parameters: `page`, `pageSize`, `originalQueueName`.
   - Return a paginated list of `DeadLetterEntry` summaries.
**Validation**:
- [ ] Request returns 200 OK with correct paginated data.

### Subtask T009: Create API endpoint: GET `/api/admin/dlq/{id}`
**Purpose**: Allow Admins to perform deep inspection of a specific failure.
**Steps**:
1. Implement handler in `DeadLetterController`:
   - Fetch a single entry by `Id`.
   - Include full `ErrorStackTrace` and `RetryHistory`.
**Validation**:
- [ ] Request returns 200 OK with full details.
- [ ] Non-existent ID returns 404.

### Subtask T010: Create API endpoint: POST `/api/admin/dlq/{id}/replay`
**Purpose**: Trigger a manual retry of a failed message.
**Steps**:
1. Implement handler in `DeadLetterController`:
   - Fetch the entry from the DB.
   - Re-publish the `Payload` to the `OriginalQueueName` using `RabbitMQ.Client`.
   - Set status to `Replayed` in the DB.
**Validation**:
- [ ] Message appears in the original queue.
- [ ] DB status updates to `Replayed`.

### Subtask T011: Create API endpoint: DELETE `/api/admin/dlq/{id}/purge`
**Purpose**: Permanently remove a failure record.
**Steps**:
1. Implement handler in `DeadLetterController`:
   - Mark the record as `Purged` (logical delete) or delete from DB (physical delete).
**Validation**:
- [ ] Entry no longer appears in `GET /api/admin/dlq`.

## Definition of Done
- [ ] Relational DB is operational with correct schema.
- [ ] All 4 admin endpoints are implemented and functional.
- [ ] Replay logic successfully re-queues messages with original correlation IDs.
- [ ] Logical or physical purging of records is working.

## Implementation Command
`spec-kitty implement WP02 --base WP01`

## Activity Log

- 2026-06-02T15:49:20Z – opencode – shell_pid=15780 – lane=doing – Assigned agent via workflow command
- 2026-06-02T15:51:12Z – opencode – shell_pid=15780 – lane=for_review – Ready for review: implemented relational DLQ storage and admin management API
- 2026-06-02T17:11:47Z – opencode – shell_pid=10460 – lane=doing – Started review via workflow command
- 2026-06-02T17:15:53Z – opencode – shell_pid=10460 – lane=done – Review passed: Implementation successfully covers all subtasks (T006-T011), with correct DB schema and functional API endpoints for dead-letter management.
