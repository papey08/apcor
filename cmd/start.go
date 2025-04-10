package cmd

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/papey08/apcor/cmd/factory"
	"github.com/papey08/apcor/internal/cli"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start",
	Long:  `start`,
	Run: func(_ *cobra.Command, _ []string) {
		manager, err := factory.Manager(configPath)
		handleError(err)

		cliServer := cli.NewServer(manager)

		lis, err := net.Listen("tcp", cliServerAddr)
		handleError(err)

		// preparing graceful shutdown
		osSignals := make(chan os.Signal, 1)
		signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)
		signal.Notify(osSignals, os.Interrupt, syscall.SIGINT)

		go func() {
			<-osSignals
			lis.Close()
			manager.Stop(context.Background(), []string{})
			close(manager.ErrStream.ErrChan)
			cliServer.GracefulStop()
		}()

		go func() {
			for err := range manager.ErrStream.ErrChan {
				fmt.Println(err.Error())
			}
		}()

		err = cliServer.Serve(lis)
		handleError(err)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
