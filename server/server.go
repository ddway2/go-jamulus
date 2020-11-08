package server

import (
	"fmt"

	"github.com/ddway2/go-jamulus/cli"
	"gopkg.in/hraban/opus.v2"
)

type Server struct {
	Conf    *cli.Config
	SSocket *Socket

	OPUSDec *opus.Decoder
}

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

// Server function
func (self *Server) Configure() {

}

func (self *Server) Start() {

}

func (self *Server) Shutdown() {

}

func (self *Server) MixEncodeTransmitData( /*Channel count, num client*/ ) {

}

func (self *Server) AcceptLoop(clr chan struct{}) {
	// Clean chan connection
	defer func() {
		if clr != nil {
			close(clr)
		}
	}()

}
