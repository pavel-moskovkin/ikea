package config

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

const configFileName = "config"

type Config struct {
	ApiListen string
	App       App
	DbConfig  DbConfig `mapstructure:"db"`
	LogLevel  string
}

type App struct {
	usdrub string
}

type DbConfig struct {
	Address  string
	Username string
	Password string
	DBName   string
	Insecure bool
}

// LoadConfig loads config from file
func LoadConfig() (Config, error) {
	v := viper.New()
	v.SetConfigName(configFileName)
	v.SetEnvPrefix("ikea")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	// v.AddConfigPath("./config")

	replacer := strings.NewReplacer(".", "_")
	v.SetEnvKeyReplacer(replacer)
	v.AutomaticEnv()

	var cfg Config
	if err := v.ReadInConfig(); err != nil {
		return Config{}, errors.Wrap(err, "Failed to read config file")
	}

	err := v.Unmarshal(&cfg)
	if err != nil {
		return Config{}, errors.Wrap(err, "Failed to unmarshal config file")
	}

	return cfg, nil
}
