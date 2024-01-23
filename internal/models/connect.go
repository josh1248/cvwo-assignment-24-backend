package models

import (
	"log"

	_ "github.com/glebarez/go-sqlite"
	"github.com/jmoiron/sqlx"
)

// package scope variable representing DB connection.
var db *sqlx.DB

func ConnectToDB() {
	//declare here to prevent overshadowing of db variable:
	//https://stackoverflow.com/questions/34195360/how-to-use-global-var-across-files-in-a-package
	var err error

	//switch between :memory: and a file directory for transient/permanent DBs.
	//sqlx.Connect combines sql.Open with sql.Ping
	db, err = sqlx.Connect("sqlite", "internal/db/blogdb")
	//not needed, since it will be closed upon interruption.
	//defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connection established.")
	log.Printf("%T", db)

	/*
		createCommand := `
			CREATE TABLE users (
				id INT PRIMARY KEY,
				name TEXT NOT NULL,
				reputation INT NOT NULL
			);
		`

		_, err = db.Exec(createCommand)
		if err != nil {
			log.Fatal(err)
		}

	log.Println("Table in database created.")*/
}
