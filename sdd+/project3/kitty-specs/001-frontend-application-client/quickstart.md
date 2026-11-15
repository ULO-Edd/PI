# Quickstart Guide: Frontend Application Client

This guide explains how to spin up, develop, and test the Frontend Application Client locally.

## Prerequisites

- **Node.js**: v18.x or higher  
- **npm**: v9.x or higher

## Setup Instructions

1. Navigate to the frontend directory:
   ```bash
   cd frontend
   ```

2. Install dependencies:
   ```bash
   npm install
   ```

3. Spin up the development server locally:
   ```bash
   npm run dev
   ```
   *The application will boot and be accessible at `http://localhost:5173/` by default.*

## Code Quality & Verification

- **Lint and Style Verification**:
  ```bash
  npm run lint
  ```

- **Unit Testing (Vitest)**:
  ```bash
  npm run test:unit
  ```

- **E2E and UI Accessibility Testing (Playwright)**:
  ```bash
  npx playwright test
  ```
