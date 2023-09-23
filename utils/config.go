package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Debug bool `json:"debug"`
	Log   bool `json:"log"`
	Port  int  `json:"port"`
}

func LoadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
