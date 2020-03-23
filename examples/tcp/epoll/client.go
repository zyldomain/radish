package main

import (
	"fmt"
	"radish/channel/epoll"
	"radish/channel/iface"
	"radish/channel/loop"
	"radish/channel/pipeline"
	"radish/core"
)

type ClientHandler struct {
	pipeline.ChannelInboundHandlerAdapter
}

func (p *ClientHandler) ChannelRead(ctx iface.ChannelHandlerContextInvoker, msg interface{}) {
	b, ok := msg.([]byte)
	if !ok {
		ctx.Write([]byte("数据错误"))
		return
	}

	fmt.Println("服务端回传消息-> " + string(b))
	ctx.FireChannelRead(msg)
}

func main() {
	g := loop.NewEpollEventGroup(1)
	b := core.NewBootstrap().Group(g).Network("tcp").SocketChannel(epoll.NIOSocket).Handler(&ClientHandler{})

	b.Bind("localhost:8080")

	for i := 0; i < 10; i++ {
		b.Channel().Write([]byte("hello"))
	}
	select {}
}
