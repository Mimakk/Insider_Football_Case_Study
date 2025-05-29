package simulator

import "insider_football_case_study/models"

func (l *League) GenerateFixtures() {
	var teams []string
	for name := range l.Teams {
		teams = append(teams, name)
	}

	if len(teams)%2 != 0 {
		teams = append(teams, "BYE")
	}

	numTeams := len(teams)
	numWeeks := numTeams - 1

	for week := 1; week <= numWeeks; week++ {
		for i := 0; i < numTeams/2; i++ {
			home := teams[i]
			away := teams[numTeams-1-i]

			if home != "BYE" && away != "BYE" {
				l.Fixtures = append(l.Fixtures, models.Match{
					Week:     week,
					HomeTeam: home,
					AwayTeam: away,
					Played:   false,
				})
			}
			l.Fixtures = append(l.Fixtures, models.Match{
				Week:     week + (numTeams - 1),
				HomeTeam: away,
				AwayTeam: home,
			})
		}

		// Rotate teams
		teams = append([]string{teams[0]}, append([]string{teams[len(teams)-1]}, teams[1:len(teams)-1]...)...)
	}
}
