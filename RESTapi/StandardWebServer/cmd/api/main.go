package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/killabayte/Go/RESTapi/StandardWebServer/internal/app/api"
)

var (
	configPath string = "configs/api.toml"
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/api.toml", "path to config file in toml format")
}

func main() {
	log.Println("It works!")
	config := api.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Println("Can't decode config file", err)
	}

	server := api.New(config)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
