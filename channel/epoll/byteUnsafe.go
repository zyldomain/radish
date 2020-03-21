package epoll

import (
	"golang.org/x/sys/unix"
	"net"
	"radish/channel/iface"
	"radish/channel/util"
	"sync"
)

var pool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 32)
	},
}

type ByteUnsafe struct {
	channel iface.Channel
}

func NewByteUnsafe(channel iface.Channel) *ByteUnsafe {
	return &ByteUnsafe{channel: channel}
}

func (b *ByteUnsafe) Read(links *util.ArrayList) {

	buf := pool.Get().([]byte)
	for {
		n, err := unix.Read(b.channel.FD(), buf)
		if err != nil || n == 0 {
			if err == unix.EAGAIN {
				return
			}
			//fmt.Println("error : " + err.Error())
			return
		}
		tmp := make([]byte, n)
		copy(tmp, buf)
		links.Add(tmp)
	}

	pool.Put(buf)
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
