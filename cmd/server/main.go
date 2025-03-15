package main

import (
	"fmt"
	"log"
	"os"
	"url_shortener/internal/database"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Run migrations only when the command is "go run main.go migrate"
	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		err = database.RunMigrations(db)
		if err != nil {
			log.Fatal("Migration failed:", err)
		}
		fmt.Println("Migrations completed successfully!")
		return
	}

	fmt.Println("Server running...")
}
