package main

import (
	"docker/client/cmd"
)

func main() {
	tcmd := cmd.NewDockerCommand()
	tcmd.Execute()
}
