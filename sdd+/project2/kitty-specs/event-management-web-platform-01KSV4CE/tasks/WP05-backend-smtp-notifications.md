---
work_package_id: WP05
title: Backend SMTP Email Service
dependencies:
- WP04
requirement_refs:
- C-002
- FR-009
planning_base_branch: feature/sdd-plus-project2
merge_target_branch: feature/sdd-plus-project2
branch_strategy: Planning artifacts for this feature were generated on feature/sdd-plus-project2. During /spec-kitty.implement this WP may branch from a dependency-specific base, but completed changes must merge back into feature/sdd-plus-project2 unless the human explicitly redirects the landing branch.
subtasks:
- T020
- T021
- T022
- T023
phase: Phase 5 - SMTP Messaging Integration
assignee: ''
agent: "opencode"
shell_pid: "1604"
history:
- timestamp: '2026-05-30T00:31:40Z'
  agent: system
  action: Prompt generated via /spec-kitty.tasks
authoritative_surface: backend/internal/email/
execution_mode: code_change
owned_files:
- backend/internal/email/**
tags: []
---

# Work Package Prompt: WP05 – Backend SMTP Email Service

## Objective
Integrate net/smtp email services, draft core HTML messaging templates, and trigger notifications during lifecycle states.

---

## Subtask Details

### Subtask T020: Implement net/smtp Email Client
*   **Purpose:** Build core SMTP dispatch interface reading parameters from environment.
*   **Steps:**
    1. Write client implementation in Go.
*   **Files:** `backend/internal/email/client.go`
*   **Validation:** Dispatch mock connections without connection timeouts.

### Subtask T021: Create Messaging HTML Templates
*   **Purpose:** Build standard rich responsive HTML messaging templates.
*   **Steps:**
    1. Create Registration Confirmation template.
    2. Create Event Cancellation warning template.
    3. Create Organizer Verification approval template.
*   **Files:** `backend/internal/email/templates.go`
*   **Validation:** Templates render correctly in HTML clients.

### Subtask T022: Integrate Event Email Triggers
*   **Purpose:** Hook automated alerts directly to backend event loops.
*   **Steps:**
    1. Add async email dispatch trigger on event cancellation and successful attendee sign-up.
*   **Files:** `backend/internal/email/hooks.go`
*   **Validation:** Registration triggers confirm emails in async background routines.

### Subtask T023: Mail Catcher Compose Setup
*   **Purpose:** Enable local webmail container capture.
*   **Steps:**
    1. Append Mailpit/MailHog container mapping to `docker-compose.yml`.
*   **Files:** `docker-compose.yml` (Note: docker-compose modifications are integrated safely)
*   **Validation:** Mail catcher dashboard shows all outbound emails.

## Activity Log

- 2026-05-30T21:07:01Z – opencode – shell_pid=15300 – Started implementation via action command
- 2026-05-30T21:08:16Z – opencode – shell_pid=15300 – SMTP email client, HTML templates, and Mailpit integration finalized. Async hooks prepared for lifecycle events.
- 2026-05-30T21:08:25Z – opencode – shell_pid=1604 – Started review via action command
- 2026-05-30T21:08:36Z – opencode – shell_pid=1604 – Review passed: SMTP client implemented with net/smtp, HTML templates created, and Mailpit successfully added to docker-compose orchestration.
