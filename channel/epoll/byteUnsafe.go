package epoll

import (
	"fmt"
	"golang.org/x/sys/unix"
	"net"
	"radish/channel/iface"
	"radish/channel/util"
)

type ByteUnsafe struct {
	channel iface.Channel
}

func NewByteUnsafe(channel iface.Channel) *ByteUnsafe {
	return &ByteUnsafe{channel: channel}
}

func (b *ByteUnsafe) Read(links *util.ArrayList) {
	for {
		buf := make([]byte, 1024)
		n, err := unix.Read(b.channel.FD(), buf)
		if err != nil || n == 0 {
			if err == unix.EAGAIN {
				fmt.Println("unix.EAGAIN", links.Size())
				return
			}
			fmt.Println("error : " + err.Error())
			unix.Close(b.channel.FD())
		}
		links.Add(buf[:n])
	}

}

func (b *ByteUnsafe) Write(buf []byte) (int, error) {
	return unix.Write(b.channel.FD(), buf)
}

func (u *ByteUnsafe) Bind(address string) {
	l, err := net.ResolveTCPAddr("tcp", address)

	if err != nil {
		panic(err)
	}
	sa := &unix.SockaddrInet4{Port: l.Port}
	copy(sa.Addr[:], l.IP)

	unix.Bind(u.channel.FD(), sa)
}
