package pkg

import "net"

type tcpClient struct {
	addr string
	conn net.Conn
}

func NewTcpClient(addr string) (*tcpClient, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &tcpClient{addr, conn}, nil
}

func (c *tcpClient) SendMessage(msg string) (string, error) {
	send := []byte(msg)
	c.conn.Write(send)

	recv := make([]byte, 1024)
	len, err := c.conn.Read(recv)
	if err != nil {
		return "", err
	}

	return string(recv[:len]), nil
}
