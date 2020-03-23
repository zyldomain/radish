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

		if key.Events & InEvents!=0 {

			key.Channel.Unsafe().Read(e.objList)
			for _, o := range e.objList.Iterator() {
				key.Channel.Read(o)
			}

			e.objList.RemoveAll()
		}

		if key.Events & OutEvents != 0 {
			//TODO
		}

	}
}

func (e *EpollEventLoop) reBuildSelector() {

}
