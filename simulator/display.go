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
	fmt.Println("Team\tPts\tGD\tGF\tGA")
	for _, r := range rows {
		s := r.Stats
		fmt.Printf("%s\t%d\t%d\t%d\t%d\n", r.Name, s.Points, s.GoalsFor-s.GoalsAgainst, s.GoalsFor, s.GoalsAgainst)
	}
}
