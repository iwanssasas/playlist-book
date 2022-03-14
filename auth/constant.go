package auth

import (
	"PLAYLISTBOOK/config"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	oauthStateString = "ini-random"
	googleApis       = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
)

func getGoogleOauth2() *oauth2.Config {
	conf := config.Get()
	return &oauth2.Config{
		RedirectURL:  "http://localhost:8080/callback/google",
		ClientID:     conf.GoogleClientId,
		ClientSecret: conf.GoogleClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

}
