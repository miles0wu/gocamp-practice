package pkg

import (
	"encoding/binary"
	"fmt"
)

/*
	Go im protocol design:
	1. Package Length: 4 bytes
	2. Header Length: 2 bytes
	3. Protocol Version: 2 bytes
	4. Operation: 4 bytes
	5. Sequence ID: 4 bytes
	6. Body: Package length - Header length
*/
func Decode(data []byte) {
	if len(data) <= 16 {
		fmt.Println("decode failed: input less then 16 bytes")
	}

	packageLen := binary.BigEndian.Uint32(data[:4])
	headerLen := binary.BigEndian.Uint16(data[4:6])
	protocolVersion := binary.BigEndian.Uint16(data[6:8])
	operation := binary.BigEndian.Uint32(data[8:12])
	sequenceId := binary.BigEndian.Uint32(data[12:16])

	fmt.Printf("packageLen: %v, headerLen: %v, version: %v, operation: %v, seqId: %v\n", packageLen, headerLen, protocolVersion, operation, sequenceId)
	fmt.Printf("body: %v\n", string(data[16:]))
}

func Encode(msg string) []byte {
	bodyLen := len(msg)
	data := make([]byte, bodyLen+16)

	binary.BigEndian.PutUint32(data[:4], uint32(cap(data)))
	binary.BigEndian.PutUint16(data[4:6], uint16(16))
	binary.BigEndian.PutUint16(data[6:8], uint16(1))
	binary.BigEndian.PutUint32(data[8:12], uint32(2))
	binary.BigEndian.PutUint32(data[12:16], uint32(3))
	copy(data[16:], msg)

	return data
}
