//+build linux

package loop

import (
	"golang.org/x/sys/unix"
	"radish/channel/iface"
)

const (
	ErrEvents = unix.EPOLLERR | unix.EPOLLHUP | unix.EPOLLRDHUP
	OutEvents = ErrEvents | unix.EPOLLOUT
	InEvents  = ErrEvents | unix.EPOLLIN | unix.EPOLLPRI
)

func (e *EpollEventLoop) processKeys(keys []iface.Key) {
	for _, key := range keys {
		if key.Flags&unix.EV_ERROR != 0 || key.Flags&unix.EV_EOF != 0 {
			unix.Close(key.Channel.FD())
			continue
		}
		if key.Filter == unix.EVFILT_READ {

			key.Channel.Unsafe().Read(e.objList)
			for _, o := range e.objList.Iterator() {
				key.Channel.Read(o)
			}

			e.objList.RemoveAll()
			if key.Channel.FD() == 9 {
				//e.selector.RemoveInterests(key.Channel, key.Filter)

			}
		}

		if key.Filter == unix.EVFILT_WRITE {
			//TODO
		}

	}
}

func (e *EpollEventLoop) reBuildSelector() {

}
