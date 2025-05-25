package server

import (
	"encoding/json"
	"net/http"
)

func handleEditMatch(w http.ResponseWriter, r *http.Request) {
	type EditRequest struct {
		Week      int    `json:"week"`
		HomeTeam  string `json:"home_team"`
		AwayTeam  string `json:"away_team"`
		HomeGoals int    `json:"home_goals"`
		AwayGoals int    `json:"away_goals"`
	}
	var req EditRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if err := league.EditMatch(req.Week, req.HomeTeam, req.AwayTeam, req.HomeGoals, req.AwayGoals); err != nil {
		http.Error(w, "Match not found", http.StatusNotFound)
		return
	}
	w.Write([]byte("Match updated successfully"))
}
