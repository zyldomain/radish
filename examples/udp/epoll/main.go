package main

import (
	"errors"
	"fmt"
	"radish/channel/epoll"
	"radish/channel/epoll/udp"
	"radish/channel/iface"
	"radish/channel/loop"
	"radish/channel/pipeline"
	"radish/core"
)

type UDPHandler struct {
	pipeline.ChannelInboundHandlerAdapter
}

func (u *UDPHandler) ChannelRead(ctx iface.ChannelHandlerContextInvoker, msg interface{}) {

	dp, ok := msg.(*udp.DataPackage)

	if !ok {
		panic(errors.New("wrong type"))
	}

	fmt.Println(string(dp.Data))

	ctx.Write(dp)

}

func main() {

	g := loop.NewEpollEventGroup(1)
	b := core.NewBootstrap().Group(g).Network("udp").SocketChannel(epoll.NIODataPackage).Handler(&UDPHandler{})

	b.Bind("localhost:8080")

	select {}
}
