package config

import (
	"github.com/kelseyhightower/envconfig"
)

var conf *Config

type Config struct {
	Port int `envconfig:"PORT" default:"8080"`
}

func Init() {
	conf = new(Config)
	envconfig.MustProcess("", conf)
}

func Get() *Config {
	return conf
}
