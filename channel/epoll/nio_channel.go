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
	*FDE
	pipeline  iface.Pipeline
	address   string
	unsafe    iface.Unsafe
	active    bool
	eventloop iface.EventLoop
	network   string
	f         *os.File
	conn      net.Conn
	ln        net.Listener
	msg       chan *iface.Pkg
}

const NIOSocket = "NIOSocket"

func init() {
	channel.SetChannel("NIOSocket", NewNIOSocketChannel)
}

func NewNIOSocketChannel(conn interface{},network string, address string, fd interface{}) iface.Channel {
	var c net.Conn
	var ok bool
	if conn != nil{
		c, ok = conn.(net.Conn)

		if !ok {
			panic(errors.New("wrong type"))
		}
	}
	epchannel := &NIOSocketChannel{
		FDE:&FDE{fd:GetFD(fd)},
		network: network,
		address: address,
		msg :make(chan *iface.Pkg, 1000),
		conn:c,
	}

	epchannel.unsafe = NewByteUnsafe(epchannel)
	epchannel.pipeline = pipeline.NewDefaultChannelPipeline(epchannel)

	if GetFD(fd) == 0 {

		conn, err := net.Dial(network, address)
		if err != nil {
			panic(err)
		}

		/*tc, ok := conn.(*net.TCPConn)

		if !ok {
			panic(errors.New("network error"))
		}*/

		/*f, err := tc.File()

		if err != nil {
			panic(err)
		}

		epchannel.f = f*/
		epchannel.fd = GetFD(1)

		epchannel.conn = conn
	}

	return epchannel
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
