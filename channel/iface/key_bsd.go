//+build  darwin netbsd freebsd openbsd dragonfly

package iface

type Key struct {
	Channel Channel

	Filter int16

	Flags uint16
}
