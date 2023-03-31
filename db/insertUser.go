package db

import (
	"fmt"

	"github.com/davidkt99/RoyalReaderBE/models"
	"github.com/davidkt99/RoyalReaderBE/util"
)

func InsertUser(user models.User) int64 {
	insertStmt := `insert into "users"("user_name", "user_password", "user_email") values($1, $2, $3) RETURNING user_id`
	var id int64
	e := db.QueryRow(insertStmt, user.UserName, util.HashAndSalt([]byte(user.Password)), user.Email).Scan(&id)
	CheckError(e)
	fmt.Println("User Inserted: ", id)
	return id
}
