package main

import (
	"encoding/binary"
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
	server := pkg.NewTcpServer(fmt.Sprintf("%s:%d", host, port), handleBasedOnLengthField)
	server.Listen()
}

// 協議設計： 前4byte為訊息長度，訊息最長可達2^32-1
func handleBasedOnLengthField(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	len, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading", err.Error())
		return
	}
	if len < 4 {
		fmt.Println("Fail to decode...")
		return
	}

	size := binary.BigEndian.Uint32(buf[:4])
	message := buf[4 : 4+size]

	conn.Write(append([]byte("Message Received: "), message...))
}
