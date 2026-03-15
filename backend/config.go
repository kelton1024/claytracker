package main

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort  string
	Host     string
	DBPort   string
	Name     string
	User     string
	Password string
}

func loadConfig(fileName string) (*Config, error) {

	v := viper.New()

	// Probably not idiomatic
	err := v.BindEnv("APP_PORT")
	err = v.BindEnv("HOST")
	err = v.BindEnv("DB_PORT")
	err = v.BindEnv("NAME")
	err = v.BindEnv("USER")
	err = v.BindEnv("PASSWORD")
	if err != nil {
		return nil, err
	}

	// Set up viper to read config file
	v.SetConfigFile(fileName)

	// Read config file
	err = v.ReadInConfig()
	if err != nil {
		// There is a viper specific error value we could check, but
		// I think it has a bug so this should work
		if !strings.Contains(err.Error(), "not found") {
			return nil, err
		}
	}

	// Get values that match config struct from file
	var config Config
	err = v.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
