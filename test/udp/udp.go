package main

import (
	"fmt"
	"net"
)

func main() {
	//go server()

	client()

	select {}
}



func client() {
	addr, _ := net.ResolveUDPAddr("udp", "localhost:8080")
	conn, _ := net.DialUDP("udp", nil, addr)
	conn.Write([]byte("hello"))
	buf := make([]byte, 1024)
	conn.ReadFromUDP(buf)

	fmt.Println(string(buf))

}
