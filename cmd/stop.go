package cmd

import (
	"github.com/papey08/apcor/cmd/factory"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop",
	Long:  `stop`,
	Run: func(cmd *cobra.Command, args []string) {
		cliClient, err := factory.CliClient(configPath)
		handleError(err)

		handleError(cliClient.Stop(cmd.Context(), repos, group))
	},
}

func init() {
	stopCmd.Flags().StringArrayVarP(&repos, "repos", "r", []string{}, "repo")
	stopCmd.Flags().StringVarP(group, "group", "g", "", "group")
	rootCmd.AddCommand(stopCmd)
}
