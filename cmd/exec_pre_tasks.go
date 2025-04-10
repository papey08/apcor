package cmd

import (
	"github.com/papey08/apcor/cmd/factory"
	"github.com/spf13/cobra"
)

var execPreTasksCmd = &cobra.Command{
	Use:   "exec-pre-tasks",
	Short: "exec pre tasks",
	Long:  `exec pre tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		cliClient, err := factory.CliClient(configPath)
		handleError(err)

		handleError(cliClient.ExecPreTasks(cmd.Context()))
	},
}

func init() {
	rootCmd.AddCommand(execPreTasksCmd)
}
