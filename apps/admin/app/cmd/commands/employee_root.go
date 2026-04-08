package commands

import (
	"github.com/spf13/cobra"
)

func NewEmployeeRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "identities",
		Short: "Manage identities",
	}

	cmd.AddCommand(NewEmployeeCreateCmd())
	cmd.AddCommand(NewEmployeeDeleteCmd())
	cmd.AddCommand(NewEmployeeListCmd())

	return cmd
}
