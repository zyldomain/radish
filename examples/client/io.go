package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	ln, err := net.Listen("tcp4", "localhost:8081")

	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()

		if err != nil {
			panic(err)
		}

		go Loop(conn)

	}
}

func Loop(conn net.Conn) {
	for {
		buf := make([]byte, 128)
		n, err := conn.Read(buf)

		if err == io.EOF {
			return
		}

		fmt.Println("客户端发送消息-> " + string(buf[:n]))
		conn.Write([]byte("服务端发送消息-> " + string(buf[:n])))
	}

}
