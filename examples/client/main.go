package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp4", "localhost:8080")

	if err != nil {
		panic(err)
	}
	for i := 0; i < 100; i++ {
		time.Sleep(time.Second)
		conn.Write([]byte("hello"))
		buf := make([]byte, 128)
		n, _ := conn.Read(buf)
		fmt.Println(string(buf[:n]))
	}
}
