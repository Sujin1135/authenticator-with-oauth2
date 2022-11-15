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
	ClientId    string
	RedirectUrl string
}

func (auth *AuthProvider) Redirect() *RedirectInfo {
	return &RedirectInfo{ClientId: auth.config.ClientID, RedirectUrl: auth.config.RedirectURL}
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

var authProviders map[string]*AuthProvider

func InitAuthProviders() {
	authProviders = map[string]*AuthProvider{
		"google": {config: google.GetConfig(), apiUrl: google.GoogleApiUrl},
	}
}

func GetAuthProvider(name string) (*AuthProvider, error) {
	provider := authProviders[name]
	if provider == nil {
		return nil, fmt.Errorf("Not existed a provider by name as %s", name)
	}

	return provider, nil
}
