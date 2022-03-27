package main

import (
	"flag"
	"fmt"
	"unp/pkg"
)

var (
	host string
	port int
	msg  string
)

const delimiter = "\t"

func init() {
	flag.StringVar(&host, "host", "localhost", "server ip")
	flag.IntVar(&port, "port", 9501, "server port")
	flag.StringVar(&msg, "msg", "Hi, there.", "message")
	flag.Parse()
}

func main() {
	c, err := pkg.NewTcpClient(fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		panic(err)
	}

	msg += delimiter
	recv, err := c.SendMessage(msg)
	if err != nil {
		fmt.Println("Send message failed: ", err)
		return
	}
	fmt.Println(recv)
}
