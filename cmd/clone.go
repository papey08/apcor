package cmd

import (
	"github.com/papey08/apcor/cmd/factory"
	"github.com/spf13/cobra"
)

var (
	cloneCmd = &cobra.Command{
		Use:   "clone",
		Short: "clone",
		Long:  `clone`,
		Run: func(cmd *cobra.Command, args []string) {
			cliClient, err := factory.CliClient(configPath)
			handleError(err)

			handleError(cliClient.Clone(cmd.Context(), args, group))
		},
	}
)

func init() {
	cloneCmd.Flags().StringArrayVarP(&repos, "repos", "r", []string{}, "repo")
	cloneCmd.Flags().StringVarP(group, "group", "g", "", "group")
	rootCmd.AddCommand(cloneCmd)
}
