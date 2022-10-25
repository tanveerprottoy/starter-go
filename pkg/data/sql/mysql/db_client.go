package mysql

import (
	"database/sql"
	"log"
	"txp/restapistarter/pkg/configutil"

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
		User:   configutil.GetEnvValue("DB_USER"),
		Passwd: configutil.GetEnvValue("DB_PASS"),
		Net:    netType,
		Addr:   configutil.GetEnvValue("DB_HOST") + ":" + configutil.GetEnvValue("DB_PORT"),
		DBName: configutil.GetEnvValue("DB_NAME"),
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
