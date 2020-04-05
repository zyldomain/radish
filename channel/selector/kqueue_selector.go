// +build darwin netbsd freebsd openbsd dragonfly

package selector

import (
	"fmt"
	"golang.org/x/sys/unix"
	"radish/channel/iface"
	"radish/channel/util"
	"sync"
)

type KqueueSelector struct {
	epfd       int
	fd_channel map[int]iface.Channel
	eplock     sync.RWMutex
	size       int
	events     []unix.Kevent_t
}

func OpenSelector() (iface.Selector, error) {
	epfd, err := unix.Kqueue()

	if err != nil {
		return nil, err
	}

	s := &KqueueSelector{
		epfd:       epfd,
		fd_channel: make(map[int]iface.Channel),
		size:       0,
		events:     make([]unix.Kevent_t, 64),
	}
	return s, err
}

func (es *KqueueSelector) AddRead(channel iface.Channel) {
	es.AddInterests(channel, unix.EVFILT_READ)
}

func (es *KqueueSelector) AddWrite(channel iface.Channel) {
	es.AddInterests(channel, unix.EVFILT_WRITE)
}

func (es *KqueueSelector) AddReadWrite(channel iface.Channel) {
	es.AddRead(channel)
	es.AddWrite(channel)
}

func (es *KqueueSelector) RemoveRead(channel iface.Channel) {
	es.RemoveInterests(channel, unix.EVFILT_READ)
}

func (es *KqueueSelector) RemoveWrite(channel iface.Channel) {
	es.RemoveInterests(channel, unix.EVFILT_WRITE)
}

func (es *KqueueSelector) RemoveReadWrite(channel iface.Channel) {
	es.RemoveRead(channel)
	es.RemoveWrite(channel)
}

func (es *KqueueSelector) AddInterests(channel iface.Channel, filters int16) error {
	es.eplock.Lock()
	if _, ok := es.fd_channel[channel.FD()]; !ok {

		es.size++
		if es.size > len(es.events) {
			es.events = make([]unix.Kevent_t, 2*es.size)
		}
	}
	es.fd_channel[channel.FD()] = channel
	es.eplock.Unlock()
	_, err := unix.Kevent(es.epfd, []unix.Kevent_t{{Ident: uint64(channel.FD()), Filter: filters, Flags: unix.EV_ADD}}, nil, nil)

	return err
}

func (es *KqueueSelector) RemoveInterests(channel iface.Channel, filters int16) error {
	_, err := unix.Kevent(es.epfd, []unix.Kevent_t{{Ident: uint64(channel.FD()), Filter: filters, Flags: unix.EV_DELETE}}, nil, nil)
	return err
}

func (es *KqueueSelector) SelectWithTimeout(timeout int64) ([]iface.Key, error) {
	var t *unix.Timespec
	if timeout == 0 {
		t = nil
	} else {
		to := util.TimeToTimeSpec(timeout)
		t = &to
	}
	n, err := unix.Kevent(es.epfd, nil, es.events, t)
	if err != nil {
		fmt.Println("kevent", err)
		return nil, err
	}
	keys := make([]iface.Key, n)
	for i := 0; i < n; i++ {
		kevent := es.events[i]

		ch, ok := es.fd_channel[int(kevent.Ident)]

		if !ok {
			fmt.Println("channel lost")
			continue
		}
		keys[i] = iface.Key{
			Channel: ch,
			Filter:  kevent.Filter,
			Flags:   kevent.Flags,
		}
	}

	if n >= len(es.events) {
		es.events = append(es.events, make([]unix.Kevent_t, len(es.events))...)
	}

	return keys, nil
}

func (es *KqueueSelector) Select() ([]iface.Key, error) {
	return es.SelectWithTimeout(0)
}
