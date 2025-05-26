package simulator

import (
	"insider_football_case_study/models"
	"testing"
)

func TestSimulateMatchesFilter(t *testing.T) {
	league := League{
		Teams: map[string]*models.Team{
			"A": {Name: "A", Strength: 1.1},
			"B": {Name: "B", Strength: 1.0},
			"C": {Name: "C", Strength: 1.0},
		},
		Fixtures: []models.Match{
			{Week: 1, HomeTeam: "A", AwayTeam: "B"},
			{Week: 2, HomeTeam: "A", AwayTeam: "C"},
		},
	}

	league.simulateMatches(func(m *models.Match) bool {
		return m.Week == 2
	})

	if league.Fixtures[0].Played {
		t.Errorf("Week 1 should not be played")
	}
	if !league.Fixtures[1].Played {
		t.Errorf("Week 2 should be played")
	}
}
