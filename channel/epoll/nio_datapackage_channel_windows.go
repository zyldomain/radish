// +build windows

package epoll

import (
	"radish/channel/epoll/udp"
	"radish/channel/iface"
	"radish/channel/util"
)

func (ec *NIODataPackageChannel) doReadMessages(links *util.ArrayList) {

}

func (ec *NIODataPackageChannel) write(msg interface{}) (int, error) {
	ec.eventloop.AddPackage(ec, &iface.Pkg{
		Event: iface.WRITE,
		Data:  msg,
	})
	return 0, nil
}

func (ec *NIODataPackageChannel) bind(address string) {
}
func (ec *NIODataPackageChannel) SetNonBolcking() {

}

func (ec *NIODataPackageChannel) AddWriteMsg(pkg *iface.Pkg) {
	ec.msg <- pkg
}

func (ec *NIODataPackageChannel) ReadLoop() {
	for {
		buf := make([]byte, 2048)
		n, sa, err := ec.conn.ReadFrom(buf)
		if err != nil {
			ec.eventloop.RemoveChannel(ec)
			ec.conn.Close()
			break
		}

		ec.eventloop.AddPackage(ec, &iface.Pkg{
			Event: iface.READ,
			Data: &udp.DataPackage{
				Data: buf[:n],
				Addr: sa,
			},
		})
	}
}

func (ec *NIODataPackageChannel) WriteLoop() {
	for p := range ec.msg {
		dp, ok := p.Data.(*udp.DataPackage)
		if !ok {
			panic("wrong type")
		}

		_, err := ec.conn.WriteTo(dp.Data, dp.Addr)

		if err != nil {
			continue
		}

	}
}

func (ec *NIODataPackageChannel) close() {
	ec.active = false
	ec.conn.Close()
	ec.eventloop.RemoveChannel(ec)
	ec.pipeline.ChannelInActive(ec)
}
