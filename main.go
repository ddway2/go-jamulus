package main

import (
	"github.com/ddway2/go-jamulus/cli"
	"github.com/ddway2/go-jamulus/server"
)

func main() {
	var conf cli.Config

	s, err := server.NewServer(&conf)
	if err != nil {
		server.LogDie("%s", err)
	}

	if err := s.Run(); err != nil {
		server.LogDie("%s", err)
	}
}
