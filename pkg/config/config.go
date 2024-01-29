package config

import (
	"github.com/spf13/viper"
)

type Configure struct {
	PORT         string `mapstructure:"PORT"`
	ADMINPORT    string `mapstructure:"ADMINPORT"`
	BSERVICEPORT string `mapstructure:"BSERVICEPORT"`
	SECRETKEY    string `mapstructure:"SECRETKEY"`
	REDISHOST    string `mapstructure:"REDISHOST"`
}

func LoadConfigure() (*Configure, error) {
	var cfg Configure

	viper.SetConfigFile("../../.env")
	err := viper.ReadInConfig()

	err = viper.Unmarshal(&cfg)

	if err != nil {
		return &Configure{}, nil
	}

	return &cfg, nil
}
