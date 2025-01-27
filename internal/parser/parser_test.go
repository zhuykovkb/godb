package parser_test

import (
	"github.com/stretchr/testify/assert"
	"goconcurrency/internal/model"
	"goconcurrency/internal/parser"
	"testing"
)

func TestParse_ValidCommands(t *testing.T) {
	tests := []struct {
		input       string
		expected    *parser.Command
		expectedErr string
	}{
		{"SET key value", &parser.Command{Type: model.CmdTypeSet, Key: "key", Value: "value"}, ""},
		{"GET key", &parser.Command{Type: model.CmdTypeGet, Key: "key"}, ""},
		{"DEL key", &parser.Command{Type: model.CmdTypeDel, Key: "key"}, ""},
	}

	for _, test := range tests {
		cmd, err := parser.Parse(test.input)
		if test.expectedErr != "" {
			assert.ErrorContains(t, err, test.expectedErr)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.expected, cmd)
		}
	}
}

func TestParse_InvalidCommands(t *testing.T) {
	tests := []struct {
		input       string
		expectedErr string
	}{
		{"", "parse: undefined command"},
		{"INVALID", "parse: invalid INVALID command"},
		{"SET key", "parse: invalid SET command"},
		{"GET", "parse: invalid GET command"},
		{"DEL", "parse: invalid DEL command"},
		{"SET key value extra", "parse: invalid SET command"},
		{"GET key extra", "parse: invalid GET command"},
		{"DEL key extra", "parse: invalid DEL command"},
	}

	for _, test := range tests {
		_, err := parser.Parse(test.input)
		assert.ErrorContains(t, err, test.expectedErr)
	}
}
