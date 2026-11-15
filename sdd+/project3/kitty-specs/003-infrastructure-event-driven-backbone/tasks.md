# Work Packages: Infrastructure Event-Driven Backbone

This document outlines the work packages for implementing the infrastructure event-driven backbone.

## Work Packages

- **WP01**: Foundation Setup & Message Broker (Dependencies: None)
- **WP02**: Dead-Letter Service & DB Persistence (Dependencies: WP01)
- **WP03**: Analytics Worker Service (Dependencies: WP01)
- **WP04**: Notification Worker Service (Dependencies: WP01)
- **WP05**: API Gateway & DLQ Admin Endpoints (Dependencies: WP02)
- **WP06**: Deployment Manifests (k8s/docker-compose) (Dependencies: WP01, WP02, WP03, WP04)

---

## WP01: Foundation Setup & Message Broker
**Summary**: Configure RabbitMQ and foundational infrastructure for the event-driven system.
**Priority**: High
**Independent Test**: Verify RabbitMQ startup and basic connectivity with test producer/consumer.
**Included Subtasks**:
- [x] T001: [P] Setup RabbitMQ Docker Compose definition.
- [x] T002: [P] Define `checkin-exchange` in RabbitMQ setup.
- [x] T003: [P] Configure RabbitMQ queues (`analytics-queue`, `notification-queue`, `deadletter-queue`).
- [x] T004: [P] Implement base `Shared` library for event models.
- [x] T005: [P] Implement `CorrelationId` middleware/helper.
**Estimated Prompt Size**: ~350 lines
**Requirement Refs**: FR-001, FR-002

---

## WP02: Dead-Letter Service & DB Persistence
**Summary**: Implement the service to persist and manage failed events.
**Priority**: High
**Independent Test**: Verify failed messages from worker retries are correctly logged in the DLQ relational database.
**Included Subtasks**:
- [x] T006: [P] Initialize EF Core SQLite DB context for `DeadLetterEntry`.
- [x] T007: [P] Implement `DeadLetterService` API.
- [x] T008: [P] Create API endpoint: GET `/api/admin/dlq`.
- [x] T009: [P] Create API endpoint: GET `/api/admin/dlq/{id}`.
- [x] T010: [P] Create API endpoint: POST `/api/admin/dlq/{id}/replay`.
- [x] T011: [P] Create API endpoint: DELETE `/api/admin/dlq/{id}/purge`.
**Estimated Prompt Size**: ~450 lines
**Requirement Refs**: FR-007, FR-011, FR-012, FR-013

---

## WP03: Analytics Worker Service
**Summary**: Build worker to process check-in events for analytics.
**Priority**: High
**Independent Test**: Check-in event consumption updates analytics storage.
**Included Subtasks**:
- [x] T012: [P] Initialize `AnalyticsWorker` service.
- [x] T013: [P] Configure RabbitMQ connection and `analytics-queue` consumption.
- [x] T014: [P] Implement `CheckInEventConsumer` logic.
- [x] T015: [P] Implement Polly retry policy (3 attempts, exponential backoff).
- [x] T016: [P] Implement DLX routing on final failure.
**Estimated Prompt Size**: ~400 lines
**Requirement Refs**: FR-003, FR-005, FR-006, FR-007

---

## WP04: Notification Worker Service
**Summary**: Build worker to process check-in events for notifications.
**Priority**: High
**Independent Test**: Check-in event consumption triggers notification call.
**Included Subtasks**:
- [x] T017: [P] Initialize `NotificationWorker` service.
- [x] T018: [P] Configure RabbitMQ connection and `notification-queue` consumption.
- [x] T019: [P] Implement `CheckInEventConsumer` logic.
- [x] T020: [P] Implement Polly retry policy (3 attempts, exponential backoff).
- [x] T021: [P] Implement DLX routing on final failure.
**Estimated Prompt Size**: ~400 lines
**Requirement Refs**: FR-004, FR-005, FR-006, FR-007

---

## WP05: API Gateway & DLQ Admin Endpoints
**Summary**: Secure and expose DLQ administration via the API Gateway.
**Priority**: Medium
**Independent Test**: Admin user can successfully call DLQ endpoints.
**Included Subtasks**:
- [x] T022: [P] Configure Gateway routing to `DeadLetterService`.
- [x] T023: [P] Implement JWT role check for DLQ endpoints (Admin access).
- [x] T024: [P] Integrate gateway with DLQ management endpoints.
**Estimated Prompt Size**: ~250 lines
**Requirement Refs**: FR-011, FR-012

---

## WP06: Deployment Manifests (k8s/docker-compose)
**Summary**: Define infrastructure deployment configuration.
**Priority**: Medium
**Independent Test**: Deploy all services to test cluster/local-docker.
**Included Subtasks**:
- [x] T025: [P] Create `backend/deploy/local/docker-compose.yml`.
- [x] T026: [P] Create `backend/deploy/k8s/rabbitmq.yaml`.
- [x] T027: [P] Create `backend/deploy/k8s/workers.yaml`.
- [x] T028: [P] Create `backend/deploy/k8s/deadletter.yaml`.
- [x] T029: [P] Create `backend/deploy/k8s/configmaps-secrets.yaml`.
**Estimated Prompt Size**: ~300 lines
**Requirement Refs**: FR-008, FR-009, FR-010
