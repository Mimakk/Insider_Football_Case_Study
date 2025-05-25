package server

import (
	"fmt"
	"net/http"
)

func SetupServer() {
	initLeague()
	// http.HandleFunc("/simulate/all", handleSimulateAll)
	// http.HandleFunc("/simulate/week", handleSimulateWeek)
	// http.HandleFunc("/match/edit", handleEditMatch)
	// http.HandleFunc("/table", handleTable)
	// http.HandleFunc("/predict", handlePredict)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
