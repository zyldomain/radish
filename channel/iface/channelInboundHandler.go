package iface

type ChannelInboundHandler interface {
	ChannelRead(ctx ChannelHandlerContextInvoker, msg interface{})
	ChannelActive(ctx ChannelHandlerContextInvoker, msg interface{})
	ChannelInActive(ctx ChannelHandlerContextInvoker, msg interface{})
}
