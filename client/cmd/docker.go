package cmd

import "github.com/spf13/cobra"

func NewDockerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "docker [OPTIONS] COMMAND [ARG...]",
		Short: "A self-sufficient runtime for containers",
	}

	cmd.AddCommand(
		NewBuildCommand(),
	)
	return cmd
}
