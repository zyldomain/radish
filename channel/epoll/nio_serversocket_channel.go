package epoll

import (
	"errors"
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

func NewNIOServerSocketChannel(conn interface{}, network string, address string, fd interface{}) iface.Channel {

	var c net.Conn
	var ok bool
	if conn != nil {
		c, ok = conn.(net.Conn)

		if !ok {
			panic(errors.New("wrong type"))
		}
	}
	ln, err := net.Listen(network, address)

	if err != nil {
		panic(err)
	}
	//fd, err := unix.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	l, ok := ln.(*net.TCPListener)
	if !ok {
		panic(errors.New("network error"))
	}

	f, err := l.File()
	if err != nil {
		panic(err)
	}

	epchannel := &NIOSocketChannel{
		FDE:     &FDE{fd: GetFD(f.Fd())},
		network: network,
		address: address,
		f:       f,
		ln:      ln,
		msg:     make(chan *iface.Pkg, 1000),
		conn:    c,
	}

	ssChannel := &NIOServerSocketChannel{
		NIOSocketChannel: epchannel,
	}
	ssChannel.unsafe = NewMessageUnsafe(ssChannel)
	ssChannel.pipeline = pipeline.NewDefaultChannelPipeline(ssChannel)

	return ssChannel
}
