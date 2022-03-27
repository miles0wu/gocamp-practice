package pkg

import (
	"fmt"
	"net"
)

type tcpServer struct {
	listenAddr string
	handler    func(conn net.Conn)
}

func NewTcpServer(listenAddr string, handler func(conn net.Conn)) *tcpServer {
	return &tcpServer{listenAddr: listenAddr, handler: handler}
}

func (s tcpServer) Listen() error {
	l, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("receive error connection: ", err)
			continue
		}

		go s.handler(conn)
	}
}
