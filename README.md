# Spec-Driven Development with Spec-Kitty

**Evaluating Spec-Driven Development (SDD) against "vibe coding" for AI-assisted software development.**

This is the source code for the final-year project *"Spec-Driven Development – Module 2: Spec-Kitty"*, Bachelor's degree in Computer Engineering, School of Technology and Management (ESTG), Polytechnic of Leiria, June 2026.

## What this project is about

AI coding agents can produce working software quickly. When they only get loose prompts (*vibe coding*), though, the result is often incomplete, hard to check, or drifts away from the requirements. **Spec-Driven Development** reverses this: specifications, plans and task breakdowns become the main artefacts, and code is generated and reviewed against them.

This project uses **[Spec-Kitty](https://pypi.org/project/spec-kitty-cli/)**, an SDD framework that drives coding agents through a structured, kanban-based workflow with git-worktree isolation. We used it to build the same domain at three levels of complexity and measured the cost of doing so.

All applications implement one domain: an **Event Management Solution for ESTG / IPLeiria**, covering event publication, user registration and event management.

### The Spec-Kitty workflow

```
charter ─► specify ─► plan ─► tasks ─► implement ─► review ─► accept ─► merge
 (once)       │                            ▲           │
              │                            └─ rejected ┘
              └──────────── next feature / iteration ◄────────────────┘
```

| Step          | What happens                                                                                          |
|---------------|-------------------------------------------------------------------------------------------------------|
| **charter**   | An interview produces `charter.md`, the project's architectural "DNA" (stack, quality gates, constraints). Every later step reads it. |
| **specify**   | A feature description becomes a structured spec under `kitty-specs/<feature>/`, in its own worktree.  |
| **plan**      | Produces the implementation plan, research notes, data model, contracts (OpenAPI) and a quickstart.   |
| **tasks**     | Splits the plan into **work packages** (WP01, WP02...), each with its own prompt file.                |
| **implement** | The agent implements one WP at a time, moving it through the lanes `planned → doing → review → done`. |
| **review**    | Each WP is checked against its spec. Flawed WPs go back to `planned`.                                 |
| **accept**    | Final validation of the whole feature.                                                                |
| **merge**     | The feature worktree is merged into the main branch and cleaned up.                                   |

You can browse the specs, plans, task lists and WP prompts for each project in its `kitty-specs/` folder.

## Repository layout

```
.
├── vibe/                        # Baseline: vibe coding (no SDD)
│   ├── vibe-project1.1/         # Project 1 – Java / Spring Boot (Gemini CLI)
│   └── vibe-project1.2/         # Project 1 – Node.js / Express / Prisma
│
├── sdd/                         # SDD with minimal preparation
│   └── Project1.2-app/          # Project 1.2 – Python / FastAPI
│                                #   given only the raw project statement (no charter)
│
└── sdd+/                        # SDD with good practices
    ├── project1/                # Project 1.3 – Go REST API (charter interview, MCP, prompt engineering)
    ├── project2/                # Project 2  – Go API + Vue 3 web platform
    └── project3/                # Project 3  – .NET microservices + RabbitMQ + Vue 3
```

### The projects

| Project | Approach | Stack | Scope |
|---|---|---|---|
| **1.1** `vibe/vibe-project1.1` | Vibe coding | Java, Spring Boot, H2 | Basic events REST API. Requirements pasted into the agent with little guidance. |
| `vibe/vibe-project1.2` | Vibe coding | Node.js, Express, Prisma (SQLite) | Another vibe-coded take on the Project 1 API. |
| **1.2** `sdd/Project1.2-app` | SDD (raw statement) | Python, FastAPI, SQLAlchemy, JWT | Events and users REST API. A "brute-force" Spec-Kitty run, with all WPs reviewed together. |
| **1.3** `sdd+/project1` | SDD + best practices | Go, PostgreSQL, Swagger, Docker | Events and users REST API with auth. Built after a charter interview, with MCP tooling and a WP-by-WP review. |
| **2** `sdd+/project2` | SDD + best practices | Go API, Vue 3 + Vite + Tailwind, PostgreSQL, Mailpit | Event management web platform: UI on top of an extended Project 1 API, e-mail, responsive design. |
| **3** `sdd+/project3` | SDD + best practices | .NET 10, RabbitMQ, PostgreSQL, Nginx gateway, Vue 3, Docker Compose / Kubernetes | Distributed platform: Identity, Events and Registrations services, Notification and Analytics workers, and a dead-letter service talking over a message broker. Real-time check-in and an admin portal. |

## Results

Metrics were collected with **OpenCode Monitor** (session-level: tool calls, duration, cost) and **Google AI Studio** (model-level: requests, tokens, errors). The coding agent was [OpenCode](https://opencode.ai/), using Google models, except Project 1.1, which used Gemini CLI.

| Project / Approach | Tokens      | Est. cost | Tool calls | Duration |
|--------------------|-------------|-----------|------------|----------|
| 1.1 – Vibe coding  | 6,932,177   | n/a       | 74         | 15 m     |
| 1.2 – SDD          | 14,146,448  | $3.53     | 134        | 58 m     |
| 1.3 – SDD+         | 48,837,903  | $12.93    | 320        | 1 h 52 m |
| 2 – SDD+           | 73,318,747  | $20.37    | 384        | 2 h 03 m |
| 3 – SDD+           | 209,442,915 | $27.71    | 1,037      | 13 h 57 m |

**Takeaways**

- **Vibe coding is fast and cheap but incomplete.** Project 1.1 finished in 15 minutes but did not implement user CRUD or authentication.
- **SDD adds organisation, traceability and control.** Every change has a spec, a plan, work packages and a review trail.
- **That rigour costs more.** Tokens, tool calls and time grow sharply with complexity. Project 3 alone used more than all the other projects combined.
- **Spec-Kitty works best** where clear requirements, documentation and validation matter more than fast prototyping. It needs careful control of scope, specs and iteration cycles.
- **Local models** (Llama 3.1 8B, Qwen 2.5 Coder 14B and others) were tried on a school machine and dropped. They hallucinated or misread the context, so cloud models were used instead.

> These applications were built to evaluate the methodology. They are **not production-ready**, and some services still have build errors.

## Running a project

Each project is self-contained. As a quick reference:

```bash
# Spec-Kitty itself
pip install spec-kitty-cli

# Project 1.1 (Java)
cd vibe/vibe-project1.1 && mvn spring-boot:run          # H2 file DB in ./data, console at /h2-console

# Project 1 vibe (Node)
cd vibe/vibe-project1.2 && cp .env.example .env && npm install && npx prisma db push && npm start

# Project 1.2 (FastAPI)
cd sdd/Project1.2-app/vibe-project1.2 && pip install -r requirements.txt && uvicorn main:app --reload

# Project 1.3 / Project 2 (Go + Postgres)
cd sdd+/project1 && docker compose up --build           # Swagger UI exposed by the API
cd sdd+/project2 && docker compose up --build

# Project 3 (.NET microservices + Vue)
cd sdd+/project3/backend/deploy/local && docker compose up -d --build
cd sdd+/project3/frontend && npm install && npm run dev
```

For details, see each project's `kitty-specs/*/quickstart.md` and `sdd+/project3/quickstart.md`.

## Authors

- **Eduardo Amado**
- **Tiago Ramos**

**Supervisor:** Prof. Ricardo Gomes · **Co-supervisor:** Prof. Vítor Távora

Polytechnic of Leiria, School of Technology and Management (ESTG), Department of Computer Engineering, 2025/26.
