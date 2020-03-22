package epoll

import (
	"golang.org/x/sys/unix"
	"net"
	"os"
	"radish/channel"
	"radish/channel/iface"
	"radish/channel/pipeline"
	"radish/channel/util"
)

type NIOServerSocketChannel struct {
	*NIOSocketChannel
	address string
	ln      net.Listener
	f       *os.File
}

const NIOServerSocket = "NIOServerSocket"

func init() {
	channel.SetChannel("NIOServerSocket", NewNIOServerSocketChannel)
}

func NewNIOServerSocketChannel(network string, address string, fd int) iface.Channel {

	ln, _ := net.Listen(network, address)
	//fd, err := unix.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	l, _ := ln.(*net.TCPListener)
	f, _ := l.File()

	epchannel := &NIOSocketChannel{
		fd: int(f.Fd()),
	}

	ssChannel := &NIOServerSocketChannel{
		NIOSocketChannel: epchannel,
		address:          address,
		ln:               ln,
		f:                f,
	}
	ssChannel.unsafe = NewMessageUnsafe(ssChannel)
	ssChannel.pipeline = pipeline.NewDefaultChannelPipeline(ssChannel)

	return ssChannel
}
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
		c := NewNIOSocketChannel("tcp", addr.String(), nfd)

		links.Add(c)
	}
}
