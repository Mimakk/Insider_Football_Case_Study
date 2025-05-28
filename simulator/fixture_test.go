package simulator

import (
	"insider_football_case_study/models"
	"testing"
)

func TestGenerateFixtures(t *testing.T) {
	teams := []models.Team{
		{Name: "A"},
		{Name: "B"},
	}
	l := NewLeague(teams)
	l.GenerateFixtures()

	if len(l.Fixtures) != 2 {
		t.Errorf("expected 2 fixtures, got %d", len(l.Fixtures))
	}
	if l.Fixtures[0].HomeTeam == l.Fixtures[0].AwayTeam {
		t.Errorf("home and away teams should not be the same")
	}
}
