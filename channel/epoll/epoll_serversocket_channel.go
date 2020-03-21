package epoll

import (
	"net"
	"os"
	"radish/channel"
	"radish/channel/iface"
	"radish/channel/pipeline"
)

type EpollServerSocketChannel struct {
	*EpollChannel
	address string
	ln      net.Listener
	f       *os.File
}

const EpollServerSocket = "EpollServerSocketChannel"

func init() {
	channel.SetChannel("EpollServerSocketChannel", NewEpollServerSocketChannel)
}

func NewEpollServerSocketChannel(address string) iface.Channel {

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
		ln:           ln,
		f:            f,
	}
	ssChannel.unsafe = NewMessageUnsafe(ssChannel)
	ssChannel.pipeline = pipeline.NewDefaultChannelPipeline(ssChannel)

	return ssChannel
}
