# Research & Decisions: Event Management RESTful API

This document records the key technical decisions made during the planning phase.

## Decision: Technology Stack

- **Decision**: The project will be built using **Python 3.11**, the **FastAPI** framework, and **SQLite** as the database.
- **Rationale**: This stack was chosen based on the planning interrogation. FastAPI is a modern, high-performance web framework for Python that is easy to learn and use. SQLite is a lightweight, serverless, self-contained SQL database engine that is ideal for development and small to medium-sized applications. Python is a widely-used language with a rich ecosystem of libraries.
- **Alternatives considered**:
    - **Node.js/Express**: A popular choice for building APIs, but the team preferred Python.
    - **Java/Spring Boot**: A robust and widely-used framework, but considered to have a steeper learning curve for this project's scope.
    - **PostgreSQL/MySQL**: More powerful and scalable databases, but SQLite is sufficient for the current requirements and simpler to set up.
