package simulator

import "insider_football_case_study/models"

type League struct {
	Teams    map[string]*models.Team
	Fixtures []models.Match
	Results  []models.Match
}

func NewLeague(teamList []models.Team) *League {
	teams := make(map[string]*models.Team)
	for i := range teamList {
		t := teamList[i]
		teams[t.Name] = &t
	}
	return &League{Teams: teams}
}

func (l *League) TotalWeeks() int {
	max := 0
	for _, m := range l.Fixtures {
		if m.Week > max {
			max = m.Week
		}
	}
	return max
}

// Interface to abstract League for server (Dependency Inversion)
type LeagueAPI interface {
	SimulateAllMatches()
	SimulateWeek(week int)
	GetTable() []TeamRow
	PredictRemainingWeeks(fromWeek int) []TeamRow
	EditMatch(week int, home, away string, hg, ag int) error
	GetFixtures() []models.Match
	GetResults() []models.Match
	TotalWeeks() int
}
