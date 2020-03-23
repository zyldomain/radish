// +build windows

package iface

type ChannelRW interface {
	WriteLoop()

	ReadLoop()


	AddWriteMsg(pkg *Pkg)
}
