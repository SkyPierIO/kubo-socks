package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Port      int `json:"port"`
	SocksPort int `json:"socksPort"`
}

func LoadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	if err != nil {
		fmt.Println("Loading default config...")
		config = Config{8081, 1080}
		return config
	} else {
		defer configFile.Close()
		jsonParser := json.NewDecoder(configFile)
		jsonParser.Decode(&config)
		return config
	}
}
