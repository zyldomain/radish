package channel

import (
	"golang.org/x/sys/unix"
)

type ServerSocketAcceptor struct {
	*ChannelInboundHandlerAdapter

	childHandler ChannelHandler

	childGroup *EpollEventGroup
}

func NewServerSocketAccptor(childHandler ChannelHandler, childGroup *EpollEventGroup) *ServerSocketAcceptor {
	return &ServerSocketAcceptor{
		childHandler: childHandler,
		childGroup:   childGroup,
	}
}

func (ssa *ServerSocketAcceptor) ChannelRead(ctx *ChannelHandlerContext, msg interface{}) {

	sc, ok := msg.(Channel)

	if !ok {
		panic("channel failed")
	}

	ssa.initChannel(sc)

	ctx.FireChannelRead(msg)
}

func (ssa *ServerSocketAcceptor) initChannel(sc Channel) {
	if ssa.childHandler != nil {
		sc.Pipeline().AddLast(ssa.childHandler)
	}

	ssa.childGroup.Next().Register(sc, []int16{unix.EVFILT_READ, unix.EVFILT_WRITE})
}
