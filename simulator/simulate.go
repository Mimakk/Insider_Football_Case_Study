package simulator

func (l *League) SimulateAllMatches() {
	for i := range l.Fixtures {
		if !l.Fixtures[i].Played {
			match := &l.Fixtures[i]
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
