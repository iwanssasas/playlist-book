package config

import (
	"github.com/kelseyhightower/envconfig"
)

var conf *Config

type Config struct {
	Port       int    `envconfig:"PORT" default:"8080"`
	DbName     string `envconfig:"DB_NAME" default:"PLAYLISTBOOK"`
	DbHost     string `envconfig:"DB_HOST" default:"127.0.0.1"`
	DbPort     int    `envconfig:"DB_PORT" default:"3306"`
	DbUsername string `envconfig:"DB_USER" default:"root2"`
	DbPassword string `envconfig:"DB_PASS" default:"root2"`
}

func Init() {
	conf = new(Config)
	envconfig.MustProcess("", conf)
}

func Get() *Config {
	return conf
}
