// +build windows

package udp

import (
	"golang.org/x/sys/windows"
	"net"
)

type DataPackage struct {
	Data []byte
	Sa   windows.Sockaddr
	Addr net.Addr
}