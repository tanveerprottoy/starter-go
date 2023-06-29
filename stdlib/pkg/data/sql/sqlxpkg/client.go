package sqlxpkg

import (
	"fmt"
	"log"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/tanveerprottoy/starter-go/stdlib/pkg/config"
)

var (
	instance *Client
	once     sync.Once
)

type Client struct {
	DB *sqlx.DB
}

func GetInstance() *Client {
	once.Do(func() {
		instance = new(Client)
		instance.init()
	})
	return instance
}

func (d *Client) init() {
	// connection properties.
	info := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.GetEnvValue("DB_HOST"), config.GetEnvValue("DB_PORT"), config.GetEnvValue("DB_USER"), config.GetEnvValue("DB_PASS"), config.GetEnvValue("DB_NAME"))
	var err error
	d.DB, err = sqlx.Open("postgres", info)
	if err != nil {
		panic(err)
	}
	// ping is necessary to create connection
	err = d.DB.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("Successfully connected!")
	// create table if not exists
	_, err = d.DB.Exec("CREATE TABLE IF NOT EXISTS books (id uuid PRIMARY KEY DEFAULT gen_random_uuid(), title VARCHAR NOT NULL, author VARCHAR NOT NULL, publication_year INT, created_at BIGINT, updated_at BIGINT)")
	if err != nil {
		panic(err)
	}

}
