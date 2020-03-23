package core

import (
	"errors"
	"radish/channel"
	"radish/channel/iface"
	"strings"
	"sync"
)

type ServerBootstrap struct {
	childGroup iface.EventGroup

	parentGroup iface.EventGroup

	childHandler iface.ChannelHandler

	parentHandler iface.ChannelHandler

	wg sync.WaitGroup

	serverSocketChannel string

	network string
}

func NewServerBootstrap() *ServerBootstrap {
	sb := &ServerBootstrap{}
	sb.network = "tcp"
	return sb
}

func (b *ServerBootstrap) ChildGroup(cg iface.EventGroup) *ServerBootstrap {
	b.childGroup = cg

	return b
}

func (b *ServerBootstrap) ParentGroup(pg iface.EventGroup) *ServerBootstrap {
	b.parentGroup = pg
	return b
}

func (b *ServerBootstrap) ChildHandler(handler iface.ChannelHandler) *ServerBootstrap {
	b.childHandler = handler
	return b
}

func (b *ServerBootstrap) ParentHandler(handler iface.ChannelHandler) *ServerBootstrap {
	b.parentHandler = handler

	return b
}

func (b *ServerBootstrap) ServerSocketChannel(name string) *ServerBootstrap {
	b.serverSocketChannel = name
	return b
}

func (b *ServerBootstrap) Bind(address string) *ServerBootstrap {

	if b.childGroup == nil || b.parentGroup == nil {
		panic(errors.New("no executor "))
	}
	b.initAndRegisterChannel(address)

	b.wg.Add(1)
	return b
}

func (b *ServerBootstrap) initAndRegisterChannel(address string) {
	f, err := channel.GetChannel(b.serverSocketChannel)
	if err != nil {
		panic(err)
	}

	ssc := f(nil,b.network, address, 0)
	if b.parentHandler != nil {
		ssc.Pipeline().AddLast(b.parentHandler)
	}
	ssc.Pipeline().AddLast(channel.NewServerSocketAccptor(b.childHandler, b.childGroup))

	b.parentGroup.Next().Register(ssc)
}

func (b *ServerBootstrap) NetWrok(name string) *ServerBootstrap {
	b.network = strings.ToLower(name)
	return b
}

func (b *ServerBootstrap) Sync() {
	b.wg.Wait()
}

func (b *ServerBootstrap) Shutdown() {
	b.wg.Done()
}
