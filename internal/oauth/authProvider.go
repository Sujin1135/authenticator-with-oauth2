package oauth

import (
	"auth/internal/oauth/google"
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"io"
	"net/http"
)

type AuthProvider struct {
	config *oauth2.Config
	apiUrl string
}

type RedirectInfo struct {
	clientId    string
	redirectUrl string
}

func (auth *AuthProvider) Redirect() *RedirectInfo {
	return &RedirectInfo{clientId: auth.config.ClientID, redirectUrl: auth.config.RedirectURL}
}

func (auth *AuthProvider) Callback(authCode string) ([]byte, error) {
	token, err := auth.config.Exchange(context.Background(), authCode)
	if err != nil {
		return nil, fmt.Errorf("failed getting an access token of Oauth 2.0: %s", err.Error())
	}
	response, err := http.Get(auth.apiUrl + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed extract body contents from response body : %s", err.Error())
	}
	return contents, nil
}

var AuthProviders = map[string]*AuthProvider{
	"google": &AuthProvider{config: google.GoogleConfig, apiUrl: google.GoogleApiUrl},
}