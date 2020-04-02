package epoll

import (
	"radish/channel/iface"
	"radish/channel/util"
)

type MessageUnsafe struct {
	channel iface.Channel
}

func NewMessageUnsafe(channel iface.Channel) *MessageUnsafe {
	return &MessageUnsafe{channel: channel}
}

func (u *MessageUnsafe) Read(links *util.ArrayList) {
	c, _ := u.channel.(AbstractChannel)
	c.doReadMessages(links)

}

func (u *MessageUnsafe) Write(msg interface{}) (int, error) {
	c, _ := u.channel.(AbstractChannel)
	return c.write(msg)
}

func (u *MessageUnsafe) Bind(address string) {
	c, _ := u.channel.(AbstractChannel)
	c.bind(address)
}

func (u *MessageUnsafe) Close() {
	c, _ := u.channel.(AbstractChannel)
	c.close()
}
