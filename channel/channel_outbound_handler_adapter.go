package channel

type ChannelOutboundHandlerAdapter struct {
}

func (a *ChannelOutboundHandlerAdapter) Write(ctx *ChannelHandlerContext, msg interface{}) {
	ctx.Write(msg)
}
func (a *ChannelOutboundHandlerAdapter) Bind(ctx *ChannelHandlerContext, address string) {
	ctx.Bind(address)
}

func (a *ChannelOutboundHandlerAdapter) ChannelHandlerAdded(ctx *ChannelHandlerContext) {

}

func (a *ChannelOutboundHandlerAdapter) ChannelHandlerRemoved(ctx *ChannelHandlerContext) {

}
