package iface

type ChannelHandlerContextInvoker interface {
	Handler() ChannelHandler
	Pipeline() Pipeline
	Bind(address string)
	FireChannelActive(msg interface{})
	FireChannelInActive(msg interface{})
	FireChannelRead(msg interface{})
	Write(msg interface{})
	RemoveSelf()
}
