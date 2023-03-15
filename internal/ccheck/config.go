package ccheck

import (
	"fmt"

	"github.com/spf13/viper"
)

type CCheckConfig struct{}

func GetCCheckConfig() (config CCheckConfig) {
	viper.SetConfigFile(".ccheck")
	viper.SetConfigType("toml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	viper.Unmarshal(&config)
	return
}
