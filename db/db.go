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

	// hardcoded
	// insertStmt := `insert into "Chapters"("name", "content", "read") values($1, $2, false) RETURNING id`
	// id := 0
	// e := db.QueryRow(insertStmt, "testName", "testContent").Scan(&id)
	// CheckError(e)

	// fmt.Println(id)

	// rows, err := db.Query(`SELECT "name","content" FROM "Chapters"`)
	// CheckError(err)

	// defer rows.Close()
	// for rows.Next() {
	// 	var name string
	// 	var content string

	// 	err = rows.Scan(&name, &content)
	// 	CheckError(err)

	// 	fmt.Println(name, content)
	// }

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
