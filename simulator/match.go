package simulator

import (
	"insider_football_case_study/models"
	"math/rand"
)

func SimulateMatch(home, away models.Team) (int, int) {
	base := 1.5
	homeGoals := int(home.Strength * base * rand.Float64())
	awayGoals := int(away.Strength * base * rand.Float64())
	return homeGoals, awayGoals
}
