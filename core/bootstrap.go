package core

import (
	"errors"
	"golang.org/x/sys/unix"
	"radish/channel"
	"radish/channel/iface"
	"sync"
)

type Bootstrap struct {
	childGroup *channel.EpollEventGroup

	parentGroup *channel.EpollEventGroup

	childHandler iface.ChannelHandler

	parentHandler iface.ChannelHandler

	wg sync.WaitGroup
}

func NewBootstrap() *Bootstrap {
	return &Bootstrap{}
}

func (b *Bootstrap) ChildGroup(cg *channel.EpollEventGroup) *Bootstrap {
	b.childGroup = cg

	return b
}

func (b *Bootstrap) ParentGroup(pg *channel.EpollEventGroup) *Bootstrap {
	b.parentGroup = pg
	return b
}

func (b *Bootstrap) ChildHandler(handler iface.ChannelHandler) *Bootstrap {
	b.childHandler = handler
	return b
}

func (b *Bootstrap) ParentHandler(handler iface.ChannelHandler) *Bootstrap {
	b.parentHandler = handler

	return b
}

func (b *Bootstrap) Bind(address string) *Bootstrap {

	if b.childGroup == nil || b.parentGroup == nil {
		panic(errors.New("no executor "))
	}
	ssc := channel.NewEpollServerSocketChannel(address)

	if b.parentHandler != nil {
		ssc.Pipeline().AddLast(b.parentHandler)
	}
	ssc.Pipeline().AddLast(channel.NewServerSocketAccptor(b.childHandler, b.childGroup))

	b.parentGroup.Next().Register(ssc, []int16{unix.EVFILT_READ})

	/*doBind := func() {
		ssc.Bind(address)
	}
	ssc.EventLoop().AddTask(util.NewTask(doBind))*/

	b.wg.Add(1)
	return b
}

func (b *Bootstrap) Sync() {
	b.wg.Wait()
}

func (b *Bootstrap) Shutdown() {
	b.wg.Done()
}
