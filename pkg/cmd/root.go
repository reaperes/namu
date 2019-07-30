package cmd

import (
	"github.com/spf13/cobra"
)

func Root() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "namu",
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	return cmd
}
