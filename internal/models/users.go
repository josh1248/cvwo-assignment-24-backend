package models

import (
	"strconv"

	"github.com/josh1248/cvwo-assignment-24-backend/internal/entities"
)

// Use StructScan here if memory becomes an issue.
func FindAllUsers() ([]entities.User, error) {
	var users []entities.User
	err := db.Select(&users, "SELECT * FROM users")
	return users, err
}

// ID is unique, so return type is a single item
func FindUserByID(id string) (entities.User, error) {

	var user entities.User

	id_int, err := strconv.Atoi(id)
	if err != nil {
		return user, err
	}

	err = db.Get(&user, "SELECT * FROM users WHERE id = ?", id_int)
	//sql.ErrNoRows to seperate different errors.
	return user, err
}

// Create a new user with a given name.
// names do not need to be unique.
// returns ID
func CreateUser(name string) error {
	//Using the behaviour of id as an INT PRIMARY KEY row,
	//SQLite can generate incremental IDs for us if we leave it empty.
	//Documentation: https://www.sqlite.org/autoinc.html
	_, err := db.Exec("INSERT INTO users (name, reputation) VALUES (?, ?)", name, 0)
	return err

	// createStmt.LastInsertId() possible here?
}
