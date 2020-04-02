// +build linux darwin netbsd freebsd openbsd dragonfly

package epoll

import (
	"golang.org/x/sys/unix"
	"radish/channel/epoll/udp"
	"radish/channel/util"
)

func (ec *NIODataPackageChannel) doReadMessages(links *util.ArrayList) {
	buf := make([]byte, 2048)
	n, sa, err := unix.Recvfrom(ec.fd, buf, 0)
	if err != nil || n == 0 {
		if err == unix.EAGAIN {
			return
		}
		return
	}

	dp := &udp.DataPackage{Sa: sa}
	if n >= 2048 {
		dp.Data = buf
	} else {
		dp.Data = buf[:n]
	}

	links.Add(dp)
}

func (ec *NIODataPackageChannel) write(msg interface{}) (int, error) {

	dp, ok := msg.(*udp.DataPackage)

	if !ok {
		panic("wrong type")
	}
	err := unix.Sendto(ec.fd, dp.Data, 0, dp.Sa)

	return len(dp.Data), err
}

func (ec *NIODataPackageChannel) bind(address string) {
}
func (ec *NIODataPackageChannel) SetNonBolcking() {
	unix.SetNonblock(ec.fd, true)
}
func (ec *NIODataPackageChannel) close() {
	ec.active = false
	unix.Close(ec.fd)
	ec.pipeline.ChannelInActive(ec)
}
