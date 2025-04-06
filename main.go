package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"photo-saver/server"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/server.json", "config server file path")
}

func main() {
	flag.Parse()

	config := server.NewConfig()

	getConfig(config)

	server := server.NewServer(config)
	err := server.Start()
	if err != nil {
		fmt.Println("server error: ", err)
	}
}

func getConfig(config *server.Config) {
	confFile, err := os.Open(configPath)
	if err != nil {
		fmt.Println("config error: ", err)
	}
	defer confFile.Close()

	decoder := json.NewDecoder(confFile)
	if err := decoder.Decode(&config); err != nil {
		fmt.Println("config error: ", err)
	}
}
