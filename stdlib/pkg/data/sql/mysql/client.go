package mysql

import (
	"database/sql"
	"log"
	"sync"
	"sync/atomic"

	"github.com/tanveerprottoy/starter-go/pkg/config"

	"github.com/go-sql-driver/mysql"
)

const (
	netType = "tcp"
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
	cfg := mysql.Config{
		User:   config.GetEnvValue("DB_USER"),
		Passwd: config.GetEnvValue("DB_PASS"),
		Net:    netType,
		Addr:   config.GetEnvValue("DB_HOST") + ":" + config.GetEnvValue("DB_PORT"),
		DBName: config.GetEnvValue("DB_NAME"),
	}
	var err error
	d.DB, err = sql.Open("mysql", cfg.FormatDSN())
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
