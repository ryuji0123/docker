package cmd

import (
	"docker/client/cli"
	"github.com/spf13/cobra"
)

func NewBuildCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build [OPTIONS] PATH | URL | -",
		Short: "Send build str to host",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := cli.APIClient{}
			return client.ImageBuild()
		},
	}
	return cmd
}
