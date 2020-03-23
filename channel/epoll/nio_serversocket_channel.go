package epoll

import (
	"net"
	"radish/channel"
	"radish/channel/iface"
	"radish/channel/pipeline"
)

type NIOServerSocketChannel struct {
	*NIOSocketChannel
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
		fd:      int(f.Fd()),
		network: network,
		address: address,
		f:       f,
		ln:      ln,
	}

	ssChannel := &NIOServerSocketChannel{
		NIOSocketChannel: epchannel,
	}
	ssChannel.unsafe = NewMessageUnsafe(ssChannel)
	ssChannel.pipeline = pipeline.NewDefaultChannelPipeline(ssChannel)

	return ssChannel
}
