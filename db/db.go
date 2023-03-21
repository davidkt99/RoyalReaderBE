package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "goServer"
	password = "1234"
	dbname   = "royalRoadReaderDB"
)

func DBSetup() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")

	// hardcoded
	insertStmt := `insert into "Chapters"("name", "content", "read") values('CH14', 'Content: Hello World!', true)`
	_, e := db.Exec(insertStmt)
	CheckError(e)

	rows, err := db.Query(`SELECT "name","content" FROM "Chapters"`)
	CheckError(err)

	defer rows.Close()
	for rows.Next() {
		var name string
		var content string

		err = rows.Scan(&name, &content)
		CheckError(err)

		fmt.Println(name, content)
	}

	// // dynamic
	// insertDynStmt := `insert into "Students"("Name", "Roll") values($1, $2)`
	// _, e = db.Exec(insertDynStmt, "Jane", 2)
	// CheckError(e)

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
