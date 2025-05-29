package server

import (
	"insider_football_case_study/database"
	"insider_football_case_study/models"
	"insider_football_case_study/simulator"
	"log"
)

var league simulator.LeagueAPI

func initLeague() {
	var teams []models.Team
	err := database.DB.Select(&teams, "SELECT name, strength FROM teams")
	if err != nil {
		log.Fatalf("Failed to fetch teams from DB: %v", err)
	}
	realLeague := simulator.NewLeague(teams)
	realLeague.GenerateFixtures()
	league = realLeague
}
