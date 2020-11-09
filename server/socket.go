package server

import (
	"net"
	"time"

	"github.com/ddway2/go-jamulus/protocol"
)

type Socket struct {
	server    *Server
	localIP   string
	localPort uint16
	conn      net.PacketConn
}

// NewSocket Create new internal socket
func NewSocket(serv *Server) (*Socket, error) {
	s := &Socket{
		server: serv,
	}
	return s, nil
}

// Start begin socket listening
func (s *Socket) Start() error {
	c, err := net.ListenPacket("udp", s.localIP+":"+string(s.localPort))
	if err != nil {
		return err
	}
	go s.serveUDP(c)
	return nil
}

func (s *Socket) serveUDP(c net.PacketConn) {
	defer c.Close()

	c.SetReadDeadline(time.Now().Add(time.Second))
	buf := make([]byte, protocol.MAX_SIZE_BYTES_NETW_BUF)

	for {
		n, addr, err := c.ReadFrom(buf)
		if err != nil {
			continue
		}
	}
}
