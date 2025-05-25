package server

import (
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
	league.PredictRemainingWeeks(week)
	w.Write([]byte("Prediction complete."))
}
