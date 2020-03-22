package epoll

import (
	"golang.org/x/sys/unix"
	"net"
	"radish/channel/iface"
	"radish/channel/util"
)

type MessageUnsafe struct {
	channel iface.Channel
}

func NewMessageUnsafe(channel iface.Channel) *MessageUnsafe {
	return &MessageUnsafe{channel: channel}
}

func (u *MessageUnsafe) Read(links *util.ArrayList) {
	c, _ := u.channel.(AbstractChannel)
	c.doReadMessages(links)

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
