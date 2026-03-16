package auth

import (
	"buybikeshop/libs/go/config"

	"golang.org/x/oauth2"
)

func ProvideOAuthConfig(c *config.Config) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     c.String("oauth.client_id"),
		ClientSecret: "",
		RedirectURL:  c.String("oauth.redirect_url"),
		Scopes:       c.StringSlice("oauth.scopes"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  c.String("oauth.authorization_url"),
			TokenURL: c.String("oauth.token_url"),
		},
	}
}
