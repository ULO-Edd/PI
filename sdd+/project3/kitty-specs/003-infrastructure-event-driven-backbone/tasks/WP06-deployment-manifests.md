---
work_package_id: WP06
title: Deployment Manifests (k8s/docker-compose)
lane: "done"
dependencies: []
base_branch: main
base_commit: 9ff5b3f23042cbde8fe0ecd69b704270fb064ac6
created_at: '2026-06-02T16:14:18.114725+00:00'
subtasks: [T025, T026, T027, T028, T029]
requirement_refs:
- FR-008
- FR-009
- FR-010
shell_pid: "39196"
agent: "opencode"
reviewed_by: "Tiago Ramos"
review_status: "approved"
---

# WP06: Deployment Manifests (k8s/docker-compose)

## Objective
Provide consistent deployment configurations for both local development and production environments.

## Context
The platform consists of multiple services and a broker. Orchestration manifests ensure that networking, environment variables, and secrets are handled consistently across environments.

## Detailed Guidance

### Subtask T025: Create `backend/deploy/local/docker-compose.yml`
**Purpose**: Enable a "one-command" local startup for the entire backbone.
**Steps**:
1. Create a Docker Compose file including:
   - `rabbitmq`: Using image `rabbitmq:3-management`.
   - `analytics-worker`: Built from `backend/src/Services/AnalyticsWorker/Dockerfile`.
   - `notification-worker`: Built from `backend/src/Services/NotificationWorker/Dockerfile`.
   - `dead-letter-service`: Built from `backend/src/Services/DeadLetterService/Dockerfile`.
2. Configure environment variables for each service to point to the `rabbitmq` container.
3. Set up a network for inter-container communication.
**Validation**:
- [ ] `docker-compose up` boots all components and they can communicate.

### Subtask T026: Create `backend/deploy/k8s/rabbitmq.yaml`
**Purpose**: Define the production RabbitMQ cluster configuration.
**Steps**:
1. Create a Kubernetes `StatefulSet` for RabbitMQ to ensure persistent storage.
2. Define a `Service` (ClusterIP) for internal communication.
3. Configure resources (CPU/Memory limits) and health probes.
**Validation**:
- [ ] Manifest validates with `kubectl apply --dry-run=client`.

### Subtask T027: Create `backend/deploy/k8s/workers.yaml`
**Purpose**: Define the deployment of background worker services.
**Steps**:
1. Create `Deployment` manifests for `AnalyticsWorker` and `NotificationWorker`.
2. Configure `replica` counts for scaling.
3. Implement `livenessProbe` and `readinessProbe` checking the broker connection.
**Validation**:
- [ ] Manifests are valid and follow K8s best practices.

### Subtask T028: Create `backend/deploy/k8s/deadletter.yaml`
**Purpose**: Define the deployment of the Dead-Letter management service.
**Steps**:
1. Create a `Deployment` and `Service` for the `DeadLetterService`.
2. Ensure it has the necessary environment variables for its isolated relational database.
**Validation**:
- [ ] Manifest is valid.

### Subtask T029: Create `backend/deploy/k8s/configmaps-secrets.yaml`
**Purpose**: Securely manage environment configurations and credentials.
**Steps**:
1. Define `ConfigMap` for non-sensitive settings (e.g., RabbitMQ host, queue names).
2. Define `Secret` templates for sensitive data (e.g., DB passwords, API keys).
3. Map these as environment variables in the worker and service deployments.
**Validation**:
- [ ] Secrets are correctly mounted as environment variables.

## Definition of Done
- [ ] `docker-compose.yml` fully functional for local dev.
- [ ] All Kubernetes YAMLs (`rabbitmq`, `workers`, `deadletter`, `config/secrets`) are created and valid.
- [ ] Environment variables are mapped correctly across all manifests.

## Implementation Command
`spec-kitty implement WP06 --base WP04`

## Activity Log

- 2026-06-02T16:14:20Z – opencode – shell_pid=35872 – lane=doing – Assigned agent via workflow command
- 2026-06-02T16:16:09Z – opencode – shell_pid=35872 – lane=for_review – Ready for review: implemented local docker-compose and production Kubernetes manifests
- 2026-06-02T17:23:46Z – opencode – shell_pid=39196 – lane=doing – Started review via workflow command
- 2026-06-02T17:24:04Z – opencode – shell_pid=39196 – lane=done – Review passed: Implementation successfully covers all subtasks (T025-T029), providing functional local docker-compose setup and valid production Kubernetes manifests.
