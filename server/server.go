package server

import (
	"github.com/ddway2/go-jamulus/cli"
)

type Server struct{}

func NewServer(conf *cli.Config) (*Server, error) {
	s := &Server{}
	return s, nil
}

// Server function
func (self *Server) Configure() {

}

func (self *Server) Start() error {
	return nil
}

func (self *Server) Shutdown() {

}

func (self *Server) AcceptLoop(clr chan struct{}) {
	// Clean chan connection
	defer func() {
		if clr != nil {
			close(clr)
		}
	}()

}
