package test

import (
	"fmt"
	"net"
)

func RadishGo(writebuf []byte, address string) {
	conn, err := net.Dial("tcp4", address)

	if err != nil {
		panic(err)
	}
	conn.Write(writebuf)
	buf := make([]byte, 128)
	n, _ := conn.Read(buf)
	fmt.Println(string(buf[:n]))
	conn.Close()

}
