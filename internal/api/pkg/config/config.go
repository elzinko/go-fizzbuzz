package config

import (
	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

var Config *Configuration

type Configuration struct {
	Server ServerConfiguration
	Log    LogConfiguration
}

type ServerConfiguration struct {
	Port   string
	Secret string
	Mode   string
}

type LogConfiguration struct {
	Level           string
	AppendToLogFile bool
}

func Load(basePath string, configPath string) *Configuration {
	var configuration *Configuration

	viper.SetConfigFile(basePath + configPath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	Config = configuration

	return Config
}

func GetConfig() *Configuration {
	return Config
}
