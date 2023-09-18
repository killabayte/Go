package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/killabayte/Go/RESTapi/StandardWebServer/internal/app/api"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "path", "configs/api.toml", "path to config file")
}

func main() {
	flag.Parse()
	if configPath == "" {
		fmt.Println("No config path specified, using default")
		configPath = "configs/api.toml"
	}
	log.Println("It works!")
	config := api.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Println("Error decoding config file:", err)
	}

	server := api.New(config)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
