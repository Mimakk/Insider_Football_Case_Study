package server

import (
	"insider_football_case_study/models"
	"insider_football_case_study/simulator"
)

var league simulator.LeagueAPI

func initLeague() {
	teams := []models.Team{
		{Name: "Fenerbahce", Strength: 1.1},
		{Name: "Besiktas", Strength: 1.0},
		{Name: "Galatasaray", Strength: 1.2},
		{Name: "Trabzonspor", Strength: 0.9},
	}
	realLeague := simulator.NewLeague(teams)
	realLeague.GenerateFixtures()
	league = realLeague
}
