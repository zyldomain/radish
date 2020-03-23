// +build linux darwin netbsd freebsd openbsd dragonfly

package epoll

import (
	"errors"
)

type FDE struct {
	fd int
}

func (f *FDE) FD() int {
	return f.fd
}

func GetFD(fd interface{}) int {
	f, ok := fd.(uintptr)

	if !ok {

		ff, ok := fd.(int)
		if ok {
			return ff
		}
		panic(errors.New("wrong type"))
	}

	return int(f)

}

func GetRightFD(fd int) int {
	return fd
}
