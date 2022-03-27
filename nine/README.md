# Week 9

### 1. To summarize several ways of unpacking socket sticky packets: fix length/delimiter based/length field based frame decoder. try to give examples.

* fixed length
```
$ make run-fixed-length-server
$ make run-fixed-length-client msg="aloha"
```

* delimiter based
```
$ make run-delimiter-based-server
$ make run-delimiter-based-client msg="aloha"
```

* length field based
```
$ make run-length-field-based-server
$ make run-length-field-based-client msg="aloha"
```

### 2. Implement a decoder that decodes the goim protocol from a socket connection

```
$ make run-goim-example
go run cmd/goim/main.go
packageLen: 39, headerLen: 16, version: 1, operation: 2, seqId: 3
body: Hello from goim client!
```
