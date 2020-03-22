package epoll

import "radish/channel/util"

type AbstractChannel interface {
	doReadMessages(links *util.ArrayList)
}
