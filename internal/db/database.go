package db

import (
	"errors"
	"goconcurrency/internal/compute"
	"goconcurrency/internal/logger"
	"goconcurrency/internal/storage"
	"os"
	"strings"
)

type Db struct {
	storage storage.Storage
	compute compute.ParserInterface
}

func StartDB(s storage.Storage, c compute.ParserInterface) *Db {
	return &Db{
		storage: s,
		compute: c,
	}
}

func (d *Db) HandleReq(in string) (string, error) {
	in = strings.TrimSpace(in)
	if strings.ToLower(in) == "exit" {
		logger.Info("Bye bye!")

		os.Exit(0)
	}

	cmd, err := d.compute.Parse(in)
	if err != nil {
		logger.Warn(err.Error())
		return "", err
	}

	switch cmd.Type {
	case compute.CmdTypeSet:
		logger.Info("Setting key-value store")
		d.storage.Set(cmd.Key, cmd.Value)
		return "", nil
	case compute.CmdTypeGet:
		logger.Info("Getting key-value store")
		r, ok := d.storage.Get(cmd.Key)
		if !ok {
			logger.Warn("key not found")
			return "", errors.New("key not found")
		}
		return r, nil
	case compute.CmdTypeDel:
		logger.Info("Deleting key-value store")
		d.storage.Del(cmd.Key)
		return "", nil
	default:
		logger.Info("Unknown command")
	}
	return "", errors.New("unknown command")
}
