package commands

import (
	"buybikeshop/apps/office/app/api"
	"buybikeshop/apps/office/app/pkg"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"buybikeshop/libs/go/auth"
	"buybikeshop/libs/go/config"
	"buybikeshop/libs/go/persistance"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go.uber.org/dig"
)

func NewHttpRootCommand() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "serve",
		Short: "Run office server",
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
			_ = c.Provide(auth.ProvideKratos)

			_ = pkg.Module(c)
			_ = api.Module(c)

			if err = c.Invoke(func(router *gin.Engine, config *config.Config, ctx context.Context) {
				srv := &http.Server{
					Addr:    fmt.Sprintf(":%d", config.Int("port")),
					Handler: router,
				}

				log.Printf("Run server on :%d\n", config.Int("port"))
				go func() {
					if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
						log.Fatalf("listen: %s\n", err)
					}
				}()

				quit := make(chan os.Signal, 1)
				signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
				<-quit

				log.Println("Shutting down server...")
				appCancel()

				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()

				if err = srv.Shutdown(ctx); err != nil {
					log.Fatal("Server forced to shutdown:", err)
				}

				log.Println("Server exiting")
			}); err != nil {
				panic(err)
			}
		},
	}

	return cmd
}
