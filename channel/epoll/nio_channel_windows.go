// +build windows

package epoll

import (
	"fmt"
	"radish/channel/iface"
	"radish/channel/util"
)

func (ec *NIOSocketChannel) doReadMessages(links *util.ArrayList) {

}

func (ec *NIOSocketChannel) write(msg interface{}) (int, error) {
	ec.eventloop.AddPackage(ec, &iface.Pkg{
		Event: iface.WRITE,
		Data:  msg,
	})

	return 0, nil
}

func (ec *NIOSocketChannel) bind(address string) {

}

func (ec *NIOSocketChannel) SetNonBolcking() {

}

func (ec *NIOSocketChannel) AddWriteMsg(pkg *iface.Pkg) {
	ec.msg <- pkg
}

func (ec *NIOSocketChannel) ReadLoop() {
	for {
		buf := make([]byte, 2048)
		n, err := ec.conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			ec.eventloop.RemoveChannel(ec)
			ec.conn.Close()
			break
		}

		ec.eventloop.AddPackage(ec, &iface.Pkg{
			Event: iface.READ,
			Data:  buf[:n],
		})
	}
}

func (ec *NIOSocketChannel) WriteLoop() {

	for p := range ec.msg {
		b, ok := p.Data.([]byte)

		if !ok {
			panic("wrong type")
		}

		_, err := ec.conn.Write(b)

		if err != nil {
			ec.conn.Close()

			close(ec.msg)
			ec.eventloop.RemoveChannel(ec)
			break
		}
	}

}

func (ec *NIOSocketChannel) close() {
	ec.active = false
	ec.conn.Close()
	ec.eventloop.RemoveChannel(ec)
	ec.pipeline.ChannelInActive(ec)
}
