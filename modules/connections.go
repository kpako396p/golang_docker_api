package connections

import (
	"flag"
	"fmt"
	"net"
)

const (
	host string = "192.168.1.200"
	port string = "6666"
)

type Server struct {
	proto   string
	addr    string
	handler func(c *net.TCPConn) error
}

func DetectMode() string {
	arguments := flag.String("mode", "server", "server or client")
	flag.Parse()
	switch *arguments {
	case "server":
		fmt.Println("SERVER")
		s := &Server{proto: "tcp", addr: net.JoinHostPort(host, port)}
		s.ListenAndGo()
	case "client":
		fmt.Println("CLIENT")
	}
	return (*arguments)
}
