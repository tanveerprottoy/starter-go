package sqlx

import (
	"github.com/jmoiron/sqlx"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

const (
	netType = "tcp"
)

var (
	DB  *sqlx.DB
	err error
)

func InitDBClient() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASS"),
		Net:    netType,
		Addr:   os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		DBName: os.Getenv("DB_NAME"),
	}
	DB, err = sqlx.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}
	// ping is necessary to create connection
	err = DB.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("Successfully connected!")
}
