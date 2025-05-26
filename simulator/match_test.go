package simulator

import (
	"insider_football_case_study/models"
	"testing"
)

func TestApplyMatchResult_HomeWin(t *testing.T) {
	home := &models.Team{Name: "Fenerbahce"}
	away := &models.Team{Name: "Besiktas"}

	applyMatchResult(home, away, 3, 1)

	if home.Stats.Points != 3 || home.Stats.Wins != 1 {
		t.Errorf("home should have 3 points and 1 win: %+v", home.Stats)
	}
	if away.Stats.Losses != 1 || away.Stats.Points != 0 {
		t.Errorf("away should have 1 loss and 0 points: %+v", away.Stats)
	}
}

func TestApplyMatchResult_AwayWin(t *testing.T) {
	home := &models.Team{Name: "Galatasaray"}
	away := &models.Team{Name: "Trabzonspor"}

	applyMatchResult(home, away, 1, 2)

	if away.Stats.Points != 3 || away.Stats.Wins != 1 {
		t.Errorf("away should have 3 points and 1 win: %+v", away.Stats)
	}
	if home.Stats.Losses != 1 || home.Stats.Points != 0 {
		t.Errorf("home should have 1 loss and 0 points: %+v", home.Stats)
	}
}

func TestApplyMatchResult_Draw(t *testing.T) {
	home := &models.Team{Name: "Galatasaray"}
	away := &models.Team{Name: "Fenerbahce"}

	applyMatchResult(home, away, 2, 2)

	if home.Stats.Points != 1 || home.Stats.Draws != 1 {
		t.Errorf("home should have 1 point and 1 draw: %+v", home.Stats)
	}
	if away.Stats.Points != 1 || away.Stats.Draws != 1 {
		t.Errorf("away should have 1 point and 1 draw: %+v", away.Stats)
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
