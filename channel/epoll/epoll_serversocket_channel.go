package epoll

import (
	"net"
	"radish/channel/pipeline"
)

type EpollServerSocketChannel struct {
	*EpollChannel
	address string
}

func NewEpollServerSocketChannel(address string) *EpollServerSocketChannel {

	ln, _ := net.Listen("tcp", address)
	//fd, err := unix.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	l, _ := ln.(*net.TCPListener)
	f, _ := l.File()
	epchannel := &EpollChannel{
		fd: int(f.Fd()),
	}

	ssChannel := &EpollServerSocketChannel{
		EpollChannel: epchannel,
		address:      address,
	}
	ssChannel.unsafe = NewMessageUnsafe(ssChannel)
	ssChannel.pipeline = pipeline.NewDefaultChannelPipeline(ssChannel)

	return ssChannel
}
