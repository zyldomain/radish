package channel

import (
	"radish/channel/iface"
	"radish/channel/pipeline"
)

type ServerSocketAcceptor struct {
	*pipeline.ChannelInboundHandlerAdapter

	childHandler iface.ChannelHandler

	childGroup iface.EventGroup
}

func NewServerSocketAccptor(childHandler iface.ChannelHandler, childGroup iface.EventGroup) *ServerSocketAcceptor {
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

	ssa.childGroup.Next().Register(sc)
}
