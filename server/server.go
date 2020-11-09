package server

import (
	"fmt"
	"sync"

	"github.com/ddway2/go-jamulus/cli"
	"github.com/ddway2/opus"
)

// Info contains server informations
type Info struct {
}

// Server is the main application
type Server struct {
	mu   sync.Mutex
	done chan bool

	Conf    *cli.Config
	SSocket *Socket

	OPUSDec *opus.Decoder
}

// NewServer initialize server from Config
func NewServer(conf *cli.Config) (*Server, error) {
	s := &Server{
		Conf: conf,
	}
	var err error
	s.SSocket, err = NewSocket(s)
	if err != nil {
		return nil, fmt.Errorf("Unable to create socket")
	}

	return s, nil
}

// Start server
func (self *Server) Start() {

}

// Shutdown server if available
func (self *Server) Shutdown() {

}

// MixEncodeTransmitData preare data from other clients
func (self *Server) MixEncodeTransmitData( /*Channel count, num client*/ ) {

}
