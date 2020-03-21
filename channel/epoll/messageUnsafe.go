package epoll

import (
	"golang.org/x/sys/unix"
	"net"
	"radish/channel/iface"
	"radish/channel/util"
)

var count = 0

type MessageUnsafe struct {
	channel iface.Channel
	list    *util.ArrayList
}

func NewMessageUnsafe(channel iface.Channel) *MessageUnsafe {
	return &MessageUnsafe{channel: channel, list: util.NewArrayList()}
}

func (u *MessageUnsafe) Read(links *util.ArrayList) {
	for {
		nfd, sa, err := unix.Accept(u.channel.FD())
		u.list.Add(nfd)
		if err != nil {
			if err == unix.EAGAIN {
				return
			} else {
				panic(err)
			}

		}
		c := NewEpollChannel(nfd, sa)

		links.Add(c)
	}

}

func (u *MessageUnsafe) Write(buf []byte) (int, error) {
	return 0, nil
}

func (u *MessageUnsafe) Bind(address string) {
	l, err := net.ResolveTCPAddr("tcp", address)

	if err != nil {
		panic(err)
	}
	sa := &unix.SockaddrInet4{Port: l.Port}
	copy(sa.Addr[:], l.IP)

	unix.Bind(u.channel.FD(), sa)
}
