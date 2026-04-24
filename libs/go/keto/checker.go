package keto

import (
	"buybikeshop/libs/go/config"

	ory "github.com/ory/keto-client-go"
)

type Checker struct {
	ory.APIClient
}

func ProvideKetoChecker(config *config.Config) Checker {
	if config.String("keto.read") == "" {
		panic("keto.read is required")
	}
	configuration := ory.NewConfiguration()
	configuration.Servers = []ory.ServerConfiguration{
		{
			URL: config.String("keto.read"),
		},
	}
	readClient := ory.NewAPIClient(configuration)
	return Checker{APIClient: *readClient}
}
