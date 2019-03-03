package connections

import (
	"log"
	"net"
)

func (s *Server) ListenAndGo() error {
	tcpaddr, _ := net.ResolveTCPAddr(s.proto, s.addr)
	ln, err := net.ListenTCP(s.proto, tcpaddr)
	if err != nil {
		log.Println("Failed to listen for tcp connections on address ", s.addr, " Error: ", err)
		return err
	}

	for {
		conn, err := ln.AcceptTCP()
		if err != nil {
			log.Println("Failed to accept connection ", conn, " due to error ", err)
			continue
		}
		log.Println("Client ", conn.RemoteAddr(), " connected")
		go s.handler(conn)
	}
	return nil
}
