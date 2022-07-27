package config

import (
	"fmt"

	"fioepq9.cn/checkin_ecycloud/logger"
	"github.com/spf13/viper"
)

type login struct {
	Url string `mapstructure:"url"`
}

type checkin struct {
	Url string `mapstructure:"url"`
}

type user struct {
	Name     string `mapstructure:"name"`
	Shortcut string `mapstructure:"shortcut"`
	Email    string `mapstructure:"email"`
	Passwd   string `mapstructure:"passwd"`
}

type config struct {
	Login   login   `mapstructure:"login"`
	Checkin checkin `mapstructure:"checkin"`
	Users   []user  `mapstructure:"user"`
}

var C config

func init() {
	viper.SetConfigFile("./config.yaml")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		logger.L.Info("")
		panic(fmt.Errorf("Fatal error in config file: %s \n", err))
	}
	err = viper.Unmarshal(&C)
	if err != nil {
		panic(err)
	}
}
