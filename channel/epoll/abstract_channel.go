package epoll

import "radish/channel/util"

type AbstractChannel interface {
	doReadMessages(links *util.ArrayList)
	write(msg interface{}) (int, error)
	bind(address string)
	close()
}
