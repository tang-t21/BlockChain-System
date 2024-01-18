package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
)

const leading_zeros = 22
const fake_leading_zeros = leading_zeros - 1
const total_bits = 256

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

type StringList []string

func (s *StringList) String() string {
	return fmt.Sprintf("%v", *s)
}
func (s *StringList) Set(value string) error {
	*s = append(*s, value)
	return nil
}

func (sl *StringList) Find(str string) int {
	for idx, s := range *sl {
		if s == str {
			return idx
		}
	}
	return -1
}
