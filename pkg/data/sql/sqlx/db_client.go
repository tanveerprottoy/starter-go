package sqlx

import (
	"log"
	"os"
	"sync"
	"sync/atomic"

	"github.com/jmoiron/sqlx"

	"github.com/go-sql-driver/mysql"
)

const (
	netType = "tcp"
)

var (
	instance    *DBClient
	once        sync.Once
	mu          sync.Mutex
	initialized uint32
)

type DBClient struct {
	DB *sqlx.DB
}

func GetInstance() *DBClient {
	once.Do(func() {
		instance = new(DBClient)
		instance.init()
	})
	return instance
}

func GetInstanceMutex() *DBClient {
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()
		if instance == nil {
			instance = new(DBClient)
			instance.init()
		}
	}
	return instance
}

func GetInstanceAtomic() *DBClient {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}
	mu.Lock()
	defer mu.Unlock()
	if initialized == 0 {
		instance = new(DBClient)
		instance.init()
		atomic.StoreUint32(&initialized, 1)
	}
	return instance
}

func (d *DBClient) init() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASS"),
		Net:    netType,
		Addr:   os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		DBName: os.Getenv("DB_NAME"),
	}
	var err error
	d.DB, err = sqlx.Open("mysql", cfg.FormatDSN())
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
