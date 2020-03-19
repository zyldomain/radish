package channel

type ChannelInitializer struct {
	*ChannelInboundHandlerAdapter
	initChannel func(pipeline Pipeline)
}

func (ci *ChannelInitializer) ChannelHandlerAdded(ctx *ChannelHandlerContext) {
	ci.initChannel(ctx.Pipeline())

	ctx.prev.next = ctx.next

	ctx.next.prev = ctx.prev

	ctx.next = nil
	ctx.prev = nil
}

func (ci *ChannelInitializer) SetInit(init func(pipeline Pipeline)) *ChannelInitializer {
	ci.initChannel = init
	return ci
}
