package main

import (
	"insider_football_case_study/database"
	"insider_football_case_study/server"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  .env file not found, using defaults")
	}
	database.InitDB()
	server.SetupServer()
}
