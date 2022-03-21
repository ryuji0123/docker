package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

func newDaemonCommand() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "dockerd [OPTIONS]",
		Short: "A self-sufficient runtime for containers",
		RunE: func(cmd *cobra.Command, args []string) error {
			daemonCli := NewDaemonCli()
			return daemonCli.start()
		},
	}
	return cmd, nil
}

func init() {
	logrus.SetOutput(os.Stdout)
}

func main() {
	onError := func(err error) {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	cmd, err := newDaemonCommand()
	if err != nil {
		onError(err)
	}
	if err := cmd.Execute(); err != nil {
		onError(err)
	}
}
