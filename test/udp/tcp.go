package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 0 ; i < 10 ; i++{
		conn, err := net.Dial("tcp4", "localhost:8080")

		if err != nil {
			panic(err)
		}
		conn.Write([]byte("hello"))
		buf := make([]byte, 128)
		n, _ := conn.Read(buf)
		fmt.Println(string(buf[:n]))
		conn.Close()
	}
}
