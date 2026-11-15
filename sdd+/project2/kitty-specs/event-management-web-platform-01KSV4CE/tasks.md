# Work Package Task List: Event Management Web Platform

This document outlines the subtask index and detailed work packages for executing the Event Management Web Platform.

---

## Subtask Index

| ID | Description | Work Package | Parallel |
|---|---|---|---|
| T001 | Copy Project 1 Go REST API codebase into backend folder | WP01 | [P] |
| T002 | Create backend Dockerfile | WP01 | [P] |
| T003 | Configure docker-compose.yml with PostgreSQL and Backend services | WP01 | No |
| T004 | Configure database migrations hook on backend launch | WP01 | No |
| T005 | Execute local backend verification checks | WP01 | No |
| T006 | Create database migration schema for Categories | WP02 | No |
| T007 | Create database migration schema for Organizer verification | WP02 | No |
| T008 | Create database migration schema for Rich Event metadata | WP02 | No |
| T009 | Create database migration schema for Registration Presence status | WP02 | No |
| T010 | Execute database migrations end-to-end verification | WP02 | No |
| T011 | Implement User model role and status attributes in Go | WP03 | [D] |
| T012 | Implement JWT token role-validation mechanisms | WP03 | No | [D] |
| T013 | Implement Admin verification endpoints for Organizers | WP03 | No | [D] |
| T014 | Implement Admin CRUD Category handlers | WP03 | No | [D] |
| T015 | Add backend Admin validation tests | WP03 | No | [D] |
| T016 | Implement Organizer event CRUD logic | WP04 | No |
| T017 | Implement Organizer attendee roster and status actions | WP04 | No |
| T018 | Implement backend CSV generator and download route | WP04 | [P] |
| T019 | Add backend Organizer integration tests | WP04 | No |
| T020 | Implement email client wrapper using net/smtp | WP05 | [P] |
| T021 | Create Registration/Cancellation HTML template packages | WP05 | [P] |
| T022 | Integrate automated email dispatch hooks | WP05 | No |
| T023 | Append SMTP mail catcher wrapper to docker-compose | WP05 | No |
| T024 | Initialize Vue 3 Frontend app under /frontend | WP06 | [P] |
| T025 | Configure Tailwind CSS integrations and base properties | WP06 | [P] |
| T026 | Define high-contrast basic accessible focus rules | WP06 | No |
| T027 | Create multi-stage Frontend Dockerfile | WP06 | [P] |
| T028 | Map frontend development services inside docker-compose | WP06 | No |
| T029 | Implement Pinia store auth mechanisms | WP07 | No |
| T030 | Implement Vue Router and client guards | WP07 | No |
| T031 | Build accessible Organizer/Attendee sign-up forms | WP07 | No |
| T032 | Build accessible login portal views | WP07 | No |
| T033 | Add frontend validation visual feedback and focus states | WP07 | No |
| T034 | Create searchable/categorized Event Discovery grid | WP08 | No |
| T035 | Build detailed Event Info and seat counter cards | WP08 | No |
| T036 | Create Attendee tracking schedule views | WP08 | No |
| T037 | Incorporate visual status notifications for sign-ups | WP08 | No |
| T038 | Construct Organizer Event CRUD templates and forms | WP09 | No |
| T039 | Build interactive presence confirmation roster tables | WP09 | No |
| T040 | Integrate stats charts and CSV export triggers | WP09 | No |
| T041 | Introduce screen-reader table labels and focus controls | WP09 | No |
| T042 | Design Admin Organizer approval workspace | WP10 | No |
| T043 | Design Admin Category CRUD layout tables | WP10 | No |
| T044 | Design Admin high-level system-wide aggregate charts | WP10 | No |
| T045 | Perform cross-role end-to-end integration layout reviews | WP10 | No |

---

## Work Packages

### Work Package 1: Repository Layout & Backend Setup
*   **Goal:** Establish multi-component directory layout and verify baseline Project 1 backend execution in Docker.
*   **Priority:** High
*   **Test Criteria:** Backend compiles, launches, and successfully connects to the PostgreSQL container in Docker.
*   **Dependencies:** None

#### Included Subtasks
- [ ] T001 Copy Project 1 Go REST API codebase into backend folder (WP01)
- [ ] T002 Create backend Dockerfile (WP01)
- [ ] T003 Configure docker-compose.yml with PostgreSQL and Backend services (WP01)
- [ ] T004 Configure database migrations hook on backend launch (WP01)
- [ ] T005 Execute local backend verification checks (WP01)

---

### Work Package 2: Database Extensions & Migration Scripts
*   **Goal:** Expand the database schema with new relational models and data states.
*   **Priority:** High
*   **Test Criteria:** Migrations run without syntax errors and establish expected constraints in Postgres.
*   **Dependencies:** WP01

#### Included Subtasks
- [ ] T006 Create database migration schema for Categories (WP02)
- [ ] T007 Create database migration schema for Organizer verification (WP02)
- [ ] T008 Create database migration schema for Rich Event metadata (WP02)
- [ ] T009 Create database migration schema for Registration Presence status (WP02)
- [ ] T010 Execute database migrations end-to-end verification (WP02)

---

