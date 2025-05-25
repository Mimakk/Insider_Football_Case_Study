package server

import (
	"net/http"
	"strconv"
)

func handleSimulateAll(w http.ResponseWriter, r *http.Request) {
	league.SimulateAllMatches()
	w.Write([]byte("All matches simulated."))
}

func handleSimulateWeek(w http.ResponseWriter, r *http.Request) {
	weekStr := r.URL.Query().Get("week")
	week, err := strconv.Atoi(weekStr)
	if err != nil {
		http.Error(w, "Invalid week", http.StatusBadRequest)
		return
	}
	league.SimulateWeek(week)
	w.Write([]byte("Week simulated."))
}
