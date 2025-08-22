# 📝 Simple Go Notes API

A minimal yet powerful RESTful API built with **Go**, **PostgreSQL**, and **JWT Authentication**. Users can register, log in, and manage their personal notes. Built with production best practices in mind — perfect for cloud-native, containerized deployments.

---

## ✨ Features

- ✅ User Registration & Login (JWT-based Auth)
- ✅ Password hashing using bcrypt
- ✅ CRUD for Notes (Create, Read, Update, Delete)
- ✅ PostgreSQL for persistent storage
- ✅ Clean project structure (Go modules)
- ✅ Docker-ready (soon)

---

## 🔧 Tech Stack

| Layer      | Tool                                      |
| ---------- | ----------------------------------------- |
| Language   | [Go](https://golang.org)                  |
| Database   | [PostgreSQL](https://www.postgresql.org/) |
| Auth       | JWT (via `golang-jwt/jwt`)                |
| ORM Layer  | SQLC (type-safe queries)                  |
| Passwords  | Bcrypt hashing                            |
| Deployment | Docker (optional)                         |

---

## 🚀 Getting Started

### Prerequisites

- Go 1.21+
- PostgreSQL 17+
- `sqlc` (optional, if you regenerate DB code)
- Docker (optional)

---

### 🔌 Setup

1. **Clone the repo**

```bash
git clone https://github.com/mikegav23/simple-go-todo.git
cd simple-go-todo
```

2. **Setup environment variables**

- Rename .env.example to simply .env
- Set the postgres url and the jwt secret

3. **Run the app (server)**

```bash
go run cmd/app/main.go
```
