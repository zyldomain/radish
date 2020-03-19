package channel

import "radish/channel/iface"

type ChannelInitializer struct {
	*ChannelInboundHandlerAdapter
	initChannel func(pipeline iface.Pipeline)
}

func NewChannelInitializer(initChannel func(pipeline iface.Pipeline)) iface.ChannelHandler {
	return &ChannelInitializer{
		initChannel: initChannel,
	}
}

func (ci *ChannelInitializer) ChannelHandlerAdded(ctx iface.ChannelHandlerContextInvoker) {
	c, _ := ctx.(*ChannelHandlerContext)
	ci.initChannel(ctx.Pipeline())

	c.prev.next = c.next

	c.next.prev = c.prev

	c.next = nil
	c.prev = nil
}
