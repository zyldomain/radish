package channel

import (
	"fmt"
	"golang.org/x/sys/unix"
	"radish/channel/iface"
	"sync"
)

type EpollSelector struct {
	epfd       int
	fd_channel map[int]iface.Channel
	eplock     sync.RWMutex
	size       int
	events     []unix.Kevent_t
}

func OpenEpollSelector() (*EpollSelector, error) {
	epfd, err := unix.Kqueue()

	if err != nil {
		return nil, err
	}

	s := &EpollSelector{
		epfd:       epfd,
		fd_channel: make(map[int]iface.Channel),
		size:       0,
		events:     make([]unix.Kevent_t, 64),
	}

	unix.Kevent(s.epfd, []unix.Kevent_t{{Ident: 0, Filter: unix.EVFILT_USER, Flags: unix.EV_ADD | unix.EV_CLEAR}}, nil, nil)
	return s, err
}

func (es *EpollSelector) AddInterests(channel iface.Channel, filters int16) error {
	es.eplock.Lock()
	if _, ok := es.fd_channel[channel.FD()]; !ok {
		es.fd_channel[channel.FD()] = channel
		es.size++
		if es.size > len(es.events) {
			es.events = make([]unix.Kevent_t, 2*es.size)
		}
	}
	es.eplock.Unlock()
	_, err := unix.Kevent(es.epfd, []unix.Kevent_t{{Ident: uint64(channel.FD()), Filter: filters, Flags: unix.EV_ADD}}, nil, nil)

	return err
}

func (es *EpollSelector) RemoveInterests(channel iface.Channel, filters int16) error {
	_, err := unix.Kevent(es.epfd, []unix.Kevent_t{{Ident: uint64(channel.FD()), Filter: filters, Flags: unix.EV_DELETE}}, nil, nil)
	return err
}

func (es *EpollSelector) SelectWithTimeout(timeout *unix.Timespec) ([]iface.Key, error) {
	n, err := unix.Kevent(es.epfd, nil, es.events, timeout)
	if err != nil {
		fmt.Println(err)
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
		}
	}

	return keys, nil
}

func (es *EpollSelector) Select() ([]iface.Key, error) {
	return es.SelectWithTimeout(nil)
}
