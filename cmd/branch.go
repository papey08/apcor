package cmd

import (
	"fmt"

	"github.com/papey08/apcor/cmd/factory"
	"github.com/spf13/cobra"
)

var (
	branchCmd = &cobra.Command{
		Use:   "branch",
		Short: "branch",
		Long:  `branch`,
		Run: func(cmd *cobra.Command, args []string) {
			cliClient, err := factory.CliClient(configPath)
			handleError(err)

			branches, err := cliClient.Branch(cmd.Context())
			handleError(err)

			for name, branch := range branches {
				fmt.Printf("%s: %s\n", name, branch)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(branchCmd)
}
