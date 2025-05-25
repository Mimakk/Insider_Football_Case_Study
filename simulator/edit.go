package simulator

import "errors"

// This function looks through all the matches and find the given (home away and week should be given)
// and then reverse its effect them apply new results

func (l *League) EditMatch(week int, home, away string, hg, ag int) error {
	for i := range l.Fixtures {
		m := &l.Fixtures[i]
		if m.Week == week && m.HomeTeam == home && m.AwayTeam == away {
			if m.Played {
				reverseMatchResult(l.Teams[home], l.Teams[away], m.HomeGoals, m.AwayGoals)
			}
			m.HomeGoals = hg
			m.AwayGoals = ag
			m.Played = true
			applyMatchResult(l.Teams[home], l.Teams[away], hg, ag)
			return nil
		}
	}
	return errors.New("match not found")
}
