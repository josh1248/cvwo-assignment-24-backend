package models

import (
	"log"

	"github.com/josh1248/cvwo-assignment-24-backend/internal/auth"
	"github.com/josh1248/cvwo-assignment-24-backend/internal/entities"
)

// need to break abstraction here since there is no direct way to translate Golang structs into SQL queries.
// TODO: process password with backticks??
var testUsers []entities.User = []entities.User{
	{ID: 1, Name: "Jojo", Reputation: 200, Password: `letmein`},
	{ID: 2, Name: "PasswordGuy", Reputation: -30, Password: `1e!"E#@5yu6V52~\n42`},
	{ID: 3, Name: "Momo", Reputation: 1500, Password: `qwerty123`},
	{ID: 4, Name: "Geegee", Reputation: 0, Password: `letmein`},
}

// Clears junk data that may have been inputted for testing.
// Helps to avoid the need for repeated db meddling in sqlite.
func resetDB() {
	_, err := db.Exec("DROP TABLE users")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE users (
			id INTEGER PRIMARY KEY,
			name TEXT UNIQUE NOT NULL,
			reputation INT NOT NULL,
			password TEXT UNIQUE NOT NULL
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	for _, testUser := range testUsers {
		processedPW, err := auth.ProcessPassword(testUser.Password)
		if err != nil {
			log.Fatal(err)
		}

		_, err = db.Exec(`
			INSERT INTO users (id, name, reputation, password) VALUES (?, ?, ?, ?)`,
			testUser.ID, testUser.Name, testUser.Reputation, processedPW)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Table in database restarted.")
}
