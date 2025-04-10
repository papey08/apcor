package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "apcor",
		Short: "apcor",
		Long:  `apcor`,
	}

	cliServerAddr string
	configPath    string

	group *string
	repos []string
)

func handleError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func Execute() {
	rootCmd.PersistentFlags().StringVarP(&cliServerAddr, "addr", "a", "localhost:16358", "cli server address")
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "./apcor.yaml", "config path")
	handleError(rootCmd.Execute())
}
