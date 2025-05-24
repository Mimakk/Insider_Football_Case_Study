package main

import (
	"insider_football_case_study/models"
	"insider_football_case_study/simulator"
)

func main() {
	teams := []models.Team{
		{Name: "Fenerbahce", Strength: 1.1},
		{Name: "Besiktas", Strength: 1.0},
		{Name: "Galatasaray", Strength: 1.2},
		{Name: "Trabzonspor", Strength: 0.9},
	}

	league := simulator.NewLeague(teams)
	league.GenerateFixtures()
	league.SimulateAllMatches()
	league.PrintTable()
}
