package channel

type ChannelHandler interface {
	ChannelHandlerAdded(ctx *ChannelHandlerContext)
	ChannelHandlerRemoved(ctx *ChannelHandlerContext)
}
