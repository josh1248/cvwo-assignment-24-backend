package models

import (
	"log"

	"github.com/josh1248/cvwo-assignment-24-backend/internal/entities"
)

func FindAllUsers() (users []entities.User) {
	//SQL -> Golang data conversion

	//Use StructScan here if memory becomes an issue.
	err := db.Select(&users, "SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}

	return users
}
