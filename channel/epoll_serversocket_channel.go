package channel

import (
	"net"
)

type EpollServerSocketChannel struct {
	*EpollChannel
	address string
}

func NewEpollServerSocketChannel(address string) *EpollServerSocketChannel {

	addr, _ := net.ResolveTCPAddr("tcp4", address)
	ln, _ := net.ListenTCP("tcp4", addr)
	//fd, err := unix.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	f, _ := ln.File()

	epchannel := &EpollChannel{
		fd: int(f.Fd()),
	}

	ssChannel := &EpollServerSocketChannel{
		EpollChannel: epchannel,
		address:      address,
	}

	ssChannel.unsafe = NewMessageUnsafe(ssChannel)
	ssChannel.pipeline = NewDefaultChannelPipeline(ssChannel)

	return ssChannel
}
