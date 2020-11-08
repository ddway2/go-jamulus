package server

type Socket struct {
	PServer *Server
}

func NewSocket(serv *Server) (*Socket, error) {
	s := &Socket{
		PServer: serv,
	}
	return s, nil
}
