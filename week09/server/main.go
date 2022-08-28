package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"
)

//使用了length field based frame decoder
const (
	PackageLengthBytes   = 4
	HeaderLengthBytes    = 2
	ProtocolVersionBytes = 2
	OperationBytes       = 4
	SequenceIDBytes      = 4

	HeaderLength = PackageLengthBytes + HeaderLengthBytes + ProtocolVersionBytes + OperationBytes + SequenceIDBytes
)

// Depack 解码器
func Depack(buffer []byte) []byte {
	length := len(buffer)

	data := make([]byte, 32)

	if length < HeaderLength {
		return data
	}
	messageLength := ByteToInt(buffer[:PackageLengthBytes])
	if length < HeaderLength+messageLength {
		return data
	}
	site := PackageLengthBytes
	headerLength := ByteToInt(buffer[site : site+HeaderLengthBytes])
	site += HeaderLengthBytes

	protocolVersion := ByteToInt16(buffer[site : site+ProtocolVersionBytes])
	site += ProtocolVersionBytes

	operation := ByteToInt(buffer[site : site+OperationBytes])
	site += OperationBytes

	sequenceID := ByteToInt(buffer[site : site+SequenceIDBytes])
	site += SequenceIDBytes

	fmt.Printf("packageLength: %d, headerLength: %d , protocolVersion: %d, operation: %d, sequenceID: %d \n", messageLength, headerLength, protocolVersion, operation, sequenceID)

	data = buffer[HeaderLength : HeaderLength+messageLength]
	return data
}

// ByteToInt 字节转换成整形
func ByteToInt(n []byte) int {
	var x int32
	err := binary.Read(bytes.NewBuffer(n), binary.BigEndian, &x)
	if err != nil {
		return 0
	}
	return int(x)
}

func ByteToInt16(n []byte) int {
	var x int16
	err := binary.Read(bytes.NewBuffer(n), binary.BigEndian, &x)
	if err != nil {
		return 0
	}
	return int(x)
}

func main() {
	netListen, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		fmt.Printf("Fatal error: %s", err.Error())
		os.Exit(1)
	}
	defer func(l net.Listener) { _ = l.Close() }(netListen)

	log.Println("Waiting for client ...") //启动后，等待客户端访问。
	for {
		conn, err := netListen.Accept() //监听客户端
		if err != nil {
			log.Println(conn.RemoteAddr().String(), "发生了错误：", err)
			continue
		}
		log.Println(conn.RemoteAddr().String(), "tcp connection success")
		go handleConnection(conn)
	}
}

//连接处理
func handleConnection(conn net.Conn) {
	//缓冲区，存储数据
	buffer := make([]byte, 1024)
	for {
		_, err := conn.Read(buffer)
		if err != nil {
			if err.Error() != "EOF" {
				log.Println(conn.RemoteAddr().String(), "connection error: ", err)
			}
			break
		}
		log.Println(string(Depack(buffer)))
	}
	defer func(c net.Conn) { _ = c.Close() }(conn)
}
