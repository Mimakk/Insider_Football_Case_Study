package simulator

import "insider_football_case_study/models"

func (l *League) deepCopy() *League {
	// Clone teams
	clonedTeams := make(map[string]*models.Team)
	for name, team := range l.Teams {
		copy := *team
		clonedTeams[name] = &copy
	}

	// Clone fixtures
	clonedFixtures := make([]models.Match, len(l.Fixtures))
	copy(clonedFixtures, l.Fixtures)

	// Clone results
	clonedResults := make([]models.Match, len(l.Results))
	copy(clonedResults, l.Results)

	return &League{
		Teams:    clonedTeams,
		Fixtures: clonedFixtures,
		Results:  clonedResults,
	}
}
