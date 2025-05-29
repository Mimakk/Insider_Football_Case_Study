APP_NAME = league-sim
DB_URL = postgres://postgres:postgres@localhost:5433/league?sslmode=disable

run:
	go run main.go

test:
	go test ./... -v

migrate-up:
	migrate -path migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path migrations -database "$(DB_URL)" down

docker-build:
	docker build -t $(APP_NAME) .

docker-run:
	docker run -p 8080:8080 $(APP_NAME)

db-shell:
	psql -U postgres -h localhost -p 5433 -d league

tidy:
	go mod tidy

seed:
	go run main.go # Will seed if needed
