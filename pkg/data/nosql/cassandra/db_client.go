package cassandra

import (
	"log"
)

var (
	Client any
)

func InitDBClient() {

	log.Println("Successfully connected!")
}
