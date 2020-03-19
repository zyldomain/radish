package channel

import (
	"errors"
)

type HeadHandler struct {
	unsafe Unsafe
}

func NewHeadContext(pipeline Pipeline) *ChannelHandlerContext {
	ctx := NewChannelHandlerContext(pipeline, &HeadHandler{
		unsafe: pipeline.Channel().Unsafe(),
	})
	return ctx
}

func (h *HeadHandler) ChannelHandlerAdded(ctx *ChannelHandlerContext) {

}
func (h *HeadHandler) ChannelHandlerRemoved(ctx *ChannelHandlerContext) {

}

func (h *HeadHandler) Write(ctx *ChannelHandlerContext, msg interface{}) {

	buf, ok := msg.([]byte)

	if !ok {
		panic(errors.New("write failed"))
	}
	h.unsafe.Write(buf)
}

func (h *HeadHandler) Bind(ctx *ChannelHandlerContext, address string) {
	h.unsafe.Bind(address)
}
