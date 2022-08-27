package config

import "github.com/spf13/viper"

// Load load config
func Load() error {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()
}
