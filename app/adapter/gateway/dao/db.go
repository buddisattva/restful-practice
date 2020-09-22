package dao

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	// dbCon is the connection pool handler
	dbCon *sqlx.DB
)

func init() {
	var err error
	// TODO: load DB connection information from .env file
	dsn := "root:5566@tcp(localhost:3306)/restful"
	dbCon, err = sqlx.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}

	dbCon.SetConnMaxLifetime(time.Minute * 3)
	dbCon.SetMaxOpenConns(10)
	dbCon.SetMaxIdleConns(10)

	err = dbCon.Ping()
	if err != nil {
		log.Fatalf("Error on opening database connection: %s", err.Error())
	}
}
