package commands

import (
	"buybikeshop/apps/datasource/app/server"
	"buybikeshop/libs/go/config"
	"buybikeshop/libs/go/persistance"
	"context"
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
	"go.uber.org/dig"
)

func NewHttpRootCommand() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "serve",
		Short: "Run webserver",
		Args:  cobra.ArbitraryArgs,
		Run: func(cmd *cobra.Command, args []string) {
			c := dig.New()

			configPath, err := cmd.Flags().GetString("config")
			if err != nil {
				panic(err)
			}

			appContext, appCancel := context.WithCancel(context.Background())
			defer appCancel()

			_ = c.Provide(func() context.Context { return appContext })
			_ = c.Provide(config.ProvideConfig(configPath))
			_ = c.Provide(persistance.ProvideDB)
			_ = c.Provide(server.ProvideGrpcServer)

			_ = server.InitServerModules(c)
			_ = c.Provide(server.NewRegistry)

			if err = c.Invoke(func(config *config.Config, r *server.Registry, ctx context.Context) {
				lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Int("server.port")))
				if err != nil {
					log.Fatalf("failed to listen: %v", err)
				}

				fmt.Printf("server listening on port: %d\n", config.Int("server.port"))
				err = r.Server.Serve(lis)
				if err != nil {
					log.Fatalf("failed to listen: %v", err)
				}

			}); err != nil {
				panic(err)
			}
		},
	}

	return cmd
}
