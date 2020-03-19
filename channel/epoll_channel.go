package channel

import (
	"golang.org/x/sys/unix"
)

type EpollChannel struct {
	fd        int
	pipeline  Pipeline
	sa        unix.Sockaddr
	unsafe    Unsafe
	active    bool
	eventloop *EpollEventLoop
}

func NewEpollChannel(fd int, sa unix.Sockaddr) *EpollChannel {
	epchannel := &EpollChannel{
		fd: fd,
		sa: sa,
	}
	epchannel.unsafe = NewByteUnsafe(epchannel)
	epchannel.pipeline = NewDefaultChannelPipeline(epchannel)

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

func (ec *EpollChannel) Unsafe() Unsafe {
	return ec.unsafe
}

func (ec *EpollChannel) IsActive() bool {
	return ec.active
}

func (ec *EpollChannel) Bind(address string) {
	ec.pipeline.Bind(address)
}

func (ec *EpollChannel) Pipeline() Pipeline {
	return ec.pipeline
}

func (ec *EpollChannel) SetEventLoop(eventLoop *EpollEventLoop) {
	ec.eventloop = eventLoop
}

func (ec *EpollChannel) EventLoop() *EpollEventLoop {
	return ec.eventloop
}
