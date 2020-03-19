package iface

import (
	"golang.org/x/sys/unix"
)

type Selector interface {
	AddInterests(channel Channel, filters int16) error
	RemoveInterests(channel Channel, filters int16) error
	SelectWithTimeout(timeout *unix.Timespec) ([]Key, error)
	Select() ([]Key, error)
}
