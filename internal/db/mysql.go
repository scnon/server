package db

import (
	"log"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/jmoiron/sqlx"          // sqlx extensions
)

var (
	DB *sqlx.DB
)

func init() {
	var err error
	DB, err = sqlx.Open("mysql", "user:password@/database")
	if err != nil {
		log.Fatal(err)
	}
	DB.SetMaxOpenConns(10)
}
