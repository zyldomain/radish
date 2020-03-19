package iface

type ChannelHandler interface {
	ChannelHandlerAdded(ctx ChannelHandlerContextInvoker)
	ChannelHandlerRemoved(ctx ChannelHandlerContextInvoker)
}
