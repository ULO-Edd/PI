# Research & Technical Analysis: Event Management Web Platform
*Path: [kitty-specs/event-management-web-platform-01KSV4CE/research.md](research.md)*

This document summarizes the technical research and architecture recommendations supporting the implementation plan.

---

## 1. Backend Migration & Integration
*   **Decision:** Copy the complete codebase of Project 1 to the `backend/` directory of `project2`.
*   **Rationale:** The existing Go REST API has fully validated database schemas, migrations, Swagger definitions, and JWT structures. Copying it to `backend/` preserves all tested logic and serves as a solid foundation.
*   **Modifications Required:**
    1.  Update the `go.mod` module path to reflect its nested location if necessary, or preserve it as standalone under `/backend`.
    2.  Add database migrations to support additional schema additions:
        *   User table updates (Status: `pending`, `approved`, `suspended`).
        *   Event table updates (Category references, Status: `draft`, `published`, `cancelled`).
        *   Registration table updates (Status: `registered`, `present`, `no-show`, `cancelled`).

---

## 2. SMTP Notification Infrastructure
*   **Decision:** Use the Go standard library `net/smtp` package to handle email notifications.
*   **SMTP Configuration:**
    *   Configuration parameters (SMTP Host, Port, Username, Password, Sender Name) will be loaded from environment variables (`.env`) via Docker.
    *   For the local developer environment, we will use a lightweight mail catcher service (e.g., Mailpit or MailHog) in the `docker-compose.yml` file to capture and preview all outgoing emails.
*   **Notification Templates:** Simple HTML templates will be built in the backend for registration confirmation, cancellation updates, and administrator approvals.

---

## 3. Role-Based Vue.js Client Routing
*   **Decision:** Build a Vue 3 Single Page Application (SPA) utilizing Vue Router with navigation guards and Pinia for centralized state management.
*   **Layout Framework:**
    *   Use a dynamic layout wrapper (`App.vue` or custom wrappers) that reads the active user's role from the Pinia store.
    *   **Attendee Layout:** Responsive layout with user-centric navigation (Discover Events, My Registrations, Profile).
    *   **Organizer / Admin Layout:** Sidebar-based dashboard layout optimizing dense data displays, table lists, and analytics widgets.
*   **Security Guarding:** Vue Router's `beforeEach` navigation guards will inspect the active JWT in the Pinia state to prevent unauthorized users from manually navigating to administrative routes.

---

## 4. Accessibility Compliance (WCAG 2.1 AA)
*   **Decision:** Implement semantic HTML5 elements alongside Tailwind's built-in focus and state utilities.
*   **Practices:**
    1.  **Semantic Elements:** Use `<header>`, `<nav>`, `<main>`, `<section>`, `<aside>`, and `<footer>` rather than general `div` tags.
    2.  **Color Contrast:** Tailwind utility combinations will target a minimum WCAG AA contrast ratio of 4.5:1. 
    3.  **Keyboard Navigation:** All interactive elements (buttons, inputs, links, cards) will support keyboard activation. Add standard custom Tailwind styles to ensure that keyboard `:focus-visible` triggers clear, outline rings.
    4.  **Form Labels:** Ensure every input element has an associated `<label>` or `aria-label` description.
