package cmd

import (
	"buybikeshop/libs/go/ngutils/cmd/commands"
	"fmt"

	"github.com/spf13/cobra"
)

func NewRootCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use: "ngutils <command> <flags>",
	}

	cmd.AddCommand(commands.NewNGRXEntityCommand())
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
