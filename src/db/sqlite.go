package db

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite3 driver
)

var (
	db   *sql.DB
	once sync.Once
)

// GetSQLiteClient initializes and returns a thread-safe SQLite client.
func GetSQLiteClient(dbPath string) (*sql.DB, error) {
	var err error

	// Ensure the DB is initialized only once
	once.Do(func() {
		// Open the SQLite database
		db, err = sql.Open("sqlite3", fmt.Sprintf("%s?_journal_mode=WAL&_sync=FULL", dbPath))
		if err != nil {
			log.Printf("Error initializing SQLite client: %v", err)
			return
		}

		// Test the connection
		if err = db.Ping(); err != nil {
			log.Printf("Error pinging SQLite database: %v", err)
			_ = db.Close()
			db = nil
		}
	})

	// If initialization failed, return the error
	if db == nil {
		return nil, fmt.Errorf("failed to initialize SQLite client: %v", err)
	}

	return db, nil
}