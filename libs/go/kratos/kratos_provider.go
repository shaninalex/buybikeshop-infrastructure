package kratos

import (
	"buybikeshop/libs/go/config"

	ory "github.com/ory/kratos-client-go"
)

func ProvideKratos(config *config.Config) *ory.APIClient {
	configuration := ory.NewConfiguration()
	configuration.Servers = []ory.ServerConfiguration{
		{
			URL: config.String("kratos.url_browser"),
		},
	}
	return ory.NewAPIClient(configuration)
}
