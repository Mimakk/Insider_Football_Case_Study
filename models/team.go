package models

type Team struct {
	Name     string
	Strength float64
	Stats    TeamStats
}

type TeamStats struct {
	Played       int
	Wins         int
	Draws        int
	Losses       int
	GoalsFor     int
	GoalsAgainst int
	Points       int
}
