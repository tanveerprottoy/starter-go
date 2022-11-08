package mysql

import (
	"database/sql"
	"log"
	"txp/restapistarter/pkg/config"

	"github.com/go-sql-driver/mysql"
)

const (
	netType = "tcp"
)

var (
	DB  *sql.DB
	err error
)

func InitDBClient() {
	cfg := mysql.Config{
		User:   config.GetEnvValue("DB_USER"),
		Passwd: config.GetEnvValue("DB_PASS"),
		Net:    netType,
		Addr:   config.GetEnvValue("DB_HOST") + ":" + config.GetEnvValue("DB_PORT"),
		DBName: config.GetEnvValue("DB_NAME"),
	}
	DB, err = sql.Open("mysql", cfg.FormatDSN())
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
