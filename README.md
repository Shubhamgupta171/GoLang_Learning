
# GoAuthX

> **GoAuthX** — Production-ready Authentication & User Management Backend (SaaS)  
> Build secure, scalable APIs with Go, Gin, SQLite/Postgres, Redis, JWT, Docker, and CI/CD.  
> This repository contains a complete learning roadmap, daily feature plan, and a production-grade feature set you can iterate on.

---

## Table of Contents
- [Project Overview](#project-overview)
- [Why GoAuthX?](#why-goauthx)
- [Tech Stack](#tech-stack)
- [Key Features (Out-of-the-box)](#key-features-out-of-the-box)
- [Architecture (textual)](#architecture-textual)
- [Quick Start (Local)](#quick-start-local)
- [Environment Variables](#environment-variables)
- [Run with Docker (Recommended)](#run-with-docker-recommended)
- [API Endpoints (Core)](#api-endpoints-core)
- [Database Migrations & Models](#database-migrations--models)
- [Daily Improvement Roadmap (30 days)](#daily-improvement-roadmap-30-days)
- [Advanced Features to Add (SaaS-ready)](#advanced-features-to-add-saas-ready)
- [Testing & CI/CD](#testing--cicd)
- [Deployment](#deployment)
- [Security Checklist](#security-checklist)
- [Resume & Interview Tips (How to present GoAuthX)](#resume--interview-tips-how-to-present-goauthx)
- [Contributing](#contributing)
- [License](#license)

---

## Project Overview

**GoAuthX** is a single-repo, production-oriented backend template and learning project that evolves into a SaaS Authentication & User Management platform.  
Start small (SQLite + Gin + Redis + JWT) and progressively add features to make it enterprise-ready: multi-tenant support, API keys, analytics, webhooks, and more.

This README serves two purposes:
1. A developer learning guide (daily tasks + examples).
2. A project growth playbook to convert your repo into a robust SaaS backend.

---

## Why GoAuthX?

- Real-world: Implements authentication best practices, token rotation, Redis session store, role-based access control.
- Extensible: Designed to be extended into a full API platform or integrated into other microservices.
- Resume-friendly: Demonstrates security, scalability, CI/CD, and deployment skills.

---

## Tech Stack

- Language: **Go (>=1.20)**
- Web framework: **Gin**
- ORM: **GORM** (SQLite for dev, Postgres for production)
- Cache / Session store: **Redis**
- Authentication: **JWT** (access tokens) + Refresh tokens
- Containerization: **Docker / Docker Compose**
- Docs: **Swagger / OpenAPI**
- CI/CD: **GitHub Actions**
- Optional: **Nginx** (reverse proxy), **Prometheus/Grafana** (metrics), **Sentry** (error monitoring)

---

## Key Features (Out-of-the-box)

- Register / Login (hashed passwords using bcrypt)
- JWT-based access tokens + Refresh token rotation stored in Redis
- Token blacklist on logout
- Role-Based Access Control (admin, user, manager)
- Email OTP verification flow (Redis TTL based)
- Rate-limiting for sensitive endpoints (login) using Redis
- Basic user CRUD and profile endpoints
- File upload support (profile picture) — metadata in DB
- Basic request logging middleware
- Health check endpoint for readiness & liveness

---

## Architecture (textual)

Client (Web/Mobile)  
→ Reverse Proxy (optional: Nginx)  
→ GoAuthX API (Gin)  
  - Auth Service (JWT, OTP, Refresh tokens)  
  - User Service (CRUD, profile)  
  - Admin endpoints (user management)  
  - Background worker (email, cleanup)  
  - Redis (token store, rate-limit, cache)  
  - DB (SQLite dev / Postgres production)  

---

## Quick Start (Local)

> Development setup uses SQLite and local Redis for simplicity.

### Prerequisites
- Go (1.20+) installed
- Redis running locally (`redis-server`)
- Git
- Docker (optional but recommended)

### Clone
```bash
git clone https://github.com/your-username/goauthx.git
cd goauthx
```

### Local env
Create `.env` in project root (see [Environment Variables](#environment-variables)).

### Run
```bash
go run ./cmd/server
```

Open: `http://localhost:8080/health`

---

## Environment Variables

Create `.env` with (example):
```env
PORT=8080
GIN_MODE=debug

# Database
DB_DRIVER=sqlite           # or "postgres"
DB_DSN=goauthx.db          # or "host=... user=... password=... dbname=... sslmode=disable"

# Redis
REDIS_ADDR=localhost:6379
REDIS_PASS=
REDIS_DB=0

# JWT
JWT_SECRET=replace_with_strong_secret
JWT_EXP=15m
REFRESH_TOKEN_EXP=168h

# Email (SMTP) - optional
SMTP_HOST=smtp.example.com
SMTP_PORT=587
SMTP_USER=you@example.com
SMTP_PASS=yourpassword

# Other
APP_ENV=development
```

---

## Run with Docker (Recommended)

`Dockerfile` (example) and `docker-compose.yml` are provided in repo.

Start locally with Docker Compose:
```bash
docker compose up --build
```

Services launched:
- goauthx (app)
- redis
- db (optional Postgres image)
- maildev (optional SMTP dev server)

---

## API Endpoints (Core)

**Auth**
- `POST /auth/register` — Register new user  
  Body: `{ "username","email","password" }`
- `POST /auth/login` — Login (returns access + refresh tokens)  
- `POST /auth/refresh` — Rotate refresh token, issue new access token  
- `POST /auth/logout` — Invalidate refresh token + blacklist access token  
- `POST /auth/send-otp` — Send email OTP  
- `POST /auth/verify-otp` — Verify OTP

**Users**
- `GET /profile` — Get current user (protected)
- `PUT /profile` — Update profile (protected)
- `GET /admin/users` — List users (admin)
- `GET /admin/users/:id`
- `PUT /admin/users/:id`
- `DELETE /admin/users/:id`

**Utilities**
- `GET /health` — return `{status: "ok"}`
- `GET /swagger/*` — API docs (when enabled)

---

## Database Migrations & Models

Start with simple models:
```go
type User struct {
    ID        uint   `gorm:"primaryKey"`
    Username  string `gorm:"uniqueIndex;not null"`
    Email     string `gorm:"uniqueIndex;not null"`
    Password  string `gorm:"not null"`
    Role      string `gorm:"default:user"`
    AvatarURL string
    CreatedAt time.Time
    UpdatedAt time.Time
}
```

Auto-migrate at startup:
```go
db.AutoMigrate(&User{}, &OtherModels...)
```

Migrate to Postgres before production. Keep migrations deterministic (use goose / golang-migrate if needed).

---

## Daily Improvement Roadmap (30 Days)

This is a **practical daily checklist** you can follow — complete one task per day and push progress to GitHub.

### Week 1 — Core Auth & Project Setup
- Day 1: Project skeleton, health route, env config
- Day 2: User model, DB connection, auto-migrate
- Day 3: Register endpoint (password hashing)
- Day 4: Login endpoint (JWT issue)
- Day 5: Protected `/profile` route with auth middleware
- Day 6: Refresh token endpoint + Redis store
- Day 7: Logout (blacklist token), basic tests

### Week 2 — Security & Quality
- Day 8: Rate limiting (login attempts) with Redis
- Day 9: Email OTP verification flow
- Day 10: Role-based access middleware (admin)
- Day 11: Pagination & search for user list
- Day 12: File upload (profile image)
- Day 13: Input validation & structured error responses
- Day 14: Add unit tests for auth service

### Week 3 — Observability & Reliability
- Day 15: Logging with Zap / Zerolog
- Day 16: Metrics endpoint (Prometheus instrumentation)
- Day 17: Background worker (goroutine + queue) for emails
- Day 18: Webhooks: allow apps to register a webhook URL
- Day 19: API keys system for developers (scoped permissions)
- Day 20: Improve refresh token rotation & secure storage
- Day 21: Integration tests & test coverage

### Week 4 — Production & SaaS Features
- Day 22: Dockerize & docker-compose
- Day 23: Swagger docs & API versioning
- Day 24: CI pipeline (GitHub Actions)
- Day 25: Deploy to platform (Render / Railway / DigitalOcean)
- Day 26: Multi-tenant / Organization support
- Day 27: Websocket or realtime notifications (optional)
- Day 28: Billing hooks & Stripe integration (optional)
- Day 29: Security audit & common vulnerability fixes
- Day 30: Polish README, demo, and portfolio entry

---

## Advanced Features to Add (Make it SaaS-ready)

- Multi-tenant organizations with per-org billing
- API keys + scopes + rate limits per key
- OAuth2 provider support (social login)
- Admin dashboard (React/NextJS) - account management
- Audit logs (immutable event store)
- Role/permission management UI
- Usage analytics and billing reports
- License keys and plan enforcement
- CSRF protection for browser flows
- Sentry integration for error tracking
- Secrets management for keys (Vault)

---

## Testing & CI/CD

- Unit tests for services (Golang `testing`)
- Integration tests using testcontainers (Postgres/Redis)
- Linting: `golangci-lint`
- GitHub Actions:
  - Run tests
  - Run linters
  - Build Docker image
  - Push to registry (optional)
  - Deploy to staging on merge

---

## Deployment

**Staging**: Use Docker Compose or cloud service (Railway, Render, Fly.io)  
**Production**: Use managed Postgres, managed Redis. Use environment secrets rather than `.env`.

Suggested deployment steps:
1. Build Docker image (`docker build -t goauthx:latest .`)
2. Push to registry
3. Update service on host or platform
4. Use migrations on startup or run separately

---

## Security Checklist (must-follow before production)

- Use HTTPS; terminate TLS at proxy or load balancer
- Rotate JWT secret and use env secrets
- Set proper CORS policy (whitelist origins)
- Restrict trusted proxies instead of trusting all
- Use secure cookie flags if cookies used (HttpOnly, Secure, SameSite)
- Rate-limit sensitive endpoints
- Secret scanning and vault storage
- Regular dependency audits

---

## Resume & Interview Tips — How to present GoAuthX

- **Project title**: GoAuthX — Authentication & API Platform
- **One-liner**: Built a production-ready authentication platform with JWT, refresh token rotation, Redis-backed sessions, role-based access control and rate-limiting.
- **Key highlights to mention**:
  - Implemented secure JWT auth with refresh token rotation and blacklist on logout.
  - Used Redis for session and rate-limiting; designed token invalidation strategy.
  - Dockerized the application; set up GitHub Actions for CI.
  - Added automated tests and API documentation (Swagger).
  - Deployed to [your platform], configured managed Postgres and Redis in production.
- **Interview talking points**:
  - Explain the refresh-token rotation flow and why it's more secure.
  - Discuss trade-offs between storing session in Redis vs stateless JWT.
  - Explain how rate-limiting with Redis works (counters, TTLs).
  - Talk about scaling decisions: DB pooling, read replicas, caching.

---

## Contributing

Contributions are welcome! Follow these steps:
1. Fork the repo
2. Create a feature branch (`git checkout -b feat/your-feature`)
3. Add tests
4. Open a PR with a clear description

Please follow code style, tests, and commit message conventions.

---

## License

This project is released under the MIT License. See `LICENSE` for details.

---

## Final Notes

GoAuthX is intentionally modular — start small and iterate daily. Use the 30-day roadmap to keep steady progress and push frequent commits (daily). Each completed day's task becomes an incremental improvement you can show on your resume and GitHub.

Happy building! 🚀
