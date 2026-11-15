# Research Findings: Infrastructure Event-Driven Backbone

This document compiles the research, design decisions, and evaluation of alternatives conducted during Phase 0 of the Infrastructure Event-Driven Backbone feature planning.

---

## 1. Message Broker Technology Selection

- **Decision**: RabbitMQ (using the standard AMQP protocol)
- **Rationale**: RabbitMQ provides native, robust support for exchange-queue binding topologies. This fits the required consumer fan-out architecture perfectly: a single `check-in` exchange can route messages to both the `analytics-worker` queue and `notification-worker` queue via specific routing keys. Furthermore, RabbitMQ has built-in support for Dead-Letter Exchanges (DLX) via queue arguments (`x-dead-letter-exchange` and `x-dead-letter-routing-key`), allowing messages to be automatically forwarded on failure.
- **Alternatives Considered**:
  - *Apache Kafka*: Evaluated but rejected. Kafka is designed for high-throughput replayable stream logs. It introduces significant operational complexity (Zookeeper/KRaft configuration, partition management) which is disproportionate for this platform's simple queuing requirements.
  - *Redis Pub/Sub*: Evaluated but rejected. Redis Pub/Sub lacks durable, persistent queuing models and does not offer native Dead-Letter Queue (DLQ) support, which violates resilience and operational safety constraints.

---

## 2. Background Worker Framework and Runtime

- **Decision**: ASP.NET Core (.NET 10) Background Workers (`IHostedService` / `BackgroundService`)
- **Rationale**: Writing the workers as C# background services ensures perfect technological alignment with the synchronous microservices core. This allows code reuse (such as shared integration event contracts and database entity definitions under `backend/src/Shared/`), unified logging, and standard Microsoft Dependency Injection pipelines.
- **Alternatives Considered**:
  - *Node.js (TypeScript) Workers*: Evaluated but rejected. Although lightweight, introducing Node.js for backend background workers would split the backend tech stack, requiring duplicate schema definitions, separate CI/CD pipelines, and different linting/testing patterns.

---

## 3. Dead-Letter Queue (DLQ) and Failure Storage Architecture

- **Decision**: Native RabbitMQ Dead-Letter Exchange coupled with a dedicated background consumer storing failed messages in a relational database.
- **Rationale**: To meet FR-011 and FR-012, administrators must be able to inspect, query, and perform actions (replay, purge) on DLQ entries. While standard AMQP DLQs store failed messages, querying them directly via the broker is slow and limits analytical capabilities. By routing failed messages to a `dead-letter-queue` consumed by a `DeadLetterService`, we can persist the failures in an isolated database (SQLite/PostgreSQL) with rich fields (error stacks, original queues, retry count, correlation IDs). The ApiGateway can then expose standard, fast REST endpoints querying this database. When a "replay" is requested, the service fetches the message from the DB, republishes it to the target queue preserving its correlation ID, and removes/marks the DB entry.
- **Alternatives Considered**:
  - *Broker API Poll (Direct AMQP/HTTP Management)*: Querying RabbitMQ's HTTP Management API directly. Rejected because it exposes rabbitmq administrative credentials to the Gateway, lacks custom query/search capabilities, and cannot maintain a permanent historical log once messages are purged or replayed.

---

## 4. Exponential Backoff and Retry Policies

- **Decision**: Polly (the .NET resilience library)
- **Rationale**: Polly is the industry-standard policy engine for .NET. It allows defining elegant Retry policies with exponential backoff (e.g., 2s, 4s, 8s delays) for exactly 3 attempts. It integrates natively with logging to ensure transient dependency exceptions (such as network hiccups or third-party service rate-limits) are clearly recorded before final DLQ escalation.
- **Alternatives Considered**:
  - *Custom Thread.Sleep Retry Loops*: Rejected because sleeping threads block asynchronous worker execution, are highly prone to edge-case bugs, and lack structured logging hooks.
