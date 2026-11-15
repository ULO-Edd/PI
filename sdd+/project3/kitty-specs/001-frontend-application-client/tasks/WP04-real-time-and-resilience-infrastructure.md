---
work_package_id: WP04
title: Real-Time & Resilience Infrastructure
lane: "done"
dependencies: [WP01]
base_branch: 001-frontend-application-client-WP01
base_commit: 2b072ef1391046fe7e1dadb57a2004e43ed8848e
created_at: '2026-06-01T14:26:01.984831+00:00'
subtasks:
- T011
- T012
- T013
phase: Phase 4 - Messaging & Resilience Setup
assignee: ''
agent: "opencode"
shell_pid: "42048"
review_status: "approved"
reviewed_by: "Tiago Ramos"
history:
- timestamp: '2026-06-01T13:35:24Z'
  lane: planned
  agent: system
  shell_pid: ''
  action: Prompt generated via /spec-kitty.tasks
requirement_refs:
- FR-006
- FR-008
- FR-009
---

# Work Package Prompt: WP04 – Real-Time & Resilience Infrastructure

## ⚠️ IMPORTANT: Review Feedback Status

**Read this first if you are implementing this task!**

- **Has review feedback?**: Check the `review_status` field above. If it says `has_feedback`, scroll to the **Review Feedback** section immediately (right below this notice).
- **You must address all feedback** before your work is complete. Feedback items are your implementation TODO list.
- **Mark as acknowledged**: When you understand the feedback and begin addressing it, update `review_status: acknowledged` in the frontmatter.
- **Report progress**: As you address each feedback item, update the Activity Log explaining what you changed.

---

## Review Feedback

*[This section is empty initially. Reviewers will populate it if the work is returned from review. If you see feedback here, treat each item as a must-do before completion.]*

---

## Markdown Formatting
Wrap HTML/XML tags in backticks: `` `<div>` ``, `` `<script>` ``
Use language identifiers in code blocks: ````javascript`, ````bash`

---

## Objectives & Success Criteria

- Build a resilient client WebSocket service capable of maintaining real-time attendee check-in event subscriptions.
- Integrate an automatic exponential backoff reconnection strategy to recover from network disconnections.
- Build a check-in Pinia store with automatic `LocalStorage` synchronization to queue scans during disconnections.

## Context & Constraints

- Refers to `kitty-specs/001-frontend-application-client/spec.md` (defining Pinia store, exponential backoff, WebSocket client updates).
- Requires a base branch with WP01:
  ```bash
  spec-kitty implement WP04 --base WP01
  ```

## Subtasks & Detailed Guidance

### Subtask T011 - Create resilient WebSocket connection service
- **Purpose**: Maintain connection to the backend live check-in WebSocket stream.
- **Steps**:
  1. Build `/frontend/src/services/websocket.ts` exporting a unified connection wrapper.
  2. Implement `connect()`, `subscribe(eventId)`, and `disconnect()` actions.
  3. Map incoming websocket events to register attendee check-ins dynamically.
  4. Dispatch custom browser events or Pinia actions when connection drops or is established.
- **Files**:
  - `/frontend/src/services/websocket.ts`
- **Parallel?**: Yes. Can proceed once scaffolding (WP01) is complete.

### Subtask T012 - Implement exponential backoff reconnection loops
- **Purpose**: Re-establish connection safely during server or network drops without overloading servers.
- **Steps**:
  1. Build a robust retry logic within the WebSocket service:
     - Set initial delay to 1,000ms.
     - Exponentially scale retry delays: `delay = delay * 2`.
     - Cap maximum retry delay at 30,000ms.
     - Reset delay to initial state upon successful connection.
- **Files**:
  - `/frontend/src/services/websocket.ts`
- **Parallel?**: No. Must be integrated directly into WebSocket connection handler.

### Subtask T013 - Build persistent Check-In Pinia store
- **Purpose**: Manage check-in queues and automatically persist them offline to client LocalStorage.
- **Steps**:
  1. Create `/frontend/src/stores/checkIn.ts` Pinia store containing:
     - `pendingQueue: CheckInScan[]` (queued local scan items)
     - `isConnected: boolean` (indicates live WebSocket status)
  2. Configure automatic LocalStorage sync on mutation of `pendingQueue`.
  3. Ensure scan items include structured correlation IDs and client-side timestamps.
- **Files**:
  - `/frontend/src/stores/checkIn.ts`
- **Parallel?**: Yes. Can run concurrently with WebSocket wrapper (T011).

## Risks & Mitigations

- **Infinite retries server impact**: Implement max delay (30s) and random jitter (+/- 200ms) within backoff calculation to prevent massive client stampedes upon backend server recovery.
- **LocalStorage Storage quota exceeded**: Ensure store validations drop outdated scans and enforce queue limits (cap at 5,000 entries max) to prevent browser quota crashes.

## Review Guidance

- Reviewer must verify that the WebSocket client delays progressive reconnect attempts correctly (1s, 2s, 4s, 8s, up to 30s max).
- Assert that modifying `pendingQueue` in Pinia updates LocalStorage keys instantly.

## Activity Log

- 2026-06-01T13:35:24Z – system – lane=planned – Prompt generated via /spec-kitty.tasks
- 2026-06-01T14:26:06Z – opencode – shell_pid=18424 – lane=doing – Assigned agent via workflow command
- 2026-06-01T14:28:43Z – opencode – shell_pid=18424 – lane=for_review – Ready for review: WebSocket, backoff reconnect loop, and LocalStorage check-in store.
- 2026-06-01T14:52:59Z – opencode – shell_pid=42048 – lane=doing – Started review via workflow command
- 2026-06-01T14:53:11Z – opencode – shell_pid=42048 – lane=done – Review passed: Resilient WebSocket clients, exponential backoff connection retries with jitter, and LocalStorage-synced Pinia queues are fully verified.
