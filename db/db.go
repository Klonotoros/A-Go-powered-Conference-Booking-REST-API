package db

import (
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error

	DB, err = sql.Open("sqlite", "api.db")

	if err != nil {
		panic("Database could not connect: " + err.Error())
	}

	err = createTables()
	if err != nil {
		panic("Database could not connect: " + err.Error())
	}

	fmt.Println("Tables created successfully!")
}

func createTables() error {
	var err error
	createUsersTable := `
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            email TEXT NOT NULL UNIQUE,
            password TEXT NOT NULL
        )
    `
	_, err = DB.Exec(createUsersTable)
	if err != nil {
		panic("Could not create users table.")
	}

	createConferencesTable := `
        CREATE TABLE IF NOT EXISTS conferences (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            description TEXT NOT NULL,
            location TEXT NOT NULL,
            dateTime DATETIME NOT NULL,
            user_id INTEGER,
            FOREIGN KEY(user_id) REFERENCES users(id)
        )
    `

	_, err = DB.Exec(createConferencesTable)
	if err != nil {
		panic("Could not create conferences table.")
	}

	createRegistrationsTable := `
		CREATE TABLE IF NOT EXISTS registrations (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			conference_id INTEGER,
			user_id INTEGER,
			FOREIGN KEY(conference_id) REFERENCES conference(id),
			FOREIGN KEY(user_id) REFERENCES users(id)
		)
    `

	_, err = DB.Exec(createRegistrationsTable)

	if err != nil {
		panic("Could not create registrations table.")
	}

	return err
}
