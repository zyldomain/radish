package channel

import (
	"golang.org/x/sys/unix"
	"radish/channel/iface"
)

type ServerSocketAcceptor struct {
	*ChannelInboundHandlerAdapter

	childHandler iface.ChannelHandler

	childGroup *EpollEventGroup
}

func NewServerSocketAccptor(childHandler iface.ChannelHandler, childGroup *EpollEventGroup) *ServerSocketAcceptor {
	return &ServerSocketAcceptor{
		childHandler: childHandler,
		childGroup:   childGroup,
	}
}

func (ssa *ServerSocketAcceptor) ChannelRead(ctx iface.ChannelHandlerContextInvoker, msg interface{}) {

	sc, ok := msg.(iface.Channel)

	if !ok {
		panic("channel failed")
	}

	ssa.initChannel(sc)

	ctx.FireChannelRead(msg)
}

func (ssa *ServerSocketAcceptor) initChannel(sc iface.Channel) {
	if ssa.childHandler != nil {
		sc.Pipeline().AddLast(ssa.childHandler)
	}

	ssa.childGroup.Next().Register(sc, []int16{unix.EVFILT_READ})
}
