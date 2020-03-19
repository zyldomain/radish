package main

import (
	"fmt"
	"radish/channel"
	"radish/channel/iface"
	"radish/core"
)

type PrintHandler struct {
	channel.ChannelInboundHandlerAdapter
}

func (p *PrintHandler) ChannelRead(ctx iface.ChannelHandlerContextInvoker, msg interface{}) {
	b, ok := msg.([]byte)

	if !ok {
		ctx.Write([]byte("数据错误"))
		return
	}

	fmt.Println("客户端发送消息-> " + string(b))
	ctx.FireChannelRead(msg)
}

type ConvertHandler struct {
	channel.ChannelInboundHandlerAdapter
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

	cg := channel.NewEpollEventGroup(4)
	pg := channel.NewEpollEventGroup(4)

	b := core.NewBootstrap().
		ParentGroup(pg).
		ChildGroup(cg).
		ChildHandler(channel.NewChannelInitializer(
			func(pipeline iface.Pipeline) {
				pipeline.AddLast(&PrintHandler{})
				pipeline.AddLast(&ConvertHandler{})
			}))
	b.Bind("localhost:8080").Sync()
}
