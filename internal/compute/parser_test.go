package compute_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"goconcurrency/internal/compute"
	"strings"
	"testing"
)

func TestParse_ValidCommands(t *testing.T) {
	tests := []struct {
		input       string
		expected    *compute.Cmd
		expectedErr string
	}{
		{"SET key value", &compute.Cmd{Type: compute.CmdTypeSet, Key: "key", Value: "value"}, ""},
		{"GET key", &compute.Cmd{Type: compute.CmdTypeGet, Key: "key"}, ""},
		{"DEL key", &compute.Cmd{Type: compute.CmdTypeDel, Key: "key"}, ""},
		{"SET key_123 value_456", &compute.Cmd{Type: compute.CmdTypeSet, Key: "key_123", Value: "value_456"}, ""},
		{"GET key123", &compute.Cmd{Type: compute.CmdTypeGet, Key: "key123"}, ""},
	}

	p := compute.NewParser()
	for _, test := range tests {
		cmd, err := p.Parse(test.input)
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
		{"SET key@ value", "parse: invalid argument key@"},
		{"SET key value!", "parse: invalid argument value!"},
		{"GET key#", "parse: invalid argument key#"},
		{"DEL key!", "parse: invalid argument key!"},
		//{"SET  value", "parse: invalid argument "},
		//{"SET key ", "parse: invalid argument "},
		//{"SET  ", "parse: invalid argument "},
	}

	p := compute.NewParser()
	for _, test := range tests {
		_, err := p.Parse(test.input)
		assert.ErrorContains(t, err, test.expectedErr)
	}
}

func TestParse_LongArguments(t *testing.T) {
	longKey := "key" + strings.Repeat("x", 1000)
	longValue := "value" + strings.Repeat("y", 1000)

	tests := []struct {
		input    string
		expected *compute.Cmd
	}{
		{fmt.Sprintf("SET %s %s", longKey, longValue), &compute.Cmd{Type: compute.CmdTypeSet, Key: longKey, Value: longValue}},
	}

	p := compute.NewParser()
	for _, test := range tests {
		cmd, err := p.Parse(test.input)
		assert.NoError(t, err)
		assert.Equal(t, test.expected, cmd)
	}
}
