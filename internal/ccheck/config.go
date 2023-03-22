package ccheck

import (
	"fmt"

	"github.com/spf13/viper"
)

type CCheckConfig struct {
	copyright *CCheckCopyright
}

func GetCCheckConfig() (config *CCheckConfig) {
	viper.SetConfigFile(".ccheck")
	viper.SetConfigType("toml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	viper.Unmarshal(&config)
	config.copyright = NewCCheckCopyright(&[]string{
		"Copyright (c) 2023 Taylor Steinberg",
	})
	return
}
