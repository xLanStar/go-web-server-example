package db

import (
	"database/sql"
	"fmt"
)

type User struct {
	Account  string `json:"account"`
	Password string `json:"-"`
}

var (
	stmtCreateUser       *sql.Stmt
	stmtGetUserByAccount *sql.Stmt
)

func initUserDB() {
	fmt.Println("DB:", DB)
	var err error
	stmtCreateUser, err = DB.Prepare("INSERT INTO user(Account, Password) VALUES(?, ?)")
	if err != nil {
		panic(err)
	}
	stmtGetUserByAccount, err = DB.Prepare("SELECT * FROM user WHERE Account = ?")
	if err != nil {
		panic(err)
	}
}

func CreateUser(Account, Password string) *User {
	// TODO Create SQL
	// result, err := stmtCreateUser.Exec()
	return &User{
		Account:  Account,
		Password: Password,
	}
}

func GetUserByAccount(Account string) *User {
	// TODO Query SQL
	// result, err := stmtGetUserByAccount.Query(Account)
	return &User{
		Account:  "?",
		Password: "?",
	}
}
