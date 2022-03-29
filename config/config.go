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
	RoleId          int           `envconfig:"DEFAULT_ROLE_ID" default:"2"`
	Secret          string        `envconfig:"SECRET" default:"PLAYLISTBOOK"`
	ExpiredDuration time.Duration `envconfig:"EXP_DURATION" default:"10h"`
	LimitPage       int           `envconfig:"LIMITPAGE" default:"2"`
	IsEditedGoogle  bool          `envconfig:"IsEditedTrue" default:"true"`

	GoogleClientId     string `envconfig:"GOOGLE_CLIENT_ID" default:"120468716859-taqf0kjaci78btltfbkii2eoogi00gan.apps.googleusercontent.com"`
	GoogleClientSecret string `envconfig:"GOOGLE_CLIENT_SECRET" default:"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"`
}

func Init() {
	conf = new(Config)
	envconfig.MustProcess("", conf)
}

func Get() *Config {
	return conf
}

// LOGIN FALSE
// OAUTH TRUE
