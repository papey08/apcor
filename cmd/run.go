package cmd

import (
	"github.com/papey08/apcor/cmd/factory"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run",
	Long:  `run`,
	Run: func(cmd *cobra.Command, args []string) {
		cliClient, err := factory.CliClient(configPath)
		handleError(err)

		handleError(cliClient.Run(cmd.Context(), repos, group))
	},
}

func init() {
	runCmd.Flags().StringArrayVarP(&repos, "repos", "r", []string{}, "repo")
	runCmd.Flags().StringVarP(group, "group", "g", "", "group")
	rootCmd.AddCommand(runCmd)
}
