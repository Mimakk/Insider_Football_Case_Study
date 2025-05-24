package simulator

import (
	"fmt"
	"insider_football_case_study/models"
	"sort"
)

func (l *League) PrintTable() {
	type Row struct {
		Name string
		*models.Team
	}
	var rows []Row
	for name, team := range l.Teams {
		rows = append(rows, Row{name, team})
	}
	sort.Slice(rows, func(i, j int) bool {
		a := rows[i].Stats
		b := rows[j].Stats
		if a.Points != b.Points {
			return a.Points > b.Points
		}
		return (a.GoalsFor - a.GoalsAgainst) > (b.GoalsFor - b.GoalsAgainst)
	})
	fmt.Println("\nLEAGUE TABLE")
	fmt.Printf("%-15s %4s %4s %4s %4s\n", "Team", "Pts", "GD", "GF", "GA")
	for _, r := range rows {
		s := r.Team.Stats
		fmt.Printf("%-15s %4d %4d %4d %4d\n", r.Name, s.Points, s.GoalsFor-s.GoalsAgainst, s.GoalsFor, s.GoalsAgainst)
	}
}
