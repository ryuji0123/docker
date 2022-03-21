package main

import "github.com/docker/host/api/server"

type DaemonCli struct{}

func NewDaemonCli() *DaemonCli {
	return &DaemonCli{}
}

func (cli *DaemonCli) start() (err error) {
	s := server.New()
	s.ServeAPI()
	return nil
}
