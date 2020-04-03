package pipeline

import "radish/channel/iface"

type ChannelOutboundHandlerAdapter struct {
}

func (a *ChannelOutboundHandlerAdapter) Write(ctx iface.ChannelHandlerContextInvoker, msg interface{}) {
	ctx.Write(msg)
}
func (a *ChannelOutboundHandlerAdapter) Bind(ctx iface.ChannelHandlerContextInvoker, address string) {
	ctx.Bind(address)
}

func (a *ChannelOutboundHandlerAdapter) ChannelHandlerAdded(ctx iface.ChannelHandlerContextInvoker) {

}

func (a *ChannelOutboundHandlerAdapter) ChannelHandlerRemoved(ctx iface.ChannelHandlerContextInvoker) {

}

func (a *ChannelOutboundHandlerAdapter) Close(ctx iface.ChannelHandlerContextInvoker) {
	ctx.Close()
}
