package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"net"
)

func main() {
	//go server()

	client()

	select {}
}

func server() {

	l, _ := net.ListenPacket("udp", "localhost:9000")
	ln, _ := l.(*net.UDPConn)
	f, _ := ln.File()
	for {

		buf := make([]byte, 1)
		_, sa, _ := unix.Recvfrom(int(f.Fd()), buf, 0)
		//unix.Read()

		fmt.Println(string(buf))

		fmt.Println(sa)

		unix.Sendto(int(f.Fd()), buf, 0, sa)
	}
}

func client() {
	addr, _ := net.ResolveUDPAddr("udp", "localhost:8080")
	conn, _ := net.DialUDP("udp", nil, addr)
	conn.Write([]byte("hello"))
	buf := make([]byte, 1024)
	conn.ReadFromUDP(buf)

	fmt.Println(string(buf))

}
