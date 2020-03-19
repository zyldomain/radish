package channel

type Pipeline interface {
	AddLast(handler ChannelHandler) Pipeline
	Tail() *ChannelHandlerContext
	Channel() Channel
	ChannelRead(msg interface{})
	Write(msg interface{})
	Bind(address string)
	ChannelActive(msg interface{})
	ChannelInActive(msg interface{})
}
