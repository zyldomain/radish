// +build linux darwin netbsd freebsd openbsd dragonfly

package epoll

import (
	"errors"
	"golang.org/x/sys/unix"
	"net"
	"radish/channel/util"
)

func (ec *NIOSocketChannel) doReadMessages(links *util.ArrayList) {
	buf := pool.Get().([]byte)
	for {
		//buf := make([]byte, 4096)
		n, err := unix.Read(ec.fd, buf)
		if err != nil || n == 0 || buf[n-1] == 4 {
			if err == unix.EAGAIN {
				return
			}
			ec.Close()
			return
		}

		links.Add(buf[:n])
	}

	pool.Put(buf)
}

func (ec *NIOSocketChannel) write(msg interface{}) (int, error) {

	buf, ok := msg.([]byte)

	if !ok {
		panic(errors.New("wrong type"))
	}
	return unix.Write(ec.FD(), buf)
}

func (ec *NIOSocketChannel) bind(address string) {
	l, err := net.ResolveTCPAddr("tcp", address)

	if err != nil {
		panic(err)
	}
	sa := &unix.SockaddrInet4{Port: l.Port}
	copy(sa.Addr[:], l.IP)

	unix.Bind(ec.FD(), sa)
}

func (ec *NIOSocketChannel) SetNonBolcking() {
	unix.SetNonblock(ec.fd, true)
}

func (ec *NIOSocketChannel) close() {

	ec.active = false
	unix.Close(ec.fd)
	ec.pipeline.ChannelInActive(ec)
}
