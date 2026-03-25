package cmd

import (
	"buybikeshop/apps/datasource/app/cmd/commands"
	"fmt"

	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "buybikeshop-datasource",
	}

	cmd.AddCommand(commands.NewHttpRootCommand())

	cmd.PersistentFlags().String("config", "", "Configuration path. Required.")
	_ = cmd.MarkPersistentFlagRequired("config")
	return cmd
}

func Execute() int {
	c := NewRootCmd()
	if err := c.Execute(); err != nil {
		fmt.Println(err)
		return 1
	}
	return 0
}
