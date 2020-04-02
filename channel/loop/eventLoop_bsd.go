//+build  darwin netbsd freebsd openbsd dragonfly

package loop

import (
	"golang.org/x/sys/unix"
	"radish/channel/iface"
)

func (e *EpollEventLoop) processKeys(keys []iface.Key) {
	for _, key := range keys {
		if key.Flags&unix.EV_ERROR != 0 || key.Flags&unix.EV_EOF != 0 {
			key.Channel.Close()
			continue
		}
		if key.Filter == unix.EVFILT_READ {

			key.Channel.Unsafe().Read(e.objList)
			for _, o := range e.objList.Iterator() {
				key.Channel.Read(o)
			}

			e.objList.RemoveAll()
		}

		if key.Filter == unix.EVFILT_WRITE {
			//TODO
		}

	}
}

func (e *EpollEventLoop) reBuildSelector() {

}
