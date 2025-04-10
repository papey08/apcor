package cmd

import (
	"github.com/papey08/apcor/cmd/factory"
	"github.com/spf13/cobra"
)

var debugCMd = &cobra.Command{
	Use:   "debug",
	Short: "debug",
	Long:  `debug`,
	Run: func(cmd *cobra.Command, args []string) {
		cliClient, err := factory.CliClient(configPath)
		handleError(err)

		handleError(cliClient.Debug(cmd.Context(), repos, group))
	},
}

func init() {
	debugCMd.Flags().StringArrayVarP(&repos, "repos", "r", []string{}, "repo")
	debugCMd.Flags().StringVarP(group, "group", "g", "", "group")
	rootCmd.AddCommand(debugCMd)
}
