package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB *sql.DB
)

func Init() {
	fmt.Println("[DB] Init")
	// example: db_user:password@tcp(localhost:3306)/my_db
	db, err := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_DATABASE"))
	if err != nil {
		panic(err.Error())
	}
	DB = db

	initUserDB()
}

func Close() {
	DB.Close()
}
