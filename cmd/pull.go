package cmd

import (
	"github.com/papey08/apcor/cmd/factory"
	"github.com/spf13/cobra"
)

var (
	pullCmd = &cobra.Command{
		Use:   "push",
		Short: "push",
		Long:  `push`,
		Run: func(cmd *cobra.Command, args []string) {
			cliClient, err := factory.CliClient(configPath)
			handleError(err)

			if len(args) > 0 {
				handleError(cliClient.Pull(cmd.Context(), repos, group))
			}
		},
	}
)

func init() {
	pullCmd.Flags().StringArrayVarP(&repos, "repos", "r", []string{}, "repo")
	pullCmd.Flags().StringVarP(group, "group", "g", "", "group")
	rootCmd.AddCommand(pullCmd)
}
