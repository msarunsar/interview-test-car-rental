package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitCFG(config *AppConfig) error {
	viper.SetConfigName("config")
	// Add the path to look for the config file
	viper.AddConfigPath(".")
	// Read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %s\n", err)
		return err
	}

	// Define a struct to hold the configuration

	// Unmarshal the configuration from the config file into the struct
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Error unmarshalling config: %s\n", err)
		return err
	}

	return nil
}
