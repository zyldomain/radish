package idle

import (
	"radish/channel/iface"
	"radish/channel/pipeline"
	"time"
)

type TimeoutServerHandler struct {
	pipeline.ChannelInboundHandlerAdapter
	writeTimeout time.Duration
	readTimeout  time.Duration
	updateTime   time.Time
}

func NewTimeoutServerHandler(wt time.Duration, rt time.Duration) *TimeoutServerHandler {
	is := &TimeoutServerHandler{
		writeTimeout: wt,
		readTimeout:  rt,
	}

	return is
}

func (a *TimeoutServerHandler) ChannelRead(ctx iface.ChannelHandlerContextInvoker, msg interface{}) {
	ctx.FireChannelRead(msg)
}

func (a *TimeoutServerHandler) ChannelHandlerAdded(ctx iface.ChannelHandlerContextInvoker) {

}

type WTjob struct {
	ctx iface.ChannelHandlerContextInvoker
}

func NewWTjob(ctx iface.ChannelHandlerContextInvoker) *WTjob {
	return &WTjob{ctx: ctx}
}
func (w *WTjob) Run() {
	h, ok := w.ctx.Handler().(*TimeoutServerHandler)
	if !ok {
		panic("wrong type")
	}

	ticker := time.NewTicker(h.writeTimeout)
	h.updateTime = time.Now()
	for {
		select {
		case <-ticker.C:
			if time.Now().Sub(h.updateTime) > h.writeTimeout {
				//超时
				w.ctx.Pipeline().Channel().Closed()
			}
		}
	}
}
