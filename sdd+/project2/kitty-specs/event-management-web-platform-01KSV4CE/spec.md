# Event Management Web Platform Specification

## 1. Overview & Intent Summary
The Event Management Web Platform is an accessible, responsive web application that extends the existing event publication, registration, and management REST API (Project 1) into a full web-based solution. The platform serves the ESTG / IPLeiria community by enabling three primary roles: Attendees (who discover, register for, and track events), Organizers (who manage event lifecycles, track attendee presence, and analyze event metrics), and Administrators (who approve organizers, manage categories, and monitor platform-wide metrics). 

The platform guarantees accessibility (WCAG 2.1 AA compliance) and a responsive experience across all modern desktop, tablet, and mobile devices.

---

## 2. Success Criteria
*   **SC-1 (Notification Timeliness):** 100% of event registration, cancellation, and modification notifications are successfully sent to users' emails within 5 minutes of the triggering action.
*   **SC-2 (Keyboard-Only Usability):** 100% of interactive user interface components (including navigation, forms, and event lists) are navigable and operable using standard keyboard keys (Tab, Enter, Space) with a visible keyboard focus state.
*   **SC-3 (Multi-Device Visual Integrity):** The platform renders correctly on mobile (320px+), tablet (768px+), and desktop (1024px+) viewports, with no horizontal scrollbars on standard viewports, text overlapping, or truncated functional elements.
*   **SC-4 (Access/Load Performance):** Pages and dashboards must render and become interactive within 2 seconds under normal 3G/4G or broadband internet connections.

---

## 3. User Roles & Actors
*   **Attendee:** A student, teacher, or staff member of ESTG / IPLeiria who registers via an institutional or external email. They can search/filter events, register, view their past/upcoming events, and receive email updates.
*   **Organizer:** An event host who can draft, publish, edit, or cancel events, manage attendance (marking present/no-show), view statistics, and export attendee rosters. Organizers require administrative approval before they can publish events.
*   **Administrator:** Platform administrator who oversees the ecosystem. They approve or suspend organizers, manage global event categories, and monitor platform-wide metrics and statistics.

---

## 4. User Scenarios

### Scenario 1: Attendee Self-Registration & Event Discovery
*   **Actor:** Attendee
*   **Action:** The attendee registers with an institutional email, logs in, and lands on the homepage. They see a list of upcoming events categorized and searchable. They filter by "Technology," click on a workshop, view details, and register. 
*   **Outcome:** They immediately see their registration status, receive a confirmation email, and the workshop is added to their personal "Tracked Events" list.

### Scenario 2: Organizer Registration & Admin Approval
*   **Actor:** Organizer & Administrator
*   **Action:** A user registers as an organizer. The system places their account in a pending state and prevents event publication. An Administrator logs into the admin portal, reviews the list of pending organizers, and approves their application.
*   **Outcome:** The organizer receives an email approval notification, and their dashboard is unlocked, allowing them to publish events.

### Scenario 3: Attendance Tracking & CSV Roster Export
*   **Actor:** Organizer
*   **Action:** After an event ends, the organizer opens their Event Dashboard, reviews the registered attendee roster, and marks who attended ("Present") and who missed ("No-Show"). They then click "Export Attendee List."
*   **Outcome:** The attendance status is updated in the database, updating the fill and no-show statistics, and a CSV file containing names, emails, registration dates, and attendance states is downloaded to their local machine.

---

## 5. Key Entities & Data Models
*   **User:**
    *   Unique Identifier (ID)
    *   Name & Email (Institutional or External)
    *   Role (Attendee, Organizer, Administrator)
    *   Account Status (Pending, Approved, Suspended)
*   **Event:**
    *   Unique Identifier (ID)
    *   Title & Detailed Description
    *   Organizer (User ID reference)
    *   Category (Category ID reference)
    *   Maximum Capacity
    *   Start & End Timestamps
    *   Status (Draft, Published, Cancelled)
    *   Rich Metadata (Contact email, Location/Virtual Link, Additional notes)
*   **Registration:**
    *   Unique Identifier (ID)
    *   Event (Event ID reference)
    *   Attendee (User ID reference)
    *   Registration Timestamp
    *   Attendance Status (Registered, Present, No-Show, Cancelled)
