package compute

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func NewParser() *Parser {
	return &Parser{}
}

type ParserInterface interface {
	Parse(string) (*Cmd, error)
}

type Parser struct{}

func (p *Parser) Parse(cmd string) (*Cmd, error) {
	parts := strings.Fields(cmd)
	if len(parts) == 0 {
		return nil, formatError("parse: undefined command")
	}

	ctype := CmdType(parts[0])
	switch ctype {
	case CmdTypeSet:
		if len(parts) != 3 {
			return nil, formatError("parse: invalid %s command", string(ctype))
		}

		if err := validateArgument(parts[1:]...); err != nil {
			return nil, err
		}

		return &Cmd{
			Type:  ctype,
			Key:   parts[1],
			Value: parts[2],
		}, nil
	case CmdTypeGet, CmdTypeDel:
		if len(parts) != 2 {
			return nil, formatError("parse: invalid %s command", string(ctype))
		}

		if err := validateArgument(parts[1]); err != nil {
			return nil, err
		}

		return &Cmd{
			Type: ctype,
			Key:  parts[1],
		}, nil

	default:
		return nil, formatError("parse: invalid %s command", string(ctype))
	}
}

var validArgPattern = regexp.MustCompile(`^\w+$`)

func validateArgument(args ...string) error {
	for _, arg := range args {
		if !validArgPattern.MatchString(arg) {
			return formatError("parse: invalid argument %s", arg)
		}
	}
	return nil
}

func formatError(msg string, args ...interface{}) error {
	return errors.New(fmt.Sprintf(msg, args...))
}
