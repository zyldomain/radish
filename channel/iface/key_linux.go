//+build linux

package iface

type Key struct {
	Channel Channel

	Events uint32
	Fd     int32
	Pad    int32
}

