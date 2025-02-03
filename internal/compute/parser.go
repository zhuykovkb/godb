package compute

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const SetCmdArgCount = 2
const GetDelCmdArgCount = 1

func NewParser() *Parser {
	return &Parser{}
}

type ParserInterface interface {
	Parse(string) (*Cmd, error)
}

type Parser struct{}

func (p *Parser) Parse(cmd string) (*Cmd, error) {
	command, err := getCommand(cmd)
	if err != nil {
		return nil, err
	}

	err = validateArgs(command)
	if err != nil {
		return nil, err
	}

	return command, nil
}

func validateArgs(command *Cmd) error {
	switch command.Type {
	case CmdTypeSet:
		if len(command.Args) != SetCmdArgCount {
			return formatError("invalid arg count: %d need: %d", len(command.Args), SetCmdArgCount)
		}

		if err := validateArgument(command.Args...); err != nil {
			return err
		}
	case CmdTypeGet, CmdTypeDel:
		if len(command.Args) != GetDelCmdArgCount {
			return formatError("invalid arg count: %d need: %d", len(command.Args), GetDelCmdArgCount)
		}

		if err := validateArgument(command.Args...); err != nil {
			return err
		}

	default:
		return formatError("parse: invalid %s command", string(command.Type))
	}
	return nil
}

func getCommand(cmd string) (*Cmd, error) {
	if cmd == "" {
		return nil, formatError("empty command")
	}

	parts := strings.Fields(cmd)
	if len(parts) == 0 {
		return nil, formatError("parse: undefined command %s", cmd)
	}

	return &Cmd{
		Type: CmdType(parts[0]),
		Args: parts[1:],
	}, nil
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
