package cmd

import (
	"fmt"

	"buybikeshop/apps/admin/app/cmd/commands"

	"github.com/spf13/cobra"
)

func NewRootCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use: "buybikeshop-admin",
	}

	cmd.AddCommand(commands.NewHttpRootCommand())
	cmd.AddCommand(commands.NewEmployeeRootCmd())

	cmd.PersistentFlags().String("config", "", "Configuration path. Required.")
	_ = cmd.MarkPersistentFlagRequired("config")

	cmd.PersistentFlags().Bool("dev", false, "Is development mode. If true - used mock datasource")
	return cmd
}

// Execute run application
func Execute() int {
	c := NewRootCmd()

	if err := c.Execute(); err != nil {
		fmt.Println(err)
		return 1
	}
	return 0
}
