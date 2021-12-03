package config

import "github.com/spf13/viper"

type config struct {
	Database struct {
		DSN      string `mapstructure:"dsn"`
		LogLevel string `mapstructure:"logLevel"`
	} `mapstructure:"database"`
	SecretKey string `mapstructure:"secretKey"`
	Port      string `mapstructure:"port"`
}

var Env config

func Load() error {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return viper.Unmarshal(&Env)
}
