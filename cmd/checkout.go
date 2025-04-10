package cmd

import (
	"github.com/papey08/apcor/cmd/factory"
	"github.com/spf13/cobra"
)

var (
	checkoutCmd = &cobra.Command{
		Use:   "checkout",
		Short: "checkout",
		Long:  `checkout`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			cliClient, err := factory.CliClient(configPath)
			handleError(err)
			if len(args) > 0 {
				handleError(cliClient.Checkout(cmd.Context(), repos, group, args[0]))
			}
		},
	}
)

func init() {
	checkoutCmd.Flags().StringArrayVarP(&repos, "repos", "r", []string{}, "repo")
	checkoutCmd.Flags().StringVarP(group, "group", "g", "", "group")
	rootCmd.AddCommand(checkoutCmd)
}
