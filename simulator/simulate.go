package simulator

import "insider_football_case_study/models"

func (l *League) simulateMatches(filter func(match *models.Match) bool) {
	for i := range l.Fixtures {
		match := &l.Fixtures[i]
		if filter(match) && !match.Played {
			home := l.Teams[match.HomeTeam]
			away := l.Teams[match.AwayTeam]
			hg, ag := SimulateMatch(*home, *away)
			match.HomeGoals = hg
			match.AwayGoals = ag
			match.Played = true
			l.Results = append(l.Results, *match)
			applyMatchResult(home, away, hg, ag)
		}
	}
}

func (l *League) SimulateAllMatches() {
	l.simulateMatches(func(match *models.Match) bool {
		return true
	})
}

func (l *League) SimulateWeek(week int) {
	l.simulateMatches(func(match *models.Match) bool {
		return match.Week == week
	})
}

func (l *League) PredictRemainingWeeks(fromWeek int) {
	cloned := l.deepCopy()

	cloned.simulateMatches(func(m *models.Match) bool {
		return m.Week > fromWeek && !m.Played
	})
}
