package main

import (
	"bufio"
	"fmt"
	"go.uber.org/zap"
	"goconcurrency/internal/engine"
	"goconcurrency/internal/model"
	"goconcurrency/internal/parser"
	"os"
	"strings"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)

	logger.Info("This is my in-memory key-value store")
	en := engine.New()
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("-> ")
		in, err := r.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		in = strings.TrimSpace(in)
		if strings.ToLower(in) == "exit" {
			logger.Info("Bye bye!")
			break
		}

		cmd, err := parser.Parse(in)
		if err != nil {
			logger.Warn(err.Error())
			continue
		}

		switch cmd.Type {
		case model.CmdTypeSet:
			logger.Info("Setting key-value store")
			en.Set(cmd.Key, cmd.Value)
		case model.CmdTypeGet:
			logger.Info("Getting key-value store")
			r, ok := en.Get(cmd.Key)
			if !ok {
				logger.Info("key not found")
				continue
			}
			logger.Info(r)
		case model.CmdTypeDel:
			logger.Info("Deleting key-value store")
			ok := en.Del(cmd.Key)
			if !ok {
				logger.Info("Key not found")
				continue
			}
			continue
		default:
			logger.Info("Unknown command")
		}

	}
}
