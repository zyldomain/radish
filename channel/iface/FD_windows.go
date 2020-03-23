// +build windows

package iface

import "golang.org/x/sys/windows"

type FD interface {
	FD()windows.Handle
}
