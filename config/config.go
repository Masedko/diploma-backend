package config

import (
	"github.com/spf13/viper"
)

type Load struct {
	Path string
	Name string
	Type string
}

type Server struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type Database struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"db-name"`
	Port     int    `mapstructure:"port"`
	SSLMode  string `mapstructure:"ssl-mode"`
}

type Config struct {
	Server   `mapstructure:"server"`
	Database `mapstructure:"database"`
}

func LoadConfig(load Load) (*Config, error) {
	v := viper.New()
	v.SetConfigName(load.Name)
	v.SetConfigType(load.Type)
	v.AddConfigPath(load.Path)

	err := v.ReadInConfig()
	if err != nil {
		return nil, pkgerrors.Error{
			Err:  err,
			Desc: "Cannot read config",
		}
	}

	var config Config
	err = v.Unmarshal(&config)
	if err != nil {
		return nil, pkgerrors.Error{
			Err:  err,
			Desc: "Cannot unmarshal config",
		}
	}

	return &config, nil
}
