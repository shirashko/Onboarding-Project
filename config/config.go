package config

import (
	"fmt"
	"github.com/spf13/viper" // a popular configuration management library for Go.
)

// DBConfig holds the info needed for connecting to the database
// It's annotated with mapstructure tags to indicate how configuration keys should be mapped to struct fields.
type DBConfig struct {
	Host         string `mapstructure:"HOST"`
	Port         int    `mapstructure:"PORT"`
	User         string `mapstructure:"USER"`
	Password     string `mapstructure:"PASSWORD"`
	DatabaseName string `mapstructure:"DATABASE_NAME"`
}

// ServerConfig Holds information about the server (host and port). using this in the listen and serve part
type ServerConfig struct {
	Host string `mapstructure:"HOST"`
	Port int    `mapstructure:"PORT"`
}

// Config The main configuration struct that contains nested fields for database connection information and server information.
// It's also annotated with mapstructure tags.
type Config struct {
	DBConfig   DBConfig     `mapstructure:"DB"`
	ServerInfo ServerConfig `mapstructure:"SERVER"`
}

// LoadConfig load the configuration from various sources and populate the given Config struct
func LoadConfig(config interface{}) error {

	v := viper.New() // Initializes a new instance of the viper configuration manager.

	// Configuration settings
	v.SetConfigName("config")
	v.SetConfigType("json")
	v.AddConfigPath("./config")
	v.AutomaticEnv()

	// Read and unmarshal the configuration into config struct argument
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read configuration: %w", err)
	}
	if err := v.Unmarshal(&config); err != nil {
		return fmt.Errorf("failed to unmarshal ServerInfo: %w", err)
	}

	return nil
}
