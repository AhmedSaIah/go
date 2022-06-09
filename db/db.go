package main

import (
	"database/sql"
	"fmt"
)

//Not done yet

func main() {
	connStr := "user=postgres dbname=bank password=01000 host=localhost sslmode=disable"
	//port 5432

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nSuccessfully connected to database!\n")
}
