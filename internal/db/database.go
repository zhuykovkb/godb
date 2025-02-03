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
		return "", err
	}

	switch cmd.Type {
	case compute.CmdTypeSet:
		logger.Info("Setting key-value store")
		d.storage.Set(cmd.Args[0], cmd.Args[1])
		return "", nil
	case compute.CmdTypeGet:
		logger.Info("Getting key-value store")
		r, ok := d.storage.Get(cmd.Args[0])
		if !ok {
			return "", errors.New("key not found")
		}
		return r, nil
	case compute.CmdTypeDel:
		logger.Info("Deleting key-value store")
		d.storage.Del(cmd.Args[0])
		return "", nil
	default:
	}
	return "", errors.New("unknown command %s")
}
