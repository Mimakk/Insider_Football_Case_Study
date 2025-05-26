package simulator

import (
	"insider_football_case_study/models"
	"testing"
)

func TestGetTableSorting(t *testing.T) {
	league := League{
		Teams: map[string]*models.Team{
			"A": {Name: "A", Stats: models.TeamStats{Points: 6, GoalsFor: 5, GoalsAgainst: 3}},
			"B": {Name: "B", Stats: models.TeamStats{Points: 6, GoalsFor: 4, GoalsAgainst: 2}},
			"C": {Name: "C", Stats: models.TeamStats{Points: 3}},
		},
	}

	table := league.GetTable()

	if table[0].Name != "A" {
		t.Errorf("expected A first, got %s", table[0].Name)
	}
}
