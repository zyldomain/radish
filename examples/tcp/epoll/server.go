package main

import (
	"crypto/sha1"
	"encoding/base64"
	"flag"
	"fmt"
	_ "net/http/pprof"
	"radish/channel/epoll"
	"radish/channel/iface"
	"radish/channel/loop"
	"radish/channel/pipeline"
	"radish/core"
	"runtime"
	"strings"
)

type UpgradeHandler struct {
	pipeline.ChannelInboundHandlerAdapter
}

var keyGUID = []byte("258EAFA5-E914-47DA-95CA-C5AB0DC85B11")

func computeAcceptKey(challengeKey string) string {
	h := sha1.New()
	h.Write([]byte(challengeKey))
	h.Write(keyGUID)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (u *UpgradeHandler) ChannelRead(ctx iface.ChannelHandlerContextInvoker, msg interface{}) {

	//HTTP/1.1 101 Switching Protocols
	//Upgrade: websocket
	//Connection: Upgrade
	//Sec-WebSocket-Accept: Mr6+I4dZFT3ZGTVz1Y8lUBgXFNw=
	mp := make(map[string]string)
	m, _ := msg.([]byte)

	for _, v := range strings.Split(string(m), "\n")[1:] {
		arr := strings.Split(v, ": ")
		if len(arr) != 2 {
			continue
		}

		mp[arr[0]] = arr[1][:len(arr[1])-1]
	}

	p := make([]byte, 0)
	challengeKey, ok := mp["Sec-WebSocket-Key"]
	if !ok {
		//panic(errors.New(""))
	}
	p = append(p, "HTTP/1.1 101 Switching Protocols\r\nUpgrade: websocket\r\nConnection: Upgrade\r\nSec-WebSocket-Accept: "...)
	p = append(p, computeAcceptKey(challengeKey)...)
	p = append(p, "\r\n"...)
	p = append(p, "\r\n"...)
	fmt.Println(string(m))
	ctx.Write(p)
	ctx.RemoveSelf()
}

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
				pipeline.AddLast(&UpgradeHandler{}).AddLast(&PrintHandler{})
			}))
	b.Bind("localhost:8080").Sync()

}
