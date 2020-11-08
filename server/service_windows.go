package server

import "golang.org/x/sys/windows/svc"

func (self *Server) Run() error {
	isInteractive, err := svc.IsAnInteractiveSession()
	if err != nil {
		return err
	}
	if isInteractive {
		self.Start()
		return nil
	}
	return nil
}
