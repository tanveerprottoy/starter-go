package config

import (
	"log"

	"github.com/tanveerprottoy/starter-go/stdlib/pkg/file"
	"github.com/tanveerprottoy/starter-go/stdlib/pkg/jsonpkg"
)

var (
	// configs
	Configs map[string]any
)

func init() {
	pwd, _ := file.GetPWD()
	log.Println(pwd)
	b, _ := file.ReadFile(pwd + "/config/dev.json")
	_ = jsonpkg.Unmarshal(b, &Configs)
	log.Print(Configs)
}

func GetJsonValue(key string) any {
	return Configs[key]
}
