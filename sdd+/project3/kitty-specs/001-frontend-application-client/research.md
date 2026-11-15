# Research Document: Frontend Application Client

This research document analyzes key technical decisions, architectures, and alternatives for the client-side user interface.

## 1. Asynchronous Offline Queue Persistence

- **Decision**: Persist Pinia store offline scan queues to `LocalStorage`, utilizing `navigator.onLine` and WebSocket state triggers combined with an exponential backoff reconnect policy.
- **Rationale**: LocalStorage is simple, synchronous, and universally supported across browser clients. It provides excellent reliability for storing up to 5,000 pending scans without any file storage overhead. Upon reconnection, an automatic flush routine synchronizes the queue sequentially to guarantee transactional integrity.
- **Alternatives Considered**:
  - **IndexedDB**: Rejected due to asynchronous callback API complexity which increases development and testing overhead. LocalStorage easily meets the capacity constraints.
  - **Service Workers**: Rejected for basic queue synchronization. While excellent for complete offline PWA assets, in-app WebSocket state management is more predictable and easier to integrate directly with Pinia.

## 2. Secure Route Protection and Authentication State

- **Decision**: Secure HTTP-only cookies for JWT storage with Vue Router navigation guards.
- **Rationale**: Storing JWT tokens in HTTP-only cookies prevents client-side JavaScript access, virtually eliminating Cross-Site Scripting (XSS) extraction risks. Vue Router navigation guards evaluate authenticated state (via session checks) before allowing the user to enter secured organiser views.
- **Alternatives Considered**:
  - **LocalStorage Token Storage**: Rejected due to vulnerability to XSS attacks where malicious scripts could steal tokens.
  - **SessionStorage**: Rejected because it is tab-scoped, breaking multi-tab navigation and user sessions.

## 3. WCAG 2.1 AA Accessibility Validation

- **Decision**: Leverage semantic HTML, ARIA annotations built into PrimeVue UI components, styled with Tailwind CSS, and validated using automated Playwright accessibility sweeps.
- **Rationale**: PrimeVue components are engineered with built-in screen reader support and keyboard-accessible navigation (e.g., proper tab indices, roles). Tailwind CSS manages focus rings and contrast ratios cleanly. Playwright integrates directly with standard accessibility audit tools for automated CI gates.
- **Alternatives Considered**:
  - **Custom HTML Components**: Rejected due to high development effort to build keyboard accessibility and focus management from scratch.
  - **Vuetify**: Rejected as PrimeVue is the chosen library in the project constitution.
