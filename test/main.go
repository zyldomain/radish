package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"net"
	"radish/channel"
)

func main() {

	addr, _ := net.ResolveTCPAddr("tcp4", "localhost:8080")
	ln, _ := net.ListenTCP("tcp4", addr)
	selector, _ := channel.OpenEpollSelector()
	f, _ := ln.File()
	ch := channel.NewEpollChannel(int(f.Fd()), nil)
	selector.AddInterests(ch, unix.EVFILT_READ)

	for {
		keys, _ := selector.Select()

		fmt.Println(keys)
	}
}
