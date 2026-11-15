# Task List: Fix Event and User Registration Core

**Mission ID**: `01KS8SHC2ETY63EWAVDBNA8PP0`

This document defines the subtask index and individual work packages required to implement database migration fixes and secure user/event endpoints with domain validation.

---

## Subtask Index

| ID | Description | WP | Parallel |
| :--- | :--- | :--- | :--- |
| **T001** | Fix initial up-migration syntax in SQL scripts | WP01 | `[P]` |
| **T002** | Implement robust down-migration syntax for teardown | WP01 | `[P]` |
| **T003** | Verify migrations execution against uninitialized database | WP01 | |
| **T004** | Implement register domain service with security & uniqueness rules | WP02 | `[P]` |
| **T005** | Implement login validation and authentication logic | WP02 | `[P]` |
| **T006** | Implement and correct Register and Login handlers with error maps | WP02 | |
| **T007** | Write domain service tests for User model and authentication | WP02 | |
| **T008** | Implement Event CRUD domain business service validations | WP03 | `[P]` |
| **T009** | Add database repository queries supporting full event CRUD | WP03 | `[P]` |
| **T010** | Add search and query filtering support for Listing Events handler | WP03 | |
| **T011** | Implement event handlers verifying organizer identities | WP03 | |
| **T012** | Wire Event CRUD routes in the API Router | WP03 | |
| **T013** | Add domain service unit tests for event model and validations | WP03 | |
| **T014** | Implement registration and unregistration logic in domain service | WP04 | `[P]` |
| **T015** | Add domain validators for Published status, duplicates, and capacity | WP04 | `[P]` |
| **T016** | Complete Registration endpoints extracting Attendee identity from claims | WP04 | |
| **T017** | Connect registration endpoints to Router with authentication | WP04 | |
| **T018** | Write unit tests verifying registration capacity gates | WP04 | |

---

## WP01: Database Schema Migration Correction
* **Summary**: Resolve the up/down script overlap so tables persist across container launches.
* **Estimated Prompt Size**: ~250 lines
* **Included Subtasks**:
  - [x] T001 Fix initial up-migration syntax in SQL scripts
  - [x] T002 Implement robust down-migration syntax for teardown
  - [x] T003 Verify migrations execution against uninitialized database

---

## WP02: User Registration & Login Services
* **Summary**: Establish correct account creation, password hashing, and authentication services with proper error codes.
* **Estimated Prompt Size**: ~350 lines
* **Dependencies**: WP01
* **Included Subtasks**:
  - [x] T004 Implement register domain service with security & uniqueness rules
  - [x] T005 Implement login validation and authentication logic
  - [x] T006 Implement and correct Register and Login handlers with error maps
  - [x] T007 Write domain service tests for User model and authentication

---

## WP03: Event CRUD Domain & Handler Implementation
* **Summary**: Create fully functional Event management endpoints with search filters and organizer identity locks.
* **Estimated Prompt Size**: ~450 lines
* **Dependencies**: WP02
* **Included Subtasks**:
  - [x] T008 Implement Event CRUD domain business service validations
  - [x] T009 Add database repository queries supporting full event CRUD
  - [x] T010 Add search and query filtering support for Listing Events handler
  - [x] T011 Implement event handlers verifying organizer identities
  - [x] T012 Wire Event CRUD routes in the API Router
  - [x] T013 Add domain service unit tests for event model and validations

---

## WP04: Event Registration & Capacity Gates
* **Summary**: Complete registration/unregistration endpoint mappings with strict published-status and capacity validations.
* **Estimated Prompt Size**: ~400 lines
* **Dependencies**: WP03
* **Included Subtasks**:
  - [x] T014 Implement registration and unregistration logic in domain service
  - [x] T015 Add domain validators for Published status, duplicates, and capacity
  - [x] T016 Complete Registration endpoints extracting Attendee identity from claims
  - [x] T017 Connect registration endpoints to Router with authentication
  - [x] T018 Write unit tests verifying registration capacity gates
