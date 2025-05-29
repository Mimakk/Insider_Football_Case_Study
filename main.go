package main

import (
	"insider_football_case_study/database"
	"insider_football_case_study/server"

	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  .env not found — using default DB config")
	}

	database.InitDB()
	database.SeedTeamsIfEmpty()

	server.SetupServer()
}
