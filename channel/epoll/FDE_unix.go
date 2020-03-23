// +build linux darwin netbsd freebsd openbsd dragonfly


package epoll

import (
	"errors"
)

type FDE struct {
	fd int
}


func (f *FDE)FD()int{
	return f.fd
}


func GetFD(fd interface{})int{
	f, ok := fd.(int)

	if !ok {
		panic(errors.New("wrong type"))
	}

	return f

}

func GetRightFD(fd int)int{
	return fd
}

