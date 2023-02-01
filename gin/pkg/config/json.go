package config

import (
	"log"

	"github.com/tanveerprottoy/rest-api-starter-go/gin/pkg/file"
	"github.com/tanveerprottoy/rest-api-starter-go/gin/pkg/json"
)

var (
	// configs
	Configs map[string]any
)

func init() {
	pwd, _ := file.GetPWD()
	log.Println(pwd)
	b, _ := file.ReadFile(pwd + "/config/dev.json")
	_ = json.Unmarshal(b, &Configs)
	log.Print(Configs)
}

func GetJsonValue(key string) any {
	return Configs[key]
}
