package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func InitMysql() (*sql.DB, error) {
	// init mysql database
	db, err := sql.Open("mysql", os.Getenv("MYSQL_URL"))
	if err != nil {
		log.Fatal("error connecting to database", err.Error())
		return nil, err
	}

	return db, nil
}
