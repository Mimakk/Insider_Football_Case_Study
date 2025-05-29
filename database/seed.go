package database

import (
	"fmt"
	"log"

	"insider_football_case_study/models"
)

func SeedTeamsIfEmpty() {
	var count int
	err := DB.Get(&count, "SELECT COUNT(*) FROM teams")
	if err != nil {
		log.Fatalf("Failed to query teams count: %v", err)
	}

	if count > 0 {
		fmt.Println("✅ Teams already exist in DB")
		return
	}

	tx := DB.MustBegin()
	for _, team := range models.InitialTeams {
		_, err := tx.NamedExec(`INSERT INTO teams (name, strength) VALUES (:name, :strength)`, team)
		if err != nil {
			log.Fatalf("Failed to insert team: %v", err)
		}
	}
	if err := tx.Commit(); err != nil {
		log.Fatalf("Failed to commit seed: %v", err)
	}
	fmt.Println("✅ Seeded initial teams into DB")
}
