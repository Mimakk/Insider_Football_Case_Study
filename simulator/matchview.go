package simulator

import "insider_football_case_study/models"

func (l *League) GetFixtures() []models.Match {
	var fixtures []models.Match
	for _, m := range l.Fixtures {
		if !m.Played {
			fixtures = append(fixtures, m)
		}
	}
	return fixtures
}

func (l *League) GetResults() []models.Match {
	var results []models.Match
	for _, m := range l.Fixtures {
		if m.Played {
			results = append(results, m)
		}
	}
	return results
}
