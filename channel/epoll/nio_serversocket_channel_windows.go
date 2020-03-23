// +build windows

package epoll

import (
	"errors"
	"net"
	"radish/channel/iface"
	"radish/channel/util"
)

func (ssc *NIOServerSocketChannel) doReadMessages(links *util.ArrayList) {
}

func (ssc *NIOServerSocketChannel) write(msg interface{}) (int, error) {
	return 0, nil
}

func (ssc *NIOServerSocketChannel) bind(address string) {

}

func (ssc *NIOServerSocketChannel) SetNonBlocking() {

}

func (ssc *NIOServerSocketChannel) ReadLoop() {
	for {
		conn, err := ssc.ln.Accept()
		if err != nil {
			continue
		}

		tconn, ok := conn.(*net.TCPConn)

		if !ok {
			panic(errors.New("wrong type"))
		}

		/*f, err := tconn.File()
		if err != nil{
			tconn.Close()
			continue
		}*/

		c := NewNIOSocketChannel(tconn, ssc.network, conn.RemoteAddr().String(), 1)

		ssc.eventloop.AddPackage(ssc, &iface.Pkg{
			Event: iface.CONNECTED,
			Data:  c,
		})

	}
}

func (ssc *NIOServerSocketChannel) WiteLoop() {

}

func (ssc *NIOServerSocketChannel) AddWriteMsg(pkg *iface.Pkg) {
}
