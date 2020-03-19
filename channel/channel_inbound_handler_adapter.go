package channel

import "radish/channel/iface"

type ChannelInboundHandlerAdapter struct {
}

func (a *ChannelInboundHandlerAdapter) ChannelRead(ctx iface.ChannelHandlerContextInvoker, msg interface{}) {
	ctx.FireChannelRead(msg)
}
func (a *ChannelInboundHandlerAdapter) ChannelActive(ctx iface.ChannelHandlerContextInvoker, msg interface{}) {
	ctx.FireChannelActive(msg)
}
func (a *ChannelInboundHandlerAdapter) ChannelInActive(ctx iface.ChannelHandlerContextInvoker, msg interface{}) {
	ctx.FireChannelInActive(msg)
}
func (a *ChannelInboundHandlerAdapter) ChannelHandlerAdded(ctx iface.ChannelHandlerContextInvoker) {

}
func (a *ChannelInboundHandlerAdapter) ChannelHandlerRemoved(ctx iface.ChannelHandlerContextInvoker) {

}
