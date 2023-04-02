package db

import (
	"fmt"

	"github.com/davidkt99/RoyalReaderBE/models"
	"github.com/davidkt99/RoyalReaderBE/util"
)

func QueryUserVerifiyPassword(userName string, password string) int64 {
	insertStmt := `
	select * from users
	where users.user_name=$1`

	var user models.User

	rows, e := db.Query(insertStmt, userName)
	CheckError(e)

	defer rows.Close()
	rows.Next()
	eScan := rows.Scan(&user.Id, &user.UserName, &user.Email, &user.Password)
	CheckError(eScan)

	if !util.ComparePasswords(user.Password, []byte(password)) {
		fmt.Println("Password Incorrect")
		return -1
	}

	fmt.Println("Login Successful: ", user.Id)

	return user.Id
}