### Work Package 3: Backend Core Logic & Admin/Organizer Auth Guarding
*   **Goal:** Update user authorization models and add administrator management endpoints in the backend.
*   **Priority:** High
*   **Test Criteria:** Admins can query and change organizer status; JWT parses roles correctly.
*   **Dependencies:** WP02

#### Included Subtasks
- [x] T011 Implement User model role and status attributes in Go (WP03)
- [x] T012 Implement JWT token role-validation mechanisms (WP03)
- [x] T013 Implement Admin verification endpoints for Organizers (WP03)
- [x] T014 Implement Admin CRUD Category handlers (WP03)
- [x] T015 Add backend Admin validation tests (WP03)

---

### Work Package 4: Backend Event Lifecycle & Attendance Management
*   **Goal:** Add backend event lifecycles, roster management, and CSV downloads.
*   **Priority:** Medium
*   **Test Criteria:** Organizers can manage event states, log presence, and export a CSV payload.
*   **Dependencies:** WP03

#### Included Subtasks
- [ ] T016 Implement Organizer event CRUD logic (WP04)
- [ ] T017 Implement Organizer attendee roster and status actions (WP04)
- [ ] T018 Implement backend CSV generator and download route (WP04)
- [ ] T019 Add backend Organizer integration tests (WP04)

---

### Work Package 5: Backend SMTP Email Service
*   **Goal:** Implement automated email triggers for registration and cancellations.
*   **Priority:** Medium
*   **Test Criteria:** Email client can dispatch template emails caught successfully by Mailpit.
*   **Dependencies:** WP04

#### Included Subtasks
- [ ] T020 Implement email client wrapper using net/smtp (WP05)
- [ ] T021 Create Registration/Cancellation HTML template packages (WP05)
- [ ] T022 Integrate automated email dispatch hooks (WP05)
- [ ] T023 Append SMTP mail catcher wrapper to docker-compose (WP05)

---

### Work Package 6: Frontend Vue.js Scaffolding & Theme Integration
*   **Goal:** Create the Vue 3 + Tailwind CSS workspace and integrate basic accessibility themes.
*   **Priority:** High
*   **Test Criteria:** Vue application successfully boots in Dev mode and correctly applies base styles.
*   **Dependencies:** None

#### Included Subtasks
- [ ] T024 Initialize Vue 3 Frontend app under /frontend (WP06)
- [ ] T025 Configure Tailwind CSS integrations and base properties (WP06)
- [ ] T026 Define high-contrast basic accessible focus rules (WP06)
- [ ] T027 Create multi-stage Frontend Dockerfile (WP06)
- [ ] T028 Map frontend development services inside docker-compose (WP06)

---

### Work Package 7: Frontend State Management & Authentication Views
*   **Goal:** Add Pinia authorization states, client-side route guarding, and accessible sign-up pages.
*   **Priority:** High
*   **Test Criteria:** Attendees/Organizers can register/login, and routing guards block unauthorized pages.
*   **Dependencies:** WP06

#### Included Subtasks
- [ ] T029 Implement Pinia store auth mechanisms (WP07)
- [ ] T030 Implement Vue Router and client guards (WP07)
- [ ] T031 Build accessible Organizer/Attendee sign-up forms (WP07)
- [ ] T032 Build accessible login portal views (WP07)
- [ ] T033 Add frontend validation visual feedback and focus states (WP07)

---

### Work Package 8: Frontend Attendee Event Portal
*   **Goal:** Build accessible event catalogs, detail views, and user registration pages.
*   **Priority:** Medium
*   **Test Criteria:** Attendees can search and register for events and view their timeline.
*   **Dependencies:** WP07

#### Included Subtasks
- [ ] T034 Create searchable/categorized Event Discovery grid (WP08)
- [ ] T035 Build detailed Event Info and seat counter cards (WP08)
- [ ] T036 Create Attendee tracking schedule views (WP08)
- [ ] T037 Incorporate visual status notifications for sign-ups (WP08)

---

### Work Package 9: Frontend Organizer Dashboard
*   **Goal:** Build organizer event controls, check-in roster tables, and CSV exports.
*   **Priority:** Medium
*   **Test Criteria:** Organizers can create/edit events, toggle attendee presence, and trigger exports.
*   **Dependencies:** WP08

#### Included Subtasks
- [ ] T038 Construct Organizer Event CRUD templates and forms (WP09)
- [ ] T039 Build interactive presence confirmation roster tables (WP09)
- [ ] T040 Integrate stats charts and CSV export triggers (WP09)
- [ ] T041 Introduce screen-reader table labels and focus controls (WP09)

---

### Work Package 10: Frontend Administrator Workspace & Global Integration
*   **Goal:** Add administrator organizer review portals, categories, and final responsive checks.
*   **Priority:** Medium
*   **Test Criteria:** Administrators can manage categories, approve organizers, and verify desktop/mobile layouts.
*   **Dependencies:** WP09

#### Included Subtasks
- [ ] T042 Design Admin Organizer approval workspace (WP10)
- [ ] T043 Design Admin Category CRUD layout tables (WP10)
- [ ] T044 Design Admin high-level system-wide aggregate charts (WP10)
- [ ] T045 Perform cross-role end-to-end integration layout reviews (WP10)
