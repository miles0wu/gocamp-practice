package main

import "unp/pkg"

func main() {
	msg := pkg.Encode("Hello from goim client!")
	pkg.Decode(msg)
}
