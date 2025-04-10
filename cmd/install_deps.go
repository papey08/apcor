package cmd

import (
	"github.com/papey08/apcor/cmd/factory"
	"github.com/spf13/cobra"
)

var (
	platform string

	installDepsCmd = &cobra.Command{
		Use:   "install-deps [platform]",
		Short: "install dependencies",
		Long:  `install dependencies`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			manager, err := factory.Manager(configPath)
			handleError(err)
			if len(args) > 0 {
				manager.RunInstallDeps(cmd.Context(), args[0])
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(installDepsCmd)
}
