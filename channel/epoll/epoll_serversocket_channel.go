package epoll

import (
	"golang.org/x/sys/unix"
	"net"
	"radish/channel"
	"radish/channel/iface"
	"radish/channel/pipeline"
	"radish/channel/util"
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
		c := NewNIOSocketChannel(ssc.network, addr.String(), nfd)

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
