package epoll

import (
	"golang.org/x/sys/unix"
	"radish/channel/iface"
	"radish/channel/pipeline"
)

type EpollChannel struct {
	fd        int
	pipeline  iface.Pipeline
	sa        unix.Sockaddr
	unsafe    iface.Unsafe
	active    bool
	eventloop iface.EventLoop
}

func NewEpollChannel(fd int, sa unix.Sockaddr) *EpollChannel {
	epchannel := &EpollChannel{
		fd: fd,
		sa: sa,
	}
	epchannel.unsafe = NewByteUnsafe(epchannel)
	epchannel.pipeline = pipeline.NewDefaultChannelPipeline(epchannel)

	return epchannel
}

func (ec *EpollChannel) FD() int {
	return ec.fd
}

func (ec *EpollChannel) Read(msg interface{}) {
	ec.pipeline.ChannelRead(msg)
}

func (ec *EpollChannel) Write(msg interface{}) {
	ec.pipeline.Write(msg)
}

func (ec *EpollChannel) ChannelActive(msg interface{}) {
	ec.pipeline.ChannelActive(msg)
}
func (ec *EpollChannel) ChannelInActive(msg interface{}) {
	ec.pipeline.ChannelInActive(msg)
}

func (ec *EpollChannel) Unsafe() iface.Unsafe {
	return ec.unsafe
}

func (ec *EpollChannel) IsActive() bool {
	return ec.active
}

func (ec *EpollChannel) Bind(address string) {
	ec.pipeline.Bind(address)
}

func (ec *EpollChannel) Pipeline() iface.Pipeline {
	return ec.pipeline
}

func (ec *EpollChannel) SetEventLoop(eventLoop iface.EventLoop) {
	ec.eventloop = eventLoop
}

func (ec *EpollChannel) EventLoop() iface.EventLoop {
	return ec.eventloop
}
