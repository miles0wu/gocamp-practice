package main

import (
	"flag"
	"fmt"
	"net"
	"unp/pkg"
)

var (
	host string
	port int
)

func init() {
	flag.StringVar(&host, "host", "localhost", "server ip")
	flag.IntVar(&port, "port", 9501, "server port")
	flag.Parse()
}

func main() {
	server := pkg.NewTcpServer(fmt.Sprintf("%s:%d", host, port), handleBasedOnFixLength)
	server.Listen()
}

func handleBasedOnFixLength(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)

	len, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading: ", err)
		return
	}

	ret := []byte("Message Received: ")
	ret = append(ret, buf[:len]...)
	conn.Write(ret)
}
