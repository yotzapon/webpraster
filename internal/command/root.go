package command

import (
	"fmt"
	"github.com/spf13/cobra"
)

func Execute() {
	rootCmd := &cobra.Command{
		Use:   "cli",
		Short: "CLI for running tasks",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
	configureWebpRasterCommand(rootCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
