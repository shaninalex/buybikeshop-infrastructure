package cmd

import (
	"fmt"

	"buybikeshop/apps/warehouse/app/cmd/commands"

	"github.com/spf13/cobra"
)

func NewRootCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use: "buybikeshop-warehouse",
	}

	cmd.AddCommand(commands.NewHttpRootCommand())

	cmd.PersistentFlags().String("config", "", "Configuration path. Required.")
	_ = cmd.MarkPersistentFlagRequired("config")
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
