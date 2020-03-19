package iface

type ChannelOutboundHandler interface {
	Write(ctx ChannelHandlerContextInvoker, msg interface{})
	Bind(ctx ChannelHandlerContextInvoker, address string)
}
