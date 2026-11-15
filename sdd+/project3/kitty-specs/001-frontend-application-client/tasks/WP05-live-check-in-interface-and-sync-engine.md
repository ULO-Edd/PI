---
work_package_id: WP05
title: Live Check-In Interface & Sync Engine
lane: "done"
dependencies: [WP04]
base_branch: 001-frontend-application-client-WP04
base_commit: 7d527b27e0c34ecd62f4238f0284373b87e04400
created_at: '2026-06-01T14:36:29.075729+00:00'
subtasks:
- T014
- T015
- T016
- T017
phase: Phase 5 - Interactive Client & Synchronisation
assignee: ''
agent: "opencode"
shell_pid: "8248"
review_status: "approved"
reviewed_by: "Tiago Ramos"
history:
- timestamp: '2026-06-01T13:35:24Z'
  lane: planned
  agent: system
  shell_pid: ''
  action: Prompt generated via /spec-kitty.tasks
requirement_refs:
- FR-007
- FR-010
---

# Work Package Prompt: WP05 – Live Check-In Interface & Sync Engine

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

- Build the interactive, live check-in page displaying event rosters, search parameters, and scan triggers.
- Create a persistent, non-blocking warning banner notifying users instantly during network disconnections.
- Build an automatic synchronization engine that flushes queued offline scans sequentially to the backend service upon reconnecting.
- Write E2E Playwright simulations and Vitest suites verifying the offline check-in and auto-sync cycle.

## Context & Constraints

- Refers to `kitty-specs/001-frontend-application-client/spec.md` (defining live check-in, real-time sync, and non-blocking warning banner).
- Requires a base branch with WP04:
  ```bash
  spec-kitty implement WP05 --base WP04
  ```

## Subtasks & Detailed Guidance

### Subtask T014 - Build Live Check-In page layout
- **Purpose**: Provide check-in managers with a reactive attendee list and manual check-in scan controls.
- **Steps**:
  1. Build `/frontend/src/pages/LiveCheckIn.vue` mapping to private routes.
  2. Implement local search and filter to locate specific attendees on the roster.
  3. Include a scan toggle or trigger mimicking a barcode/QR scan:
     - When clicked, append a scan entry to the check-in queue in the Pinia store.
     - Play simple success/error sound and update attendee status locally.
- **Files**:
  - `/frontend/src/pages/LiveCheckIn.vue`
- **Parallel?**: Yes. Can be built concurrently with sync logic (T016).

### Subtask T015 - Build persistent warning banner
- **Purpose**: Alert users immediately (< 500ms) when the WebSocket connection is lost.
- **Steps**:
  1. Create `/frontend/src/components/NotificationBanner.vue`.
  2. Watch the `isConnected` state from the check-in Pinia store.
  3. Render a non-blocking, prominent warning banner (e.g., orange or red alert bar) at the top of the interface when `isConnected` is false.
  4. Display status: "Offline - Scans are being queued locally".
  5. Automatically dismiss banner immediately when `isConnected` becomes true.
- **Files**:
  - `/frontend/src/components/NotificationBanner.vue`
- **Parallel?**: Yes. Can proceed once components layout is complete.

### Subtask T016 - Implement queue auto-sync engine
- **Purpose**: Flush queued scans automatically and sequentially to the backend without user intervention.
- **Steps**:
  1. Add a sync watcher inside the check-in Pinia store watching the `isConnected` state.
  2. When connection changes from false to true:
     - Check if any entries exist in the LocalStorage-backed `pendingQueue`.
     - Flush entries sequentially (one by one or in small serial batches) to `/api/registrations/sync` to maintain chronological sequence.
     - For each successfully synced entry, update its local status, and remove it from the `pendingQueue`.
- **Files**:
  - `/frontend/src/stores/checkIn.ts`
- **Parallel?**: No. Requires integration within Pinia check-in store.

### Subtask T017 - Write E2E offline resilience scripts
- **Purpose**: Validate end-to-end resilience during simulated disconnections.
- **Steps**:
  1. Write `/frontend/tests/e2e/resilience.spec.ts` using Playwright:
     - Open the Live Check-In view.
     - Block network connections or trigger offline mode.
     - Verify the persistent banner appears within 500ms.
     - Trigger multiple check-in scans. Assert they are added to LocalStorage.
     - Restore connection.
     - Verify that the banner disappears, and local storage is flushed.
- **Files**:
  - `/frontend/tests/e2e/resilience.spec.ts`
- **Parallel?**: No. Runs as a final integration checkpoint.

## Risks & Mitigations

- **Double submissions during rapid reconnects**: Track client scan timestamps and correlation IDs to allow backend services to deduplicate incoming scans safely.
- **Partial sync failure**: If the network disconnects midway through a sync flush, halt the sequential flush, preserve the remaining items, and resume attempts on the next reconnect event.

## Review Guidance

- Reviewer must verify that warning banners render within 500ms of WebSocket disconnections.
- Ensure all queued scans are safely flushed sequentially upon connection recovery.

## Activity Log

- 2026-06-01T13:35:24Z – system – lane=planned – Prompt generated via /spec-kitty.tasks
- 2026-06-01T14:36:31Z – opencode – shell_pid=29164 – lane=doing – Assigned agent via workflow command
- 2026-06-01T14:39:37Z – opencode – shell_pid=29164 – lane=for_review – Ready for review: Interactive live check-in pages, warning banners, auto-sync queues, and Vitest offline synchronisation logic tests.
- 2026-06-01T14:53:21Z – opencode – shell_pid=8248 – lane=doing – Started review via workflow command
- 2026-06-01T14:53:34Z – opencode – shell_pid=8248 – lane=done – Review passed: Interactive check-in lists, connection warning banners, sequential queue auto-sync flushes, and associated test coverage are fully verified.
