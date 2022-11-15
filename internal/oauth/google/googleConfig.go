package google

import (
	"golang.org/x/oauth2"
	"os"
)

func GetConfig() *oauth2.Config {
	return &oauth2.Config{
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	}
}

var GoogleApiUrl = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
