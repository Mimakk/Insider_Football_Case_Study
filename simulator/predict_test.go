package simulator

import (
	"insider_football_case_study/models"
	"testing"
)

func TestPredictRemainingWeeks(t *testing.T) {
	league := League{
		Teams: map[string]*models.Team{
			"A": {Name: "A", Strength: 1.1},
			"B": {Name: "B", Strength: 1.0},
		},
		Fixtures: []models.Match{
			{Week: 1, HomeTeam: "A", AwayTeam: "B", Played: true},
			{Week: 2, HomeTeam: "B", AwayTeam: "A"},
		},
	}

	prediction := league.PredictRemainingWeeks(1)

	if len(prediction) != 2 {
		t.Fatalf("expected 2 teams, got %d", len(prediction))
	}

	// Validate no change to original fixtures
	if league.Fixtures[1].Played {
		t.Error("real league should not be mutated by prediction")
	}
}
