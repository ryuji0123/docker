package main

import (
	"github.com/docker/client/cmd"
)

func main() {
	tcmd := cmd.NewDockerCommand()
	tcmd.Execute()
}
