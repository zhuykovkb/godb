package parser

import (
	"errors"
	"fmt"
	"goconcurrency/internal/model"
	"strings"
)

type Command struct {
	Type  model.CmdType
	Key   string
	Value string
}

func Parse(cmd string) (*Command, error) {
	parts := strings.Fields(cmd)
	if len(parts) == 0 {
		return nil, errors.New("parse: undefined command")
	}

	ctype := model.CmdType(parts[0])
	switch ctype {
	case model.CmdTypeSet:
		if len(parts) != 3 {
			return nil, errors.New(fmt.Sprintf("parse: invalid %s command", string(ctype)))
		}
		return &Command{
			Type:  ctype,
			Key:   parts[1],
			Value: parts[2],
		}, nil
	case model.CmdTypeGet, model.CmdTypeDel:
		if len(parts) != 2 {
			return nil, errors.New(fmt.Sprintf("parse: invalid %s command", string(ctype)))
		}
		return &Command{
			Type: ctype,
			Key:  parts[1],
		}, nil

	default:
		return nil, errors.New(fmt.Sprintf("parse: invalid %s command", string(ctype)))
	}
}
