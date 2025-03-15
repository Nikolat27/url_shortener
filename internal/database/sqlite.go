package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func ExecuteQuery(db *sql.DB, query string) error {
	_, err := db.Prepare(query)
	return err
}

func RunMigrations(db *sql.DB) error {
	migrationDir, _ := filepath.Abs("./internal/database/migrations")
	files, err := ioutil.ReadDir(migrationDir)
	if err != nil {
		return err
	}

	// Sort files numerically (001_, 002_, 003_)
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	for _, file := range files {
		filename := file.Name()

		// Read the SQL file
		sqlBytes, err := os.ReadFile(migrationDir + "/" + filename)
		if err != nil {
			return err
		}

		// Execute the SQL script
		_, err = db.Exec(string(sqlBytes))
		if err != nil {
			return fmt.Errorf("error executing %s: %w", filename, err)
		}

		fmt.Println("Applied migration:", filename)
	}

	return nil
}
