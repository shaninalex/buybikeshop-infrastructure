package auth

import (
	"buybikeshop/libs/go/config"

	"golang.org/x/oauth2"
)

func ProvideOAuthConfig(config *config.Config) *oauth2.Config {
	c := config.GetOAuthConfig()
	return &oauth2.Config{
		ClientID:     c.ClientID.String(),
		ClientSecret: "",
		RedirectURL:  c.RedirectUri,
		Scopes:       c.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  c.AuthorizationUrl,
			TokenURL: c.TokenUrl,
		},
	}
}
