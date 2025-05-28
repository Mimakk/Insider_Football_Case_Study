package simulator

import (
	"insider_football_case_study/models"
	"testing"
)

func TestGetFixturesAndResults(t *testing.T) {
	league := League{
		Fixtures: []models.Match{
			{HomeTeam: "A", AwayTeam: "B", Played: false},
			{HomeTeam: "B", AwayTeam: "A", Played: true},
		},
	}

	fixtures := league.GetFixtures()
	results := league.GetResults()

	if len(fixtures) != 1 {
		t.Errorf("expected 1 fixture, got %d", len(fixtures))
	}
	if len(results) != 1 {
		t.Errorf("expected 1 result, got %d", len(results))
	}
}
