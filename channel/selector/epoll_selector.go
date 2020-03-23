// +build linux

package selector

import (
	"fmt"
	"golang.org/x/sys/unix"
	"radish/channel/iface"
	"sync"
)

const(
	readEvents = unix.EPOLLPRI | unix.EPOLLIN
	writeEvents = unix.EPOLLOUT
	readWriteEvents = readEvents | writeEvents
)

type EpollSelector struct {
	epfd       int
	fd_channel map[int]iface.Channel
	eplock     sync.RWMutex
	size       int
	events     []unix.EpollEvent
}

func OpenSelector() (iface.Selector, error) {
	epfd, err := unix.EpollCreate1(unix.EPOLL_CLOEXEC)

	if err != nil {
		return nil, err
	}

	s := &EpollSelector{
		epfd:       epfd,
		fd_channel: make(map[int]iface.Channel),
		size:       0,
		events:     make([]unix.EpollEvent, 64),
	}

	return s, err
}

func (es *EpollSelector) AddRead(channel iface.Channel) {
	es.eplock.Lock()
	if _, ok := es.fd_channel[channel.FD()]; !ok {
		es.fd_channel[channel.FD()] = channel
		es.size++
		if es.size > len(es.events) {
			es.events = make([]unix.EpollEvent, 2*es.size)
		}
	}
	es.eplock.Unlock()
	err := unix.EpollCtl(es.epfd, unix.EPOLL_CTL_ADD,channel.FD(),&unix.EpollEvent{
		Events: readEvents,
		Fd:     int32(channel.FD()),
	})

	if err != nil{
		panic(err)
	}
}

func (es *EpollSelector) AddWrite(channel iface.Channel) {
	es.eplock.Lock()
	if _, ok := es.fd_channel[channel.FD()]; !ok {
		es.fd_channel[channel.FD()] = channel
		es.size++
		if es.size > len(es.events) {
			es.events = make([]unix.EpollEvent, 2*es.size)
		}
	}
	es.eplock.Unlock()
	err := unix.EpollCtl(es.epfd, unix.EPOLL_CTL_ADD,channel.FD(),&unix.EpollEvent{
		Events: writeEvents,
		Fd:     int32(channel.FD()),
	})

	if err != nil{
		panic(err)
	}
}

func (es *EpollSelector) AddReadWrite(channel iface.Channel) {
	es.eplock.Lock()
	if _, ok := es.fd_channel[channel.FD()]; !ok {
		es.fd_channel[channel.FD()] = channel
		es.size++
		if es.size > len(es.events) {
			es.events = make([]unix.EpollEvent, 2*es.size)
		}
	}
	es.eplock.Unlock()
	err := unix.EpollCtl(es.epfd, unix.EPOLL_CTL_ADD,channel.FD(),&unix.EpollEvent{
		Events: readWriteEvents,
		Fd:     int32(channel.FD()),
	})

	if err != nil{
		panic(err)
	}
}

func (es *EpollSelector) RemoveRead(channel iface.Channel) {
	err := unix.EpollCtl(es.epfd, unix.EPOLL_CTL_DEL,channel.FD(),&unix.EpollEvent{
		Events: readEvents,
		Fd:     int32(channel.FD()),
	})

	if err != nil{
		panic(err)
	}}

func (es *EpollSelector) RemoveWrite(channel iface.Channel) {
	err := unix.EpollCtl(es.epfd, unix.EPOLL_CTL_DEL,channel.FD(),&unix.EpollEvent{
		Events: writeEvents,
		Fd:     int32(channel.FD()),
	})

	if err != nil{
		panic(err)
	}
}

func (es *EpollSelector) RemoveReadWrite(channel iface.Channel) {
	err := unix.EpollCtl(es.epfd, unix.EPOLL_CTL_DEL,channel.FD(),&unix.EpollEvent{
		Events: readWriteEvents,
		Fd:     int32(channel.FD()),
	})

	if err != nil{
		panic(err)
	}
}

func (es *EpollSelector) AddInterests(channel iface.Channel, filters int16) error {

	return nil
}

func (es *EpollSelector) RemoveInterests(channel iface.Channel, filters int16) error {
	return nil
}

func (es *EpollSelector) SelectWithTimeout(timeout int64) ([]iface.Key, error) {

	n, err := unix.EpollWait(es.epfd,es.events,int(timeout))
	if err != nil {
		fmt.Println("kevent", err)
		return nil, err
	}
	keys := make([]iface.Key, n)
	for i := 0; i < n; i++ {
		eevent := es.events[i]

		ch, ok := es.fd_channel[int(eevent.Fd)]

		if !ok {
			fmt.Println("channel lost")
			continue
		}
		keys[i] = iface.Key{
			Channel: ch,
			Events:  eevent.Events,
			Fd:      eevent.Fd,
			Pad:     eevent.Pad,
		}
	}

	if n >= len(es.events) {
		es.events = append(es.events, make([]unix.EpollEvent, len(es.events))...)
	}

	return keys, nil
}

func (es *EpollSelector) Select() ([]iface.Key, error) {
	return es.SelectWithTimeout(-1)
}

