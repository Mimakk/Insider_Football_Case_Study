package simulator

import (
	"insider_football_case_study/models"
	"testing"
)

func TestSimulateMatchesCustomFilter(t *testing.T) {
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
		return m.Week == 2 // simulate only week 2
	})

	if league.Fixtures[0].Played {
		t.Errorf("week 1 should not be played")
	}
	if !league.Fixtures[1].Played {
		t.Errorf("week 2 should be played")
	}
}

func TestNewLeague(t *testing.T) {
	teams := []models.Team{
		{Name: "A"},
		{Name: "B"},
	}
	l := NewLeague(teams)

	if len(l.Teams) != 2 {
		t.Errorf("expected 2 teams, got %d", len(l.Teams))
	}
	if _, ok := l.Teams["A"]; !ok {
		t.Errorf("team A not found in league")
	}
}
