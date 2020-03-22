package epoll

import (
	"errors"
	"golang.org/x/sys/unix"
	"net"
	"os"
	"radish/channel"
	"radish/channel/epoll/udp"
	"radish/channel/iface"
	"radish/channel/pipeline"
	"radish/channel/util"
)

const NIODataPackage = "NIOSocket"

func init() {
	channel.SetChannel("NIOSocket", NewNIODataPackageChannel)
}

type NIODataPackageChannel struct {
	fd        int
	pipeline  iface.Pipeline
	address   string
	unsafe    iface.Unsafe
	active    bool
	eventloop iface.EventLoop
	network   string
	f         *os.File
	conn      net.PacketConn
}

func NewNIODataPackageChannel(network string, address string, fd int) iface.Channel {

	np := &NIODataPackageChannel{
		fd:      fd,
		address: address,
		network: network,
	}

	if fd == -1 {
		pconn, err := net.ListenPacket(np.network, np.address)

		if err != nil {
			panic(err)
		}

		uconn, ok := pconn.(*net.UDPConn)

		if !ok {
			panic(errors.New("unknown error"))
		}

		f, err := uconn.File()

		if err != nil {
			panic(err)
		}

		np.conn = pconn

		np.f = f

		np.fd = int(f.Fd())

	}
	np.unsafe = NewByteUnsafe(np)
	np.pipeline = pipeline.NewDefaultChannelPipeline(np)

	return np
}

func (ec *NIODataPackageChannel) FD() int {
	return ec.fd
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
func (ec *NIODataPackageChannel) doReadMessages(links *util.ArrayList) {
	buf := make([]byte, 2048)
	n, sa, err := unix.Recvfrom(ec.fd, buf, 0)
	if err != nil || n == 0 {
		if err == unix.EAGAIN {
			return
		}
		return
	}

	dp := &udp.DataPackage{Sa: sa}
	if n >= 2048 {
		dp.Data = buf
	} else {
		dp.Data = buf[:n]
	}

	links.Add(dp)
}

func (ec *NIODataPackageChannel) write(msg interface{}) (int, error) {

	dp, ok := msg.(*udp.DataPackage)

	if !ok {
		panic("wrong type")
	}
	err := unix.Sendto(ec.fd, dp.Data, 0, dp.Sa)

	return len(dp.Data), err
}

func (ec *NIODataPackageChannel) bind(address string) {
}
