package simulator

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
	return rows
}
