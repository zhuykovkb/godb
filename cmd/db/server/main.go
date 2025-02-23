package main

import (
	"fmt"
	"goconcurrency/internal/config"
	"goconcurrency/internal/server"
	"log"
	"os"
)

var configFile = os.Getenv("CONFIG_FILE")

func main() {
	log.Println(fmt.Sprintf("Server starting with PID: %d", os.Getpid()))

	cfg, err := config.LoadConfig(configFile)
	if err != nil {
		log.Fatal(err)
	}

	serv, err := server.NewServer(cfg)

	if err != nil {
		log.Fatal(err)
	}

	serv.Run()
}
