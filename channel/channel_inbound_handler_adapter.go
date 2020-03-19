package channel

type ChannelInboundHandlerAdapter struct {
}

func (a *ChannelInboundHandlerAdapter) ChannelRead(ctx *ChannelHandlerContext, msg interface{}) {
	ctx.FireChannelRead(msg)
}
func (a *ChannelInboundHandlerAdapter) ChannelActive(ctx *ChannelHandlerContext, msg interface{}) {
	ctx.FireChannelActive(msg)
}
func (a *ChannelInboundHandlerAdapter) ChannelInActive(ctx *ChannelHandlerContext, msg interface{}) {
	ctx.FireChannelInActive(msg)
}
func (a *ChannelInboundHandlerAdapter) ChannelHandlerAdded(ctx *ChannelHandlerContext) {

}
func (a *ChannelInboundHandlerAdapter) ChannelHandlerRemoved(ctx *ChannelHandlerContext) {

}
