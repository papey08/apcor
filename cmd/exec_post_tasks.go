package cmd

import (
	"github.com/papey08/apcor/cmd/factory"
	"github.com/spf13/cobra"
)

var execPostTasksCmd = &cobra.Command{
	Use:   "exec-post-tasks",
	Short: "exec post tasks",
	Long:  `exec post tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		cliClient, err := factory.CliClient(configPath)
		handleError(err)
		if err := cliClient.ExecPostTasks(cmd.Context()); err != nil {
			handleError(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(execPostTasksCmd)
}
