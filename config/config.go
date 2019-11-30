package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var Config ConfigSchema

func init() {
	viper.SetConfigName("config")   // name of config file (without extension)
	viper.AddConfigPath("./config") // optionally look for config in the working directory
	err := viper.ReadInConfig()     // Find and read the config file
	if err != nil {                 // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
}

type ConfigSchema struct {
	Mongo struct {
		URI      string `mapstructure:"URI"`
		Host     string `mapstructure:"Host"`
		Username string `mapstructure:"Username"`
		Password string `mapstructure:"Password"`
	} `mapstructure:"MongoDB"`
	JWTSecret struct {
		JWTKey string `mapstructure:"JWTKey"`
	} `mapstructure:"JWTSecret"`
}
