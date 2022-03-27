host ?= "127.0.0.1"
port ?= 9501
msg =

run-fixed-length-server:
	go run cmd/fixed-length/server/server.go -host=$(host) -port=$(port)

run-fixed-length-client:
	go run cmd/fixed-length/client/client.go -host=$(host) -port=$(port) -msg=$(msg)

run-delimiter-based-server:
	go run cmd/fixed-length/server/server.go -host=$(host) -port=$(port)

run-delimiter-based-client:
	go run cmd/fixed-length/client/client.go -host=$(host) -port=$(port) -msg=$(msg)

run-length-field-based-server:
	go run cmd/length-field-based/server/server.go -host=$(host) -port=$(port)

run-length-field-based-client:
	go run cmd/length-field-based/client/client.go -host=$(host) -port=$(port) -msg=$(msg)

run-goim-example:
	go run cmd/goim/main.go