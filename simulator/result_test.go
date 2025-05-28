package simulator

import (
	"insider_football_case_study/models"
	"testing"
)

func TestReverseMatchResult_Draw(t *testing.T) {
	home := &models.Team{Name: "A", Stats: models.TeamStats{Draws: 1, Points: 1, Played: 1}}
	away := &models.Team{Name: "B", Stats: models.TeamStats{Draws: 1, Points: 1, Played: 1}}

	reverseMatchResult(home, away, 1, 1)

	if home.Stats.Draws != 0 || away.Stats.Draws != 0 {
		t.Errorf("draw reversal failed")
	}
	if home.Stats.Points != 0 || away.Stats.Points != 0 {
		t.Errorf("points reversal failed")
	}
}

func TestReverseMatchResult_AwayWin(t *testing.T) {
	home := &models.Team{Name: "A", Stats: models.TeamStats{Losses: 1, Played: 1}}
	away := &models.Team{Name: "B", Stats: models.TeamStats{Wins: 1, Points: 3, Played: 1}}

	reverseMatchResult(home, away, 1, 3)

	if home.Stats.Losses != 0 || away.Stats.Wins != 0 {
		t.Errorf("win/loss reversal failed")
	}
	if away.Stats.Points != 0 {
		t.Errorf("away points should be 0, got %d", away.Stats.Points)
	}
}
