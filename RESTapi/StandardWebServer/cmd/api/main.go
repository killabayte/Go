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
	format     string
)

func init() {
	flag.StringVar(&configPath, "path", "", "path to config file")
	flag.StringVar(&format, "format", "", "format of config file, can be .env or .toml extension")
}

func main() {
	flag.Parse()
	config := api.NewConfig()

	if configPath == "" || format == "" {
		fmt.Println("No config path specified, using default")
		configPath = "configs/api.toml"
		_, err := toml.DecodeFile(configPath, config)
		if err != nil {
			log.Println("Error decoding config file:", err)
		}
		fmt.Println(config)
	}

	log.Println("It works!")
	server := api.New(config)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
