package main

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/killabayte/Go/RESTapi/StandardWebServer/internal/app/api"
)

var (
	configPath string = "configs/api.toml"
)

func init() {
}

func main() {
	log.Println("It works!")
	config := api.NewConfig()
	_, err := toml.DecodeFile(configPath, config)

	server := api.New(config)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
