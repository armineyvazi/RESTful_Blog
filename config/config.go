package config

import (
	"log"

	"github.com/spf13/viper"
)

var EnvConfigs *envConfigs

func InitEnvConfigs() {
	EnvConfigs = loadEnvVariables()
}

// struct to map env values
type envConfigs struct {
	LocalServerPort string `mapstructure:"LOCAL_SERVER_PORT"`
	SecretKey       string `mapstructure:"SECRET_KEY"`
	DB_USERNAME     string `mapstructure:"DB_USERNAME"`
	DB_PASSWORD     string `mapstructure:"DB_PASSWORD"`
	DB_HOST         string `mapstructure:"DB_HOST"`
	DB_PORT         string `mapstructure:"DB_PORT"`
	DB_DATABASE     string `mapstructure:"DB_DATABASE"`
}

func loadEnvVariables() (config *envConfigs) {
	// Tell viper the path/location of your env file. If it is root just add "."
	viper.AddConfigPath("./")

	// Tell viper the name of your file
	viper.SetConfigName(".env")

	// Tell viper the type of your file
	viper.SetConfigType("env")

	// Viper reads all the variables from env file and log error if any found
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	// Viper unmarshals the loaded env varialbes into the struct
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}
	return
}
