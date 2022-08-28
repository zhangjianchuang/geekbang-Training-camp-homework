package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

//使用了length field based frame decoder
const (
	ProtocolVersionBytes = 2
	OperationBytes       = 4
	SequenceIDBytes      = 4
)

// Enpack 封包
func Enpack(message []byte) []byte {
	b := append(Int32ToBytes(len(message)), Int16ToBytes(0)...)
	b = append(b, Int16ToBytes(8)...)
	b = append(b, Int32ToBytes(99)...)
	b = append(b, Int32ToBytes(10)...)
	b = append(b, message...)

	return b
}

// Int32ToBytes 整数转换成字节
func Int32ToBytes(n int) []byte {
	buf := bytes.NewBuffer([]byte{})
	_ = binary.Write(buf, binary.BigEndian, int32(n))
	return buf.Bytes()
}

// Int16ToBytes 整数转换成字节
func Int16ToBytes(n int) []byte {
	buf := bytes.NewBuffer([]byte{})
	_ = binary.Write(buf, binary.BigEndian, int16(n))
	return buf.Bytes()
}

func send(conn net.Conn, msg string) {
	_, _ = conn.Write(Enpack([]byte(msg)))
	fmt.Println("send success:", msg)
	defer func(c net.Conn) { _ = c.Close() }(conn)
}

func main() {
	connectAndSend()
}

func connectAndSend() {
	server := "localhost:8081"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Printf("Fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Printf("Fatal error: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println("connect success, sending...")
	msg := "{\"Session\":\"" + strconv.FormatInt(time.Now().Unix(), 10) + "\",\"Content\":\"message\"}"
	send(conn, msg)
}
