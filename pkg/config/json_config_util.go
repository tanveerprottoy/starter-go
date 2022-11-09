package config

import (
	"log"
	"os"
	"txp/restapistarter/pkg/json"
)

var (
	// configs
	Configs map[string]interface{}
)

func init() {
	fileBytes, _ := os.ReadFile("../config/dev.json")
	_ = json.Unmarshal(fileBytes, &Configs)
	log.Print(Configs)
}

func GetJsonValue(key string) string {
	return ""
}
