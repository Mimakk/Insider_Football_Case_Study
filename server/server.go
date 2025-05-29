package server

import (
	"fmt"
	"net/http"
)

func SetupServer() {
	initLeague()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/table", http.StatusSeeOther)
	})
	http.HandleFunc("/simulate/all", handleSimulateAll)
	http.HandleFunc("/simulate/week", handleSimulateWeek)
	http.HandleFunc("/match/edit", handleEditMatch)
	http.HandleFunc("/table", handleTable)
	http.HandleFunc("/predict", handlePredict)
	http.HandleFunc("/fixtures", handleFixtures)
	http.HandleFunc("/results", handleResults)
	http.HandleFunc("/meta", handleMeta)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
