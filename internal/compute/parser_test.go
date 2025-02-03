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
		input    string
		expected *compute.Cmd
	}{
		{"SET key value", &compute.Cmd{Type: compute.CmdTypeSet, Args: []string{"key", "value"}}},
		{"GET key", &compute.Cmd{Type: compute.CmdTypeGet, Args: []string{"key"}}},
		{"DEL key", &compute.Cmd{Type: compute.CmdTypeDel, Args: []string{"key"}}},
		{"SET key_123 value_456", &compute.Cmd{Type: compute.CmdTypeSet, Args: []string{"key_123", "value_456"}}},
		{"GET key123", &compute.Cmd{Type: compute.CmdTypeGet, Args: []string{"key123"}}},
	}

	p := compute.NewParser()
	for _, test := range tests {
		cmd, err := p.Parse(test.input)
		assert.NoError(t, err)
		assert.Equal(t, test.expected, cmd)
	}
}

func TestParse_InvalidCommands(t *testing.T) {
	tests := []struct {
		input       string
		expectedErr string
	}{
		{"", "empty command"},
		{"INVALID", "parse: invalid INVALID command"},
		{"SET key", "invalid arg count: 1 need: 2"},
		{"GET", "invalid arg count: 0 need: 1"},
		{"DEL", "invalid arg count: 0 need: 1"},
		{"SET key value extra", "invalid arg count: 3 need: 2"},
		{"GET key extra", "invalid arg count: 2 need: 1"},
		{"DEL key extra", "invalid arg count: 2 need: 1"},

		{"SET key@ value", "parse: invalid argument key@"},
		{"SET key value!", "parse: invalid argument value!"},
		{"GET key#", "parse: invalid argument key#"},
		{"DEL key!", "parse: invalid argument key!"},

		{"SET  value", "invalid arg count: 1 need: 2"},
		{"SET key ", "invalid arg count: 1 need: 2"},
		{"SET  ", "invalid arg count: 0 need: 2"},
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
		{fmt.Sprintf("SET %s %s", longKey, longValue), &compute.Cmd{Type: compute.CmdTypeSet, Args: []string{longKey, longValue}}},
	}

	p := compute.NewParser()
	for _, test := range tests {
		cmd, err := p.Parse(test.input)
		assert.NoError(t, err)
		assert.Equal(t, test.expected, cmd)
	}
}
