package main

import (
	"bufio"
	"fmt"
	"goconcurrency/internal/compute"
	"goconcurrency/internal/db"
	"goconcurrency/internal/logger"
	inmemory "goconcurrency/internal/storage/inMemory"
	"os"
)

func main() {
	//todo pass level into logger
	logger.Init()

	logger.Info("This is my in-memory key-value store")

	s := inmemory.NewEngine()
	c := compute.NewParser()

	database := db.StartDB(s, c)

	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("-> ")
		in, err := r.ReadString('\n')
		if err != nil {
			logger.Warn(err.Error())
			continue
		}

		res, err := database.HandleReq(in)

		if err != nil {
			logger.Warn(err.Error())
			continue
		}

		fmt.Println(res)

	}
}
