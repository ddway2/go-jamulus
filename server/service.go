// +build !windows

package server

func (self *Server) Run() error {
	self.Start()
	return nil
}
