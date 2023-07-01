package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"sync/atomic"

	"github.com/tanveerprottoy/starter-go/stdlib/pkg/config"

	_ "github.com/lib/pq"
)

var (
	instance    *Client
	once        sync.Once
	mu          sync.Mutex
	initialized uint32
)

type Client struct {
	DB *sql.DB
}

func GetInstance() *Client {
	once.Do(func() {
		instance = new(Client)
		instance.init()
	})
	return instance
}

func GetInstanceMutex() *Client {
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()
		if instance == nil {
			instance = new(Client)
			instance.init()
		}
	}
	return instance
}

func GetInstanceAtomic() *Client {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}
	mu.Lock()
	defer mu.Unlock()
	if initialized == 0 {
		instance = new(Client)
		instance.init()
		atomic.StoreUint32(&initialized, 1)
	}
	return instance
}

func (d *Client) init() {
	args := fmt.Sprintf(
		"host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
		config.GetEnvValue("DB_HOST"),
		config.GetEnvValue("DB_PORT"),
		config.GetEnvValue("DB_USER"),
		config.GetEnvValue("DB_PASS"),
		config.GetEnvValue("DB_NAME"),
	)
	var err error
	d.DB, err = sql.Open("postgres", args)
	if err != nil {
		panic(err)
	}
	// ping is necessary to create connection
	err = d.DB.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("Successfully connected!")
}
