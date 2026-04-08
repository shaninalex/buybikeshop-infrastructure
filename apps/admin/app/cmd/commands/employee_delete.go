package commands

import "github.com/spf13/cobra"

func NewEmployeeDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete [id]",
		Short: "Delete employee",
		Long:  "Require 1 arguments - employee id. For example:\nbuybikeshop-admin employee delete be2bb263-2f7a-4194-896b-b8a06b20e707",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	return cmd
}
