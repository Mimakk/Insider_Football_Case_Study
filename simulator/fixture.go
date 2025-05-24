package simulator

import "insider_football_case_study/models"

func (l *League) GenerateFixtures() {
	week := 1
	for home := range l.Teams {
		for away := range l.Teams {
			if home != away {
				l.Fixtures = append(l.Fixtures, models.Match{
					Week:     week,
					HomeTeam: home,
					AwayTeam: away,
					Played:   false,
				})
				week++
			}
		}
	}
}
