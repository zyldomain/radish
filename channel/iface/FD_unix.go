// +build linux darwin netbsd freebsd openbsd dragonfly

package iface

type FD interface {
	FD()int
}
