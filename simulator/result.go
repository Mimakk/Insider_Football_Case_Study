package simulator

import "insider_football_case_study/models"

func applyMatchResult(home, away *models.Team, hg, ag int) {
	home.Stats.Played++
	away.Stats.Played++
	home.Stats.GoalsFor += hg
	home.Stats.GoalsAgainst += ag
	away.Stats.GoalsFor += ag
	away.Stats.GoalsAgainst += hg

	switch {
	case hg > ag:
		home.Stats.Wins++
		home.Stats.Points += 3
		away.Stats.Losses++
	case hg < ag:
		away.Stats.Wins++
		away.Stats.Points += 3
		home.Stats.Losses++
	default:
		home.Stats.Draws++
		away.Stats.Draws++
		home.Stats.Points++
		away.Stats.Points++
	}
}

// will be used in case of edit, in order to reverse results

func reverseMatchResult(home, away *models.Team, hg, ag int) {
	home.Stats.Played--
	away.Stats.Played--
	home.Stats.GoalsFor -= hg
	home.Stats.GoalsAgainst -= ag
	away.Stats.GoalsFor -= ag
	away.Stats.GoalsAgainst -= hg

	switch {
	case hg > ag:
		home.Stats.Wins--
		home.Stats.Points -= 3
		away.Stats.Losses--
	case ag > hg:
		away.Stats.Wins--
		away.Stats.Points -= 3
		home.Stats.Losses--
	default:
		home.Stats.Draws--
		away.Stats.Draws--
		home.Stats.Points--
		away.Stats.Points--
	}
}
