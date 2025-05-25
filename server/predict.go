package server

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func handlePredict(w http.ResponseWriter, r *http.Request) {
	weekStr := r.URL.Query().Get("fromWeek")
	week, err := strconv.Atoi(weekStr)
	if err != nil {
		http.Error(w, "Invalid fromWeek", http.StatusBadRequest)
		return
	}
	predictedTable := league.PredictRemainingWeeks(week)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(predictedTable)
}
