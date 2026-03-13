package cmd

import (
	"buybikeshop/apps/datasource/app/cmd/commands"

	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "datasource",
	}

	cmd.AddCommand(commands.NewHttpRootCommand())

	cmd.PersistentFlags().String("config", "", "Configuration path. Required.")
	_ = cmd.MarkPersistentFlagRequired("config")
	return cmd
}

func Execute() int {
	return 0
}
