package keto

import (
	"buybikeshop/libs/go/config"

	ory "github.com/ory/keto-client-go"
)

type KetoChecker struct {
	ory.APIClient
}

func ProvideKetoChecker(config *config.Config) KetoChecker {
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
	return KetoChecker{APIClient: *readClient}
}
