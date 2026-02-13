package auth

import (
	ory "github.com/ory/kratos-client-go"
)

type KratosConfigProvider interface {
	KratosUrlBrowser() string
	KratosUrlAdmin() string
}

func ProvideKratos(config KratosConfigProvider) *ory.APIClient {
	configuration := ory.NewConfiguration()
	configuration.Servers = []ory.ServerConfiguration{
		{
			URL: config.KratosUrlBrowser(),
		},
	}
	return ory.NewAPIClient(configuration)
}
