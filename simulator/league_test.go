package simulator

import (
	"insider_football_case_study/models"
	"testing"
)

func TestApplyMatchResult(t *testing.T) {
	home := &models.Team{Name: "Fenerbahce"}
	away := &models.Team{Name: "Besiktas"}

	applyMatchResult(home, away, 2, 1)

	if home.Stats.Points != 3 {
		t.Errorf("expected home points to be 3, got %d", home.Stats.Points)
	}
	if away.Stats.Losses != 1 {
		t.Errorf("expected away losses to be 1, got %d", away.Stats.Losses)
	}
	if home.Stats.GoalsFor != 2 || away.Stats.GoalsAgainst != 2 {
		t.Errorf("goal stats not applied correctly")
	}
}

func TestReverseMatchResult(t *testing.T) {
	home := &models.Team{Name: "Fenerbahce", Stats: models.TeamStats{
		Points: 3, Wins: 1, GoalsFor: 2, GoalsAgainst: 1, Played: 1,
	}}
	away := &models.Team{Name: "Besiktas", Stats: models.TeamStats{
		Losses: 1, GoalsFor: 1, GoalsAgainst: 2, Played: 1,
	}}

	reverseMatchResult(home, away, 2, 1)

	if home.Stats.Points != 0 || home.Stats.Wins != 0 || home.Stats.Played != 0 {
		t.Errorf("home stats not reversed correctly")
	}
	if away.Stats.Losses != 0 || away.Stats.GoalsFor != 0 || away.Stats.Played != 0 {
		t.Errorf("away stats not reversed correctly")
	}
}

func TestGetTableSorting(t *testing.T) {
	league := League{
		Teams: map[string]*models.Team{
			"A": {Name: "A", Stats: models.TeamStats{Points: 6, GoalsFor: 5, GoalsAgainst: 3}},
			"B": {Name: "B", Stats: models.TeamStats{Points: 6, GoalsFor: 4, GoalsAgainst: 2}}, // higher GD
			"C": {Name: "C", Stats: models.TeamStats{Points: 3}},
		},
	}

	table := league.GetTable()

	if table[0].Name != "B" || table[1].Name != "A" || table[2].Name != "C" {
		if table[0].Points != 6 || table[0].GD != 2 {
			t.Errorf("Expected top team to have 6 points and 2 GD, got %+v", table[0])
		}
	}
}

func TestEditMatch(t *testing.T) {
	league := League{
		Teams: map[string]*models.Team{
			"Fenerbahce": {Name: "Fenerbahce"},
			"Besiktas":   {Name: "Besiktas"},
		},
		Fixtures: []models.Match{
			{Week: 1, HomeTeam: "Fenerbahce", AwayTeam: "Besiktas", Played: false},
		},
	}

	// Edit match to set result
	err := league.EditMatch(1, "Fenerbahce", "Besiktas", 3, 1)
	if err != nil {
		t.Fatalf("edit failed: %v", err)
	}

	fb := league.Teams["Fenerbahce"]
	bjk := league.Teams["Besiktas"]

	if fb.Stats.Points != 3 {
		t.Errorf("expected 3 points, got %d", fb.Stats.Points)
	}
	if bjk.Stats.Losses != 1 {
		t.Errorf("expected 1 loss for Besiktas, got %d", bjk.Stats.Losses)
	}
}

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
