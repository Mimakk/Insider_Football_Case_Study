package models

type Match struct {
	Week      int
	HomeTeam  string
	AwayTeam  string
	HomeGoals int
	AwayGoals int
	Played    bool
}
