package config

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/inhies/go-bytesize"
	"goconcurrency/internal/logger"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"os"
	"time"
)

var defaultConfig = Config{
	Engine: &EngineConfig{
		Type: "in-memory",
	},
	Network: &NetworkConfig{
		Address:        "127.0.0.1:3223",
		MaxConnections: 100,
		MaxMessageSize: "4096Kb",
		IdleTimeout:    5 * time.Second,
	},
	Logging: &LoggingConfig{
		Level:  "info",
		Output: "app.log",
	},
}

type Config struct {
	Engine  *EngineConfig  `yaml:"engine"`
	Network *NetworkConfig `yaml:"network"`
	Logging *LoggingConfig `yaml:"logging"`
}

type EngineConfig struct {
	Type string `yaml:"type"`
}

type NetworkConfig struct {
	Address        string        `yaml:"address"`
	MaxConnections int           `yaml:"max_connections"`
	MaxMessageSize string        `yaml:"max_message_size"`
	IdleTimeout    time.Duration `yaml:"idle_timeout"`
}
type LoggingConfig struct {
	Level  string `yaml:"level"`
	Output string `yaml:"output"`
}

func (c *NetworkConfig) GetMaxMessageSize() int {
	bs, err := bytesize.Parse(c.MaxMessageSize)
	if err != nil {
		log.Fatal(fmt.Errorf("error parsing max_message_size: %v", err))
	}

	return int(bs)
}

func LoadConfig(configFile string) (*Config, error) {
	config := defaultConfig
	confData, err := os.ReadFile(configFile)
	if err != nil {
		InitLoggerFromConfig(config.Logging)
		return &config, nil
	}

	reader := bytes.NewReader(confData)
	if reader == nil {
		return nil, errors.New("nil reader")
	}

	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("error parsing config: %w", err)
	}

	InitLoggerFromConfig(config.Logging)
	return &config, nil
}

func InitLoggerFromConfig(cfg *LoggingConfig) {
	var outputs []logger.OutputTarget

	outputs = append(outputs, logger.OutputTarget{Writer: os.Stdout})

	logFile, err := os.OpenFile(cfg.Output, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error open log file %s: %v\n", cfg.Output, err)
		os.Exit(1)
	}
	outputs = append(outputs, logger.OutputTarget{Writer: logFile})

	logCfg := logger.Config{
		Level:   cfg.Level,
		Outputs: outputs,
		Format:  "json",
	}

	logger.Init(logCfg)
}
