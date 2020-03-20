package pipeline

import (
	"errors"
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

	buf, ok := msg.([]byte)

	if !ok {
		panic(errors.New("write failed"))
	}
	h.unsafe.Write(buf)
}

func (h *HeadHandler) Bind(ctx iface.ChannelHandlerContextInvoker, address string) {
	h.unsafe.Bind(address)
}
