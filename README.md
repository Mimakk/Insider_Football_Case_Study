
# ⚽ Insider Football League Simulator

A backend simulation engine that models a football league with match predictions, real-time simulations, result editing, and persistent storage via PostgreSQL.

Built with Go, following SOLID principles, fully tested, and ready to deploy.

---

## 🚀 Features

- ✅ Simulate full season or one week at a time
- 🔄 Predict league table from any future week
- ✏️ Edit match results (with stats recalculated)
- 📊 View fixtures, results, league table, and league state
- 💾 Persistence with PostgreSQL
- 🔁 Round-robin fixture generation (fair scheduling)
- ⚙️ RESTful API with clean endpoints
- ✅ Unit tested with Go’s testing package
- 🐳 Docker support and Makefile tasks

---

## 📦 Tech Stack

- Language: **Go (Golang)**
- DB: **PostgreSQL**
- HTTP: `net/http`
- DB driver: `sqlx`, `pq`
- Migrations: `golang-migrate`
- Testing: Go's `testing` + coverage
- Dev Tools: **Makefile**, **Docker**

---

## 🧠 How It Works

The app creates a virtual league of 4 teams, generates round-robin fixtures, and allows full season simulation or week-by-week updates. Teams earn points, goals are calculated based on relative strength, and predictions can be made without committing results.

---

## 🛠 Setup

### 1. Clone and prepare

```bash
git clone https://github.com/mimakk/football-league-sim.git
cd football-league-sim
```

### 2. Run Postgres via Docker

```bash
docker compose up -d
```

> Make sure port `5433` is available

---

### 3. Run DB Migrations

```bash
make migrate-up
```

---

### 4. Seed teams and start server

```bash
make run
```

Server: [http://localhost:8080](http://localhost:8080)

---

## 📬 API Endpoints

| Method | Endpoint                   | Description                           |
|--------|----------------------------|---------------------------------------|
| GET    | `/`                        | Shows Table                           |
| GET    | `/simulate/all`            | Simulate the entire season            |
| GET    | `/simulate/week?week=2`    | Simulate a single week                |
| GET    | `/predict?fromWeek=3`      | Predict future matches & table        |
| GET    | `/table`                   | Get current league standings          |
| GET    | `/fixtures`                | See upcoming fixtures                 |
| GET    | `/results`                 | View played match results             |
| GET    | `/meta`                    | View league state (matches, weeks)    |
| POST   | `/match/edit`              | Edit a match result                   |

---

## ✏️ Example: Edit Match

```http
POST /match/edit
Content-Type: application/json

{
  "week": 3,
  "home_team": "Fenerbahce",
  "away_team": "Besiktas",
  "home_goals": 1,
  "away_goals": 2
}
```

---

## 🧪 Testing

```bash
make test
```

Includes unit tests for simulation, match results, editing, predictions, table sorting, etc.

---

## 🐳 Docker Build & Run

```bash
make docker-build
make docker-run
```

---

## 🔄 Makefile Tasks

| Task            | Purpose                            |
|-----------------|------------------------------------|
| `make run`      | Start the app                      |
| `make test`     | Run unit tests                     |
| `make migrate-up` | Apply database schema             |
| `make tidy`     | Clean go.mod and go.sum            |
| `make docker-build` | Build Docker image             |

---

## 📂 Project Structure

```
.
├── main.go
├── models/          # Team and Match structs
├── simulator/       # League logic, simulate, table, edit
├── server/          # REST handlers
├── database/        # DB init, seed, save/load matches
├── migrations/      # SQL schema files
├── postman/         # Postman collection (for testing)
├── Makefile
└── README.md
```

---

## 🧪 Postman Collection

Use the provided Postman collection under `postman/` folder to try all endpoints locally.

---

## 🧠 Design Highlights

- 🧩 **SOLID**: Separation of concerns between simulator, DB, and API
- ⚙️ Extensible LeagueAPI interface for mocking/testing
- ♻️ In-memory simulation with rollback for predictions
- 🧪 High test coverage and testable logic

---

## 📄 License

MIT
