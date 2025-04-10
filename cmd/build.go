package cmd

import (
	"github.com/papey08/apcor/cmd/factory"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "build",
	Long:  `build`,
	Run: func(cmd *cobra.Command, args []string) {
		cliClient, err := factory.CliClient(configPath)
		handleError(err)

		handleError(cliClient.Build(cmd.Context(), repos, group))
	},
}

func init() {
	buildCmd.Flags().StringArrayVarP(&repos, "repos", "r", []string{}, "repo")
	buildCmd.Flags().StringVarP(group, "group", "g", "", "group")
	rootCmd.AddCommand(buildCmd)
}
