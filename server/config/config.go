package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var DBConfig Configuration

// Configuration struct to hold configuration values
type Configuration struct {
	Database struct {
		Host         string `yaml:"host"`
		Port         int    `yaml:"port"`
		Username     string `yaml:"username"`
		Password     string `yaml:"password"`
		DatabaseName string `yaml:"database_name"`
	} `yaml:"database"`
}

// LoadConfig loads configuration from config.yaml file
func LoadConfig() error {
	// Read the YAML configuration file
	data, err := ioutil.ReadFile("config/config.yml")
	if err != nil {
		return fmt.Errorf("failed to read config file: %v", err)
	}

	// Unmarshal YAML into Config struct
	if err := yaml.Unmarshal(data, &DBConfig); err != nil {
		return fmt.Errorf("failed to unmarshal config: %v", err)
	}

	return nil
}
