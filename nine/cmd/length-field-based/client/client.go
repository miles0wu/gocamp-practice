package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"unp/pkg"
)

var (
	host string
	port int
	msg  string
)

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

	msg = string(encode(msg))
	recv, err := c.SendMessage(msg)
	if err != nil {
		fmt.Println("Send message failed: ", err)
		return
	}
	fmt.Println(recv)
}

func encode(body string) []byte {
	headerLen := 4
	bodyLen := len(body)
	ret := make([]byte, headerLen+bodyLen)

	binary.BigEndian.PutUint32(ret, uint32(bodyLen))
	copy(ret[headerLen:], body)

	return ret
}
