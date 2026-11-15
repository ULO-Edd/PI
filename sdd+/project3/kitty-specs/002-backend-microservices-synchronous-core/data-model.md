# Data Model: Backend Microservices Synchronous Core

This document defines the database schemas, entity structures, and relationships for the isolated relational databases of each microservice.

## 1. Identity Service Database

Stores user profiles and security definitions.

### Table: `Users`

| Column Name | Data Type | Constraints | Description |
|---|---|---|---|
| `Id` | Guid | Primary Key | Unique user identifier. |
| `Email` | String (255) | Required, Unique, Email format | User login email. |
| `PasswordHash` | String (500) | Required | Cryptographically hashed password. |
| `Role` | String (50) | Required (Admin, Organiser, Attendee) | Role-based access level. |
| `CreatedAt` | DateTime | Required | Account registration timestamp. |

---

## 2. Events Service Database

Stores event catalogues, venues, and live ticket inventories.

### Table: `Events`

| Column Name | Data Type | Constraints | Description |
|---|---|---|---|
| `Id` | Guid | Primary Key | Unique event identifier. |
| `Title` | String (100) | Required, Max Length 100 | Event title. |
| `Description` | String (2000) | Optional | Event detailed description. |
| `DateTime` | DateTime | Required, Must be in future | Event start timestamp. |
| `Location` | String (255) | Required | Venue name or virtual access link. |
| `Capacity` | Integer | Required, Minimum 1 | Maximum registrations allowed. |
| `AvailableTickets`| Integer | Required, 0 <= Available <= Capacity | Remaining ticket inventory. |
| `OrganiserId` | Guid | Required | Reference to creator organiser in Identity database. |

---

## 3. Registrations Service Database

Stores attendee ticket bookings and transactions. Completely isolated from the Events database.

### Table: `Registrations`

| Column Name | Data Type | Constraints | Description |
|---|---|---|---|
| `Id` | Guid | Primary Key | Unique registration identifier. |
| `EventId` | Guid | Required | Reference to event ID (validated via Events Service). |
| `AttendeeId` | Guid | Required | Reference to attendee user ID (validated via Identity Service). |
| `Status` | String (50) | Required (Pending, Confirmed, Cancelled) | Active state of the ticket booking. |
| `BookedAt` | DateTime | Required | Registration timestamp. |
| `CorrelationId` | String (100) | Required | Correlation ID used for transaction lineage. |

### Registration State Transitions

```
[Pending] -> [Confirmed] (Synchronous capacity deduction succeeds)
[Pending] -> [Cancelled] (Deduction fails or transaction rolls back)
```
