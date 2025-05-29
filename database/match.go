package database

import (
	"insider_football_case_study/models"
)

// Fetch all played matches
func GetPlayedMatches() ([]models.Match, error) {
	var matches []models.Match
	err := DB.Select(&matches, "SELECT week, home_team, away_team, home_goals, away_goals, played FROM matches WHERE played = true ORDER BY week")
	return matches, err
}

// Save match result (insert if new, update if exists)
func SaveMatchResult(m models.Match) error {
	_, err := DB.Exec(`
		INSERT INTO matches (week, home_team, away_team, home_goals, away_goals, played)
		VALUES ($1, $2, $3, $4, $5, true)
		ON CONFLICT (week, home_team, away_team)
		DO UPDATE SET home_goals = $4, away_goals = $5, played = true`,
		m.Week, m.HomeTeam, m.AwayTeam, m.HomeGoals, m.AwayGoals,
	)
	return err
}
