package models

import (
	"database/sql"

	_ "github.com/glebarez/go-sqlite"
)

func ConnectToDB() (*sql.DB, error) {
	//switch between :memory: and a file directory for transient/permanent DBs.
	db, err := sql.Open("sqlite", "internal/db/blogdb")
	//not needed, since it will be closed upon interruption.
	//defer db.Close()

	if err != nil {
		//possible shortening with named and naked returns?
		return db, err
	}

	err = db.Ping()
	if err != nil {
		return db, err
	}

	createCommand := `
		CREATE TABLE users (
			id INT PRIMARY KEY,
			name TEXT NOT NULL,
			reputation INT NOT NULL
		);
	`

	_, err = db.Exec(createCommand)
	if err != nil {
		return db, err
	} else {
		return db, nil
	}
}
