package main

import (
	"database/sql"
	"log"
)

var db *sql.DB

func DBConnect(addr string) {
	//...
	db = &sql.DB{}
}

func createUsersTable() error {
	q := `CREATE TABLE IF NOT EXIST users(
		ID SERIAL PRIMARY KEY,
		fist_name VARCHAR(255),
		last_name VARCHAR(255),
		role 	INT,
		email  VARCHAR(255)
	)`

	_, err := db.Exec(q)
	return err
}

func main() {
	err := createUsersTable()
	if err != nil {
		log.Fatal("error creating users table: ", err)
	}
}
