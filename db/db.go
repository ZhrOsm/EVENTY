package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db") // Hier kein :=, sondern =

	if err != nil {
		panic(err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventsTable := `CREATE TABLE IF NOT EXISTS events (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    location TEXT NOT NULL,
    dateTime DATETIME NOT NULL,
    user_id INTEGER
)
`
	_, err := DB.Exec(createEventsTable)
	if err != nil {
		// Detailliertere Fehlermeldung anzeigen
		fmt.Printf("Error while creating table: %v\n", err) // Ausgabe des tatsächlichen Fehlers
		panic("could not create events table")
	}
}
