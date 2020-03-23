package epoll

import (
	"errors"
	"net"
	"os"
	"radish/channel"
	"radish/channel/iface"
	"radish/channel/pipeline"
)

const NIODataPackage = "NIODataPackage"

func init() {
	channel.SetChannel("NIODataPackage", NewNIODataPackageChannel)
}

type NIODataPackageChannel struct {
	*FDE
	pipeline  iface.Pipeline
	address   string
	unsafe    iface.Unsafe
	active    bool
	eventloop iface.EventLoop
	network   string
	f         *os.File
	conn      net.PacketConn
	msg chan *iface.Pkg
}

func NewNIODataPackageChannel(conn interface{},network string, address string, fd interface{}) iface.Channel {
	var pc net.PacketConn
	var ok bool
	if conn != nil{
		pc, ok = conn.(net.PacketConn)

		if !ok {
			panic(errors.New("wrong type"))
		}
	}
	np := &NIODataPackageChannel{
		FDE:&FDE{fd:GetFD(fd)},
		address: address,
		network: network,
		msg:make(chan *iface.Pkg,1000),
		conn:pc,
	}

	if GetFD(fd) == 0 {
		pconn, err := net.ListenPacket(np.network, np.address)

		if err != nil {
			panic(err)
		}

		/*uconn, ok := pconn.(*net.UDPConn)

		if !ok {
			panic(errors.New("unknown error"))
		}

		f, err := uconn.File()

		if err != nil {
			panic(err)
		}
		np.f = f
		*/

		np.conn = pconn



		np.fd = GetFD(1)

	}
	np.unsafe = NewByteUnsafe(np)
	np.pipeline = pipeline.NewDefaultChannelPipeline(np)

	return np
}


func (ec *NIODataPackageChannel) Read(msg interface{}) {
	ec.pipeline.ChannelRead(msg)
}

func (ec *NIODataPackageChannel) Write(msg interface{}) {
	ec.pipeline.Write(msg)
}

func (ec *NIODataPackageChannel) ChannelActive(msg interface{}) {
	ec.pipeline.ChannelActive(msg)
}
func (ec *NIODataPackageChannel) ChannelInActive(msg interface{}) {
	ec.pipeline.ChannelInActive(msg)
}

func (ec *NIODataPackageChannel) Unsafe() iface.Unsafe {
	return ec.unsafe
}

func (ec *NIODataPackageChannel) IsActive() bool {
	return ec.active
}

func (ec *NIODataPackageChannel) Bind(address string) {
	ec.pipeline.Bind(address)
}

func (ec *NIODataPackageChannel) Pipeline() iface.Pipeline {
	return ec.pipeline
}

func (ec *NIODataPackageChannel) SetEventLoop(eventLoop iface.EventLoop) {
	ec.eventloop = eventLoop
}

func (ec *NIODataPackageChannel) EventLoop() iface.EventLoop {
	return ec.eventloop
}
