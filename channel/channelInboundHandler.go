package channel

type ChannelInboundHandler interface {
	ChannelRead(ctx *ChannelHandlerContext, msg interface{})
	ChannelActive(ctx *ChannelHandlerContext, msg interface{})
	ChannelInActive(ctx *ChannelHandlerContext, msg interface{})
}
