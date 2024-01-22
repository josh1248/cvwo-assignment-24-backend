package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// 127.0.0.1 refers to localhost
var secretCredentials string = `
	user=postgres
	password=xddd
	host=127.0.0.1
	port=5432
	dbname=postgres
	sslmode=disable
`

func ConnectToDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", secretCredentials)
	//defer db.Close()
	if err != nil {
		log.Fatal(err)
		return db, err
	}
	log.Println("Driver is running.")

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return db, err
	}

	log.Println("Connected to database.")

	return db, nil
}
