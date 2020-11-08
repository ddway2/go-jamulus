package server

import (
	"fmt"
	"os"

	"github.com/ddway2/go-jamulus/cli"
)

type Server struct{}

func NewServer(conf *cli.Config) (*Server, error) {
	s := &Server{}
	return s, nil
}

func LogDie(msg string, v ...interface{}) {
	fmt.Fprintln(os.Stderr, msg, v)
	os.Exit(1)
}

func LogExit(msg string, v ...interface{}) {
	fmt.Fprintln(os.Stderr, msg, v)
	os.Exit(0)
}

// Server function
func (self *Server) Configure() error {
	return nil
}

func (self *Server) Run() error {
	return nil
}
