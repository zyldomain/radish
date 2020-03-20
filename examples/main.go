package main

import (
	"fmt"
	"radish/channel/epoll"
	"radish/channel/iface"
	"radish/channel/pipeline"
	"radish/core"
)

type PrintHandler struct {
	pipeline.ChannelInboundHandlerAdapter
}

func (p *PrintHandler) ChannelRead(ctx iface.ChannelHandlerContextInvoker, msg interface{}) {
	b, ok := msg.([]byte)

	if !ok {
		ctx.Write([]byte("数据错误"))
		return
	}

	fmt.Print(b, "客户端发送消息-> "+string(b), len(b), "\n")
	ctx.FireChannelRead(msg)
}

type ConvertHandler struct {
	pipeline.ChannelInboundHandlerAdapter
}

func (p *ConvertHandler) ChannelRead(ctx iface.ChannelHandlerContextInvoker, msg interface{}) {
	b, ok := msg.([]byte)

	if !ok {
		ctx.Write([]byte("数据错误"))
		return
	}

	ctx.Write([]byte("服务端收到消息-> " + string(b) + "\n"))

}

func main() {

	cg := epoll.NewEpollEventGroup(4)
	pg := epoll.NewEpollEventGroup(4)

	b := core.NewBootstrap().
		ParentGroup(pg).
		ChildGroup(cg).
		ChildHandler(pipeline.NewChannelInitializer(
			func(pipeline iface.Pipeline) {
				pipeline.AddLast(&PrintHandler{})
				pipeline.AddLast(&ConvertHandler{})
			}))
	b.Bind("localhost:8080").Sync()
}
