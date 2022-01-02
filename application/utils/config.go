package utils

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Mongodb struct {
		Uri      string `mapstructure:"uri"`
		Database string `mapstructure:"database"`
	} `mapstructure:"mongodb"`
	Currency struct {
		Url       string `mapstructure:"url"`
		Endpoints struct {
			Live string `mapstructure:"live"`
		} `mapstructure:"endpoints"`
		ApiKey string `mapstructure:"api_key"`
	} `mapstructure:"currency"`
}

var configCache *Config

func LoadConfig() (*Config, error) {
	if configCache != nil {
		return configCache, nil
	}
	viper.SetConfigName("config")
	viper.AddConfigPath("configs")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	log.Println("load config file")
	err = viper.Unmarshal(&configCache)
	return configCache, nil
}
