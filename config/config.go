package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

var conf *Config

type Config struct {
	Port            int           `envconfig:"PORT" default:"8080"`
	DbName          string        `envconfig:"DB_NAME" default:"PLAYLISTBOOK"`
	DbHost          string        `envconfig:"DB_HOST" default:"localhost"`
	DbPort          int           `envconfig:"DB_PORT" default:"3306"`
	DbUsername      string        `envconfig:"DB_USER" default:"root2"`
	DbPassword      string        `envconfig:"DB_PASS" default:"root2"`
	RoleId          int           `envconfig:"DEFAULT_ROLE_ID" default:"1"`
	Secret          string        `envconfig:"SECRET" default:"PLAYLISTBOOK"`
	ExpiredDuration time.Duration `envconfig:"EXP_DURATION" default:"10h"`
}

func Init() {
	conf = new(Config)
	envconfig.MustProcess("", conf)
}

func Get() *Config {
	return conf
}
