package commands

import (
	"buybikeshop/apps/datasource/app/pkg"
	pb "buybikeshop/gen/grpc-buybikeshop-go/catalog"
	"buybikeshop/libs/go/config"
	"buybikeshop/libs/go/persistance"
	"context"
	"fmt"
	"log"
	"net"

	"github.com/gin-gonic/gin"
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

			if err := c.Invoke(func(router *gin.Engine, config *config.Config, ctx context.Context) {

				lis, err := net.Listen("tcp", fmt.Sprintf(":%d", int(config.Datasource.GrpcPort)))
				if err != nil {
					log.Fatalf("failed to listen: %v", err)
				}

				s := grpc.NewServer()
				pb.RegisterCatalogServer(s, &pkg.CatalogServer{})

				s.Serve(lis)
			}); err != nil {
				panic(err)
			}
		},
	}

	return cmd
}
