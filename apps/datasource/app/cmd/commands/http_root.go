package commands

import (
	"buybikeshop/apps/datasource/app/server/catalog"
	pb "buybikeshop/gen/grpc-buybikeshop-go/catalog"
	"buybikeshop/libs/go/config"
	"buybikeshop/libs/go/persistance"
	"context"
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
	"go.uber.org/dig"
	"google.golang.org/grpc"
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
			_ = c.Provide(func() *grpc.Server {
				return grpc.NewServer()
			})

			_ = c.Provide(catalog.ProvideCatalogAdapter)
			_ = c.Provide(catalog.ProvideCatalogServer)

			if err = c.Invoke(func(s *grpc.Server, config *config.Config, catalogServer *catalog.CatalogServer, ctx context.Context) {
				lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Int("server.port")))
				if err != nil {
					log.Fatalf("failed to listen: %v", err)
				}

				pb.RegisterCatalogServer(s, catalogServer)

				fmt.Printf("server listening on port: %d\n", config.Int("server.port"))
				_ = s.Serve(lis)

			}); err != nil {
				panic(err)
			}
		},
	}

	return cmd
}
