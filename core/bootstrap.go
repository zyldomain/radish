package core

import (
	"radish/channel"
	"radish/channel/iface"
	"sync"
)

type Bootstrap struct {
	group         iface.EventGroup
	network       string
	handler       iface.ChannelHandler
	wg            sync.WaitGroup
	socketChannel string
	channel       iface.Channel
}

func NewBootstrap() *Bootstrap {
	b := &Bootstrap{network: "tcp"}

	return b
}

func (b *Bootstrap) Group(group iface.EventGroup) *Bootstrap {
	b.group = group
	return b
}

func (b *Bootstrap) Network(network string) *Bootstrap {
	b.network = network

	return b
}

func (b *Bootstrap) Handler(handler iface.ChannelHandler) *Bootstrap {
	b.handler = handler

	return b
}

func (b *Bootstrap) SocketChannel(name string) *Bootstrap {
	b.socketChannel = name
	return b
}

func (b *Bootstrap) Sync() {
	b.wg.Wait()
}

func (b *Bootstrap) Bind(address string) *Bootstrap {
	if b.group == nil {
		panic("no executor")
	}

	b.initAndRegisterChannel(address)

	b.wg.Add(1)
	return b
}

func (b *Bootstrap) initAndRegisterChannel(address string) {
	f, err := channel.GetChannel(b.socketChannel)

	if err != nil {
		panic(err)
	}

	c := f(b.network, address, -1)

	if b.handler != nil {
		c.Pipeline().AddLast(b.handler)
	}
	b.channel = c
	b.group.Next().Register(c)
}

func (b *Bootstrap) Channel() iface.Channel {
	return b.channel
}
