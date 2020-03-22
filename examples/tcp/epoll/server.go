package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"radish/channel/epoll"
	"radish/channel/iface"
	"radish/channel/loop"
	"radish/channel/pipeline"
	"radish/core"
	"runtime"
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

	fmt.Println("客户端发送消息-> " + string(b))
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

	ctx.Write([]byte(string(b)))

}

func main() {
	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num)
	go func() {
		http.ListenAndServe("localhost:8999", nil)
	}()
	cg := loop.NewEpollEventGroup(num)
	pg := loop.NewEpollEventGroup(1)

	b := core.NewServerBootstrap().
		ServerSocketChannel(epoll.NIOServerSocket).
		NetWrok("tcp").
		ParentGroup(cg).
		ChildGroup(pg).
		ChildHandler(pipeline.NewChannelInitializer(
			func(pipeline iface.Pipeline) {
				pipeline.AddLast(&PrintHandler{})
				pipeline.AddLast(&ConvertHandler{})
			}))
	b.Bind("localhost:9001").Sync()

}
