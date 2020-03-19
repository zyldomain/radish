package channel

import (
	"golang.org/x/sys/unix"
	"net"
	"radish/channel/util"
)

type ByteUnsafe struct {
	channel Channel
}

func NewByteUnsafe(channel Channel) *ByteUnsafe {
	return &ByteUnsafe{channel: channel}
}

func (b *ByteUnsafe) Read(links *util.ArrayList) {
	for {
		buf := make([]byte, 1024*1024)
		_, err := unix.Read(b.channel.FD(), buf)
		if err != nil {
			if err == unix.EAGAIN {
				return
			} else {
				//panic(err)
			}
		}

		links.Add(buf)
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
