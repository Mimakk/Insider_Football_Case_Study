package simulator

import "sort"

func (l *League) GetTable() []TeamRow {
	var rows []TeamRow
	for name, t := range l.Teams {
		s := t.Stats
		rows = append(rows, TeamRow{
			Name:   name,
			Points: s.Points,
			GD:     s.GoalsFor - s.GoalsAgainst,
			GF:     s.GoalsFor,
			GA:     s.GoalsAgainst,
		})
	}

	sort.Slice(rows, func(i, j int) bool {
		if rows[i].Points == rows[j].Points {
			return rows[i].GD > rows[j].GD
		}
		return rows[i].Points > rows[j].Points
	})

	return rows
}
