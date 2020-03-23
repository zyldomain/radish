// +build linux darwin netbsd freebsd openbsd dragonfly

package epoll

import (
	"golang.org/x/sys/unix"
	"net"
	"radish/channel/util"
)

func (ssc *NIOServerSocketChannel) doReadMessages(links *util.ArrayList) {
	for {
		nfd, sa, err := unix.Accept(int(ssc.fd))
		addr := util.SockaddrToTCPOrUnixAddr(sa)
		if err != nil {
			if err == unix.EAGAIN {
				return
			} else {
				panic(err)
			}

		}
		c := NewNIOSocketChannel(nil,ssc.network, addr.String(), nfd)

		links.Add(c)
	}
}

func (ssc *NIOServerSocketChannel) write(msg interface{}) (int, error) {
	return 0, nil
}

func (ssc *NIOServerSocketChannel) bind(address string) {
	l, err := net.ResolveTCPAddr(ssc.network, address)

	if err != nil {
		panic(err)
	}
	sa := &unix.SockaddrInet4{Port: l.Port}
	copy(sa.Addr[:], l.IP)

	unix.Bind(ssc.FD(), sa)
}

func (ssc *NIOServerSocketChannel) SetNonBlocking() {
	unix.SetNonblock(ssc.fd, true)

}
