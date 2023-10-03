package config

import "github.com/spf13/viper"

// store all configuration of the application read from config file or environment variables
type Config struct {
	ServerAddress string `mapstructure:"SERVER_ADDR"`
	DbSource string `mapstructure:"DB_SOURCE"`
	DbDriver string `mapstructure:"DB_DRIVER"`
}

// read configuration from file or environment variables
func LoadConfig(path string) (config Config, err error) {
	// directory to look for file from
	viper.AddConfigPath(path)

	// name of file to read env vars
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	// if there are environment variables, they will overwrite the values from config file
	viper.AutomaticEnv()

	// read from the source of the environment variables
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	// map the env variables back to Config struct
	err = viper.Unmarshal(&config)

	return
}
