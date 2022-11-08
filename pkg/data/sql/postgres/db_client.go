package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"txp/restapistarter/pkg/config"

	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func InitDBClient() {
	args := fmt.Sprintf(
		"host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
		config.GetEnvValue("DB_HOST"),
		config.GetEnvValue("DB_PORT"),
		config.GetEnvValue("DB_USER"),
		config.GetEnvValue("DB_PASS"),
		config.GetEnvValue("DB_NAME"),
	)
	DB, err = sql.Open("postgres", args)
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
