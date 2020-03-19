package channel

type ChannelOutboundHandler interface {
	Write(ctx *ChannelHandlerContext, msg interface{})
	Bind(ctx *ChannelHandlerContext, address string)
}
