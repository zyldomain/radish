// +build windows

package epoll

import (
	"errors"
	"golang.org/x/sys/windows"
)

type FDE struct {
	fd windows.Handle
}


func (f *FDE)FD()windows.Handle{
	return f.fd
}


func GetFD(fd interface{})windows.Handle{
	f, ok := fd.(uintptr)

	if !ok {

		r, ok := fd.(int)

		if ok {
			return windows.Handle(r)
		}
		panic(errors.New("wrong type"))
	}



	return windows.Handle(f)
}

func GetRightFD(fd int)windows.Handle{
	return windows.Handle(fd)
}
