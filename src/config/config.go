package config

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	enumconfig "myclass/src/enums/config"
	"os"
)

type structure struct {
	Server struct {
		Port string          `mapstructure:"port"`
		Mode enumconfig.Mode `mapstructure:"mode"`
	} `mapstructure:"server"`
	Postgres struct {
		Url string `mapstructure:"url"`
	} `mapstructure:"postgres"`
	Db *sqlx.DB
}

var cfg structure

func GetConfig() structure {
	return cfg
}

func Load() error {
	path, err := os.Getwd()
	if err != nil {
		return err
	}

	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return err
	}

	db, err := sqlx.Connect("postgres", cfg.Postgres.Url)
	if err != nil {
		return err
	}

	cfg.Db = db

	return nil

}
