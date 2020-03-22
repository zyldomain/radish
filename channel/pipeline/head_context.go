package pipeline

import (
	"radish/channel/iface"
)

type HeadHandler struct {
	unsafe iface.Unsafe
}

func NewHeadContext(pipeline iface.Pipeline) iface.ChannelHandlerContextInvoker {
	ctx := NewChannelHandlerContext(pipeline, &HeadHandler{
		unsafe: pipeline.Channel().Unsafe(),
	})
	return ctx
}

func (h *HeadHandler) ChannelHandlerAdded(ctx iface.ChannelHandlerContextInvoker) {

}
func (h *HeadHandler) ChannelHandlerRemoved(ctx iface.ChannelHandlerContextInvoker) {

}

func (h *HeadHandler) Write(ctx iface.ChannelHandlerContextInvoker, msg interface{}) {

	h.unsafe.Write(msg)
}

func (h *HeadHandler) Bind(ctx iface.ChannelHandlerContextInvoker, address string) {
	h.unsafe.Bind(address)
}
