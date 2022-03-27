package main

import (
	"flag"
	"fmt"
	"net"
	"strings"
	"unp/pkg"
)

var (
	host string
	port int
)

const delimiter = "\t"

func init() {
	flag.StringVar(&host, "host", "localhost", "server ip")
	flag.IntVar(&port, "port", 9501, "server port")
	flag.Parse()
}

func main() {
	server := pkg.NewTcpServer(fmt.Sprintf("%s:%d", host, port), handleBasedOnDelimiter)
	server.Listen()
}

func handleBasedOnDelimiter(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	message := ""

	for {
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return
		}
		message += string(buf[:len])
		if strings.Contains(string(buf[:len]), delimiter) {
			break
		}
	}

	conn.Write([]byte("Message Received: " + message))
}
