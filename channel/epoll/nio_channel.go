package epoll

import (
	"errors"
	"net"
	"os"
	"radish/channel"
	"radish/channel/iface"
	"radish/channel/pipeline"
)

type NIOSocketChannel struct {
	fd        int
	pipeline  iface.Pipeline
	address   string
	unsafe    iface.Unsafe
	active    bool
	eventloop iface.EventLoop
	network   string
	f         *os.File
	conn      net.Conn
	ln        net.Listener
}

const NIOSocket = "NIOSocket"

func init() {
	channel.SetChannel("NIOSocket", NewNIOSocketChannel)
}

func NewNIOSocketChannel(network string, address string, fd int) iface.Channel {

	epchannel := &NIOSocketChannel{
		fd:      fd,
		network: network,
		address: address,
	}

	epchannel.unsafe = NewByteUnsafe(epchannel)
	epchannel.pipeline = pipeline.NewDefaultChannelPipeline(epchannel)

	if fd == -1 {

		conn, err := net.Dial(network, address)
		if err != nil {
			panic(err)
		}

		tc, ok := conn.(*net.TCPConn)

		if !ok {
			panic(errors.New("network error"))
		}

		f, err := tc.File()

		if err != nil {
			panic(err)
		}

		epchannel.f = f
		epchannel.fd = int(f.Fd())

		epchannel.conn = conn
	}

	return epchannel
}

func (ec *NIOSocketChannel) FD() int {
	return ec.fd
}

func (ec *NIOSocketChannel) Read(msg interface{}) {
	ec.pipeline.ChannelRead(msg)
}

func (ec *NIOSocketChannel) Write(msg interface{}) {
	ec.pipeline.Write(msg)
}

func (ec *NIOSocketChannel) ChannelActive(msg interface{}) {
	ec.pipeline.ChannelActive(msg)
}
func (ec *NIOSocketChannel) ChannelInActive(msg interface{}) {
	ec.pipeline.ChannelInActive(msg)
}

func (ec *NIOSocketChannel) Unsafe() iface.Unsafe {
	return ec.unsafe
}

func (ec *NIOSocketChannel) IsActive() bool {
	return ec.active
}

func (ec *NIOSocketChannel) Bind(address string) {
	ec.pipeline.Bind(address)
}

func (ec *NIOSocketChannel) Pipeline() iface.Pipeline {
	return ec.pipeline
}

func (ec *NIOSocketChannel) SetEventLoop(eventLoop iface.EventLoop) {
	ec.eventloop = eventLoop
}

func (ec *NIOSocketChannel) EventLoop() iface.EventLoop {
	return ec.eventloop
}
