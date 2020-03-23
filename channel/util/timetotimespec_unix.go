// +build linux darwin netbsd freebsd openbsd dragonfly

package util

import (
	"golang.org/x/sys/unix"
	"time"
)

const (
	MillionSecond = int64(time.Nanosecond) * 1000000
	Second        = MillionSecond * 1000
)

func TimeToTimeSpec(t int64) unix.Timespec {
	return unix.NsecToTimespec(t)
}
