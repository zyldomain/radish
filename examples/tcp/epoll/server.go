package main

import (
	"flag"
	"fmt"
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
	m, _ := msg.([]byte)

	fmt.Println(string(m))
	ctx.Write(msg)
}

func (p *PrintHandler) ChannelInActive(ctx iface.ChannelHandlerContextInvoker, msg interface{}) {
	fmt.Println("down")
}

type CloseHandler struct {
	pipeline.ChannelOutboundHandlerAdapter
}

func (a *CloseHandler) Close(ctx iface.ChannelHandlerContextInvoker) {
	fmt.Println("close")
	ctx.Close()
}

var address string

func init() {
	flag.StringVar(&address, "addr", "localhost:8080", "-addr localhost:8080")
}
func main() {

	//flag.Parse()
	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num)
	cg := loop.NewEpollEventGroup(1)
	pg := loop.NewEpollEventGroup(1)

	b := core.NewServerBootstrap().
		ServerSocketChannel(epoll.NIOServerSocket).
		NetWrok("tcp").
		ParentGroup(pg).
		ChildGroup(cg).
		ChildHandler(pipeline.NewChannelInitializer(
			func(pipeline iface.Pipeline) {
				pipeline.AddLast(&PrintHandler{}).AddLast(&CloseHandler{})
			}))
	b.Bind("localhost:8080").Sync()

}
