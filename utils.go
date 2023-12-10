package main

import (
	"bytes"
	"encoding/binary"
	"log"
)

var address = [20]string{"10.1.0.91:8051", "10.1.0.92:8052", "10.1.0.93:8053", "10.1.0.94:8054", "10.1.0.95:8055", "10.1.0.96:8056", "10.1.0.98:8057", "10.1.0.99:8058", "10.1.0.116:8059", "10.1.0.102:8060", "10.1.0.103:8061", "10.1.0.104:8062", "10.1.0.105:8063", "10.1.0.107:8064", "10.1.0.109:8065", "10.1.0.110:8066", "10.1.0.111:8067", "10.1.0.112:8068", "10.1.0.113:8069", "10.1.0.115:8070"}

var address_list = []string{address[5], address[8]}

func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

func IntToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return int(x)
}

func padBytes(input []byte, size int) []byte {
	padding := make([]byte, size-len(input))
	return append(padding, input...)
}
