package channel

type TailHandler struct {
}

func NewTailContext(pipeline Pipeline) *ChannelHandlerContext {
	ctx := NewChannelHandlerContext(pipeline, &TailHandler{})
	return ctx
}

func (h *TailHandler) ChannelHandlerAdded(ctx *ChannelHandlerContext) {

}
func (h *TailHandler) ChannelHandlerRemoved(ctx *ChannelHandlerContext) {

}
func (h *TailHandler) ChannelRead(ctx *ChannelHandlerContext, msg interface{}) {

}
func (h *TailHandler) ChannelActive(ctx *ChannelHandlerContext, msg interface{}) {

}
func (h *TailHandler) ChannelInActive(ctx *ChannelHandlerContext, msg interface{}) {

}
