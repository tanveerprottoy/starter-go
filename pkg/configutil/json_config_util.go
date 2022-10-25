package configutil

import (
	"log"
	"os"
	"txp/restapistarter/pkg/coreutil"
)

var (
	// configs
	Configs       map[string]interface{}
)

func init() {
	fileBytes, _ := os.ReadFile("../config/dev.json")
	_ = coreutil.Unmarshal(fileBytes, &Configs)
	log.Print(Configs)
}

func GetJsonValue(key string) string {
	return ""
}
