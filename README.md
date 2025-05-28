![Go](DONT FORGET TO PUT YOUR REPOLINK)

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
