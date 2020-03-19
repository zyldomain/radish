package main

import (
	"fmt"
	"radish/channel"
	"radish/core"
)

type PrintHandler struct {
	channel.ChannelInboundHandlerAdapter
}

func (p *PrintHandler) ChannelRead(ctx *channel.ChannelHandlerContext, msg interface{}) {
	b, ok := msg.([]byte)

	if !ok {
		fmt.Println(msg)
	} else {
		fmt.Println(string(b))
	}

	ctx.Write([]byte("服务端的消息-> " + string(b) + "\n"))
}

func main() {
	cg := channel.NewEpollEventGroup(4)
	pg := channel.NewEpollEventGroup(4)

	b := core.NewBootstrap().ParentGroup(pg).ChildGroup(cg).ChildHandler(&PrintHandler{})
	b.Bind("localhost:8080").Sync()
}
