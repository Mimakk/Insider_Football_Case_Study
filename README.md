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