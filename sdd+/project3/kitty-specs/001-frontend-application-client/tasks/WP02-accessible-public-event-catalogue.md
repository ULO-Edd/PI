---
work_package_id: WP02
title: Accessible Public Event Catalogue
lane: "done"
dependencies: [WP01]
base_branch: 001-frontend-application-client-WP01
base_commit: 2b072ef1391046fe7e1dadb57a2004e43ed8848e
created_at: '2026-06-01T14:13:54.519689+00:00'
subtasks:
- T005
- T006
- T007
phase: Phase 2 - Public Interface
assignee: ''
agent: "opencode"
shell_pid: "18392"
review_status: "approved"
reviewed_by: "Tiago Ramos"
history:
- timestamp: '2026-06-01T13:35:24Z'
  lane: planned
  agent: system
  shell_pid: ''
  action: Prompt generated via /spec-kitty.tasks
requirement_refs:
- FR-001
- FR-002
---

# Work Package Prompt: WP02 – Accessible Public Event Catalogue

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

- Create a highly responsive landing page housing the public-facing Event Catalogue components.
- Enable attendees to search, filter, and discover upcoming events efficiently.
- Enforce strict WCAG 2.1 AA accessibility guidelines, achieving zero automated accessibility errors.
- Write E2E and automated accessibility audit scripts utilizing Playwright.

## Context & Constraints

- Refers to `kitty-specs/001-frontend-application-client/spec.md` (defining Landing Page, search, and WCAG 2.1 AA compliance).
- Requires a base branch with WP01:
  ```bash
  spec-kitty implement WP02 --base WP01
  ```

## Subtasks & Detailed Guidance

### Subtask T005 - Create accessible Event Catalogue components
- **Purpose**: Create high-quality, semantic event display components.
- **Steps**:
  1. Create `components/EventCard.vue` representing a single event card:
     - Wrap the layout inside an `<article>` tag.
     - Ensure semantic headings are ordered logically (e.g., `<h3>` for event title).
     - Provide descriptive alt-text placeholders for all visual details.
  2. Create `components/EventCatalogue.vue` to render a grid list of cards:
     - Wrap grid in `<section>` with proper `aria-label="Event Catalogue"`.
     - Implement proper CSS Grid/Flexbox layouts that are fully responsive from mobile up to desktop.
- **Files**:
  - `/frontend/src/components/EventCard.vue`
  - `/frontend/src/components/EventCatalogue.vue`
- **Parallel?**: No. Components must exist before pages integration.

### Subtask T006 - Integrate Landing Page search & filtering
- **Purpose**: Map event models to the landing page and allow real-time client-side search.
- **Steps**:
  1. Implement `pages/LandingPage.vue` fetching events from the API client.
  2. Add search query input styled with Tailwind, including explicit `<label>` tags and `aria-describedby` helpers.
  3. Integrate client-side filtering matching event titles and descriptions.
  4. Ensure all list items can be navigated and activated via keyboard (Tab and Enter key).
- **Files**:
  - `/frontend/src/pages/LandingPage.vue`
- **Parallel?**: Yes. Can proceed once components layout (T005) is ready.

### Subtask T007 - Write Playwright accessibility checks
- **Purpose**: Automate testing of WCAG 2.1 AA criteria using axe-core within E2E runs.
- **Steps**:
  1. Add Playwright E2E accessibility script `tests/e2e/accessibility.spec.ts`.
  2. Configure tests to boot local app, navigate to Landing Page, and trigger axe-core accessibility analysis.
  3. Verify that zero WCAG 2.1 AA violations exist.
  4. Write assertions checking keyboard navigation paths (tabbing through elements and selecting an event).
- **Files**:
  - `/frontend/tests/e2e/accessibility.spec.ts`
- **Parallel?**: Yes. Can proceed concurrently once Landing Page structure is ready.

## Risks & Mitigations

- **Keyboard Trap risk**: Ensure dropdowns, modals, and search bars do not trap keyboard focus, allowing users to navigate cleanly out of inputs.

## Review Guidance

- Reviewer must verify that tab navigation traverses the event cards sequentially.
- Verify axe-core automated audits report 0 errors for the public pages.

## Activity Log

- 2026-06-01T13:35:24Z – system – lane=planned – Prompt generated via /spec-kitty.tasks
- 2026-06-01T14:13:57Z – opencode – shell_pid=8676 – lane=doing – Assigned agent via workflow command
- 2026-06-01T14:16:30Z – opencode – shell_pid=8676 – lane=for_review – Ready for review: Accessible Event Catalogue, landing search query filters, and automated Playwright E2E accessibility checks.
- 2026-06-01T14:52:16Z – opencode – shell_pid=18392 – lane=doing – Started review via workflow command
- 2026-06-01T14:52:29Z – opencode – shell_pid=18392 – lane=done – Review passed: Accessible public event cards, catalogues, and automated Playwright E2E accessibility checks are fully verified.
