package simulator

import (
	"insider_football_case_study/models"
	"testing"
)

func TestEditMatch(t *testing.T) {
	league := League{
		Teams: map[string]*models.Team{
			"Fenerbahce": {Name: "Fenerbahce"},
			"Besiktas":   {Name: "Besiktas"},
		},
		Fixtures: []models.Match{
			{Week: 1, HomeTeam: "Fenerbahce", AwayTeam: "Besiktas"},
		},
	}

	err := league.EditMatch(1, "Fenerbahce", "Besiktas", 3, 1)
	if err != nil {
		t.Fatalf("edit failed: %v", err)
	}

	if league.Teams["Fenerbahce"].Stats.Points != 3 {
		t.Errorf("expected 3 points, got %d", league.Teams["Fenerbahce"].Stats.Points)
	}
}