*   **Category:**
    *   Unique Identifier (ID)
    *   Name & Description

---

## 6. Functional Requirements

| ID | Title | Description | Status |
|---|---|---|---|
| FR-001 | Attendee Registration | Attendees can self-register using institutional or external email addresses. | Approved |
| FR-002 | Organizer Registration | Organizers can register on the platform but are created in a pending state. | Approved |
| FR-003 | Organizer Approval | Administrators can view pending organizers and approve or suspend their accounts. | Approved |
| FR-004 | Event Discovery | Attendees can browse, search, and filter published events by category, date, and availability. | Approved |
| FR-005 | Event Registration | Attendees can register for an event if capacity permits, reducing available slots. | Approved |
| FR-006 | Registration Tracking | Attendees can view a personal list of upcoming and past registered events. | Approved |
| FR-007 | Event Creation & Editing | Approved organizers can draft, edit, and publish events with rich metadata (name, description, category, capacity, dates). | Approved |
| FR-008 | Event Cancellation | Organizers can cancel a published event. This changes its status to Cancelled and disables registrations. | Approved |
| FR-009 | Email Notifications | The system triggers automated email notifications to attendees for registration confirmations, event cancellations, and updates. | Approved |
| FR-010 | Attendee Presence | Organizers can view the attendee roster for their events and mark individuals as present or no-show. | Approved |
| FR-011 | Attendee CSV Export | Organizers can export the attendee registration and attendance list for an event to a CSV file. | Approved |
| FR-012 | Organizer Dashboard Statistics | Organizers can view statistics including registration trends over time, fill rates, and no-show rates. | Approved |
| FR-013 | Admin Category Management | Administrators can create, edit, and delete event categories. | Approved |
| FR-014 | Admin Global Dashboard | Administrators can view platform-wide statistics including total users, active organizers, total events, and global fill rate. | Approved |

---

## 7. Non-Functional Requirements

| ID | Title | Description | Status |
|---|---|---|---|
| NFR-001 | Responsive Layout | The user interface must adjust automatically to desktop (1024px+), tablet (768px-1023px), and mobile (320px-767px) viewports with no horizontal scroll or layout breaks. | Approved |
| NFR-002 | Contrast Accessibility | All text and functional elements must achieve a minimum contrast ratio of 4.5:1 for normal text and 3:1 for large text to meet WCAG 2.1 AA standards. | Approved |
| NFR-003 | Keyboard Navigation | All interactive components, forms, and navigation menus must be accessible via standard keyboard inputs (Tab, Shift+Tab, Enter, Spacebar) with highly visible focus states. | Approved |
| NFR-004 | Semantic HTML | Frontend markup must utilize native semantic elements (e.g., `<main>`, `<nav>`, `<header>`, `<button>`, `<fieldset>`) to support screen readers. | Approved |

---

## 8. Constraints

| ID | Title | Description | Status |
|---|---|---|---|
| C-001 | Backend Integration | The frontend application must utilize the existing Go REST API (Project 1) for all data persistence and authentication verification. | Approved |
| C-002 | Email Notification Infrastructure | Email notification system must support standard SMTP transport protocols without requiring external proprietary messaging dependencies. | Approved |
| C-003 | Docker Deployment | The web application and backend components must be fully buildable and deployable via Docker and Docker Compose. | Approved |

---

## 9. Assumptions
*   **A-1:** The existing Go REST API (Project 1) exposes or will be modified to expose endpoints for organizer approvals, role-based checks, email dispatch, categories, statistics, and presence tracking.
*   **A-2:** SMTP server configuration will be provided via environment variables in the Docker Compose setup.
*   **A-3:** Standard modern web browsers (Chrome, Firefox, Safari, Edge) are used for user access.

---

## 10. Risks & Mitigations
*   **R-1:** High concurrent registrations for a popular event causing overallocation.
    *   *Mitigation:* The backend must use transactional checks to ensure registration count does not exceed capacity.
*   **R-2:** Email notification delivery failures due to SMTP server issues.
    *   *Mitigation:* Implement asynchronous queueing for email dispatches and display appropriate warnings in logs if an email fails to send.
