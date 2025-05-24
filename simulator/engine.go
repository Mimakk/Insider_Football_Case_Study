package simulator

import "insider_football_case_study/models"

type League struct {
	Teams    map[string]*models.Team
	Fictures []models.Match
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
