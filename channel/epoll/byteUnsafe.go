package epoll

import (
	"radish/channel/iface"
	"radish/channel/util"
	"sync"
)

var pool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 4096)
	},
}

type ByteUnsafe struct {
	channel iface.Channel
}

func NewByteUnsafe(channel iface.Channel) *ByteUnsafe {
	return &ByteUnsafe{channel: channel}
}

func (b *ByteUnsafe) Read(links *util.ArrayList) {
	c, _ := b.channel.(AbstractChannel)
	c.doReadMessages(links)
}

func (b *ByteUnsafe) Write(msg interface{}) (int, error) {
	c, _ := b.channel.(AbstractChannel)
	return c.write(msg)
}

func (b *ByteUnsafe) Bind(address string) {
	c, _ := b.channel.(AbstractChannel)
	c.bind(address)
}

func (b *ByteUnsafe) Close() {
	c, _ := b.channel.(AbstractChannel)
	c.close()
}
