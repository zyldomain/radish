package pipeline

import (
	"radish/channel/iface"
)

type TailHandler struct {
}

func NewTailContext(pipeline iface.Pipeline) iface.ChannelHandlerContextInvoker {
	ctx := NewChannelHandlerContext(pipeline, &TailHandler{})
	return ctx
}

func (h *TailHandler) ChannelHandlerAdded(ctx iface.ChannelHandlerContextInvoker) {

}
func (h *TailHandler) ChannelHandlerRemoved(ctx iface.ChannelHandlerContextInvoker) {

}
func (h *TailHandler) ChannelRead(ctx iface.ChannelHandlerContextInvoker, msg interface{}) {

}
func (h *TailHandler) ChannelActive(ctx iface.ChannelHandlerContextInvoker, msg interface{}) {

}
func (h *TailHandler) ChannelInActive(ctx iface.ChannelHandlerContextInvoker, msg interface{}) {

}
