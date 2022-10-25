package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"txp/restapistarter/pkg/configutil"

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
		configutil.GetEnvValue("DB_HOST"),
		configutil.GetEnvValue("DB_PORT"),
		configutil.GetEnvValue("DB_USER"),
		configutil.GetEnvValue("DB_PASS"),
		configutil.GetEnvValue("DB_NAME"),
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
