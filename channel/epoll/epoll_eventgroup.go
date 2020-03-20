package epoll

import (
	"radish/channel/iface"
	"sync/atomic"
)

type EpollEventGroup struct {
	index      int64
	size       int
	eventloops []iface.EventLoop
}

func NewEpollEventGroup(size int) iface.EventGroup {
	loops := make([]iface.EventLoop, size)

	for i := 0; i < len(loops); i++ {
		l, err := NewEpollEventLoop(int64(i))

		if err != nil {
			for j := 0; j < i; j++ {
				loops[j].Shutdown()
			}
			panic(err)
		}
		loops[i] = l
	}

	return &EpollEventGroup{
		index:      0,
		size:       size,
		eventloops: loops,
	}
}

func (eg *EpollEventGroup) Next() iface.EventLoop {
	return eg.eventloops[(atomic.AddInt64(&eg.index, 1) % int64(eg.size))]
}
