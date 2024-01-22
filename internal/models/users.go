package models

import (
	"database/sql"
	"log"

	_ "github.com/glebarez/go-sqlite"
	_ "github.com/josh1248/cvwo-assignment-24-backend/internal/entities"
)

func FindAllUsers(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM users")
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	var (
		id   int
		name string
		rep  int
	)

	for rows.Next() {
		err = rows.Scan(&id, &name, &rep)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name, rep)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
