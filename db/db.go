package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "goServer"
	password = "1234"
	dbname   = "postgres"
)

func DBSetup() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	var err error
	db, err = sql.Open("postgres", psqlconn)
	CheckError(err)

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")

}

// close database
func DBShutDown() {
	db.Close()
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
