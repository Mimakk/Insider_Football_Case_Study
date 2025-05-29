package server

import (
	"encoding/json"
	"net/http"
)

func handleMeta(w http.ResponseWriter, r *http.Request) {
	fixtures := league.GetFixtures()
	played := 0
	total := len(fixtures)

	for _, match := range fixtures {
		if match.Played {
			played++
		}
	}

	json.NewEncoder(w).Encode(map[string]int{
		"total_matches":     total,
		"played_matches":    played,
		"remaining_matches": total - played,
		"weeks_total":       league.TotalWeeks(),
	})
}
