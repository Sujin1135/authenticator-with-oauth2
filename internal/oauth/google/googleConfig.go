package google

import (
	"golang.org/x/oauth2"
	"os"
)

var GoogleConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:8000/auth/google/callback",
	ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
}

var GoogleApiUrl = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
