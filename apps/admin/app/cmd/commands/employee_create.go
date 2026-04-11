package commands

import (
	"buybikeshop/apps/admin/app/services"
	"buybikeshop/libs/go/config"
	"buybikeshop/libs/go/kratos"
	"context"

	"github.com/spf13/cobra"
	"go.uber.org/dig"
)

func NewEmployeeCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [file_path]",
		Short: "Create employee",
		Long:  "Require 1 arguments - path to json file with new employee description. For example:\nbuybikeshop-admin employee create ./resources/employee_a.json",
		Args:  cobra.MinimumNArgs(1),
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
			_ = c.Provide(kratos.ProvideKratos)

			_ = services.Module(c)

		},
	}
	return cmd
}
