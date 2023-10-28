package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
	"github.com/killabayte/Go/RESTapi/ServerAndDB/internal/app/api"
	_ "github.com/lib/pq"
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

	//Setup for default config file and local .toml config file
	if configPath == "" || format == "" {
		fmt.Println("No config path specified, using local .toml file as default")
		_, err := os.Stat("configs/api.toml")
		if err == nil {
			configPath = "configs/api.toml"
			_, err := toml.DecodeFile(configPath, config)
			if err != nil {
				log.Println("Error decoding config file:", err)
			}
		} else {
			fmt.Println("No config file provided and no config file found locally, using default values from the code")
		}
		//Setup for .toml config file
	} else if format == ".toml" {
		fmt.Println("Got an .toml format config file")
		_, err := toml.DecodeFile(configPath, config)
		if err != nil {
			log.Println("Error decoding provided .toml config file:", err)
		}
		//Setup for .env config file
	} else if format == ".env" {
		fmt.Println("Got an .env format config file")
		err := godotenv.Load(configPath)
		if err != nil {
			log.Println("Error decoding provided .env config file:", err)
		}
		bind_addr := os.Getenv("bind_addr")
		logger_level := os.Getenv("logger_level")
		config.BindAddr = bind_addr
		config.LoggerLevel = logger_level
	}

	server := api.New(config)

	log.Fatal(server.Start())
}
