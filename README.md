# ⚽ Insider Football League Simulator

A backend football league engine that simulates match results, league standings, and predictions — with full REST API, PostgreSQL persistence, and test coverage.

---

## 🚀 Features

- Simulate all matches or one week at a time
- Predict future results from any week onward
- Edit match results
- View league table, results, fixtures
- Persist teams and matches in PostgreSQL
- Clean code structure (SOLID principles)
- Full unit test coverage
- Docker-ready and CI setup

---

## 📦 Tech Stack

- **Go** 1.20
- **PostgreSQL** (via Docker)
- **REST API** (net/http)
- **SQLX** for DB access
- **migrate** for DB migrations
- **Postman** or `curl` for testing endpoints

---

## 🛠 Setup Instructions

### 1. Clone and setup:

```bash
git clone https://github.com/mimakk/insider-league.git
cd insider-league
cp .env.example .env   # if using .env files

## 🐳 Docker

Build and run:

```bash
docker build -t football-sim .
docker run -p 8080:8080 football-sim


API runs at: http://localhost:8080


---

### ✅ Coverage Command

```md
## 🧪 Run Tests with Coverage

```bash
go test ./simulator -coverprofile=coverage.out
go tool cover -html=coverage.out

## 📬 Postman Collection

You can test all API endpoints using Postman.

📁 [Download insider_football_case.postman_collection.json](./insider_football_case.postman_collection.json)

- Simulate matches
- Edit results
- View table, fixtures, predictions


## 📚 API Documentation (OpenAPI)

This project includes an `openapi.yaml` file for Swagger support.

- View and test it with [Swagger Editor](https://editor.swagger.io)
- Import it into tools like Postman, Insomnia, Stoplight, etc.

## 📘 Swagger UI

To view interactive API docs:

```bash
docker run -p 8080:8080 -p 8081:8081 football-sim



If you don’t have it:

```bash
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest