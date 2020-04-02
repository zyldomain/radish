package iface

import "radish/channel/util"

type Unsafe interface {
	Read(links *util.ArrayList)
	Write(msg interface{}) (int, error)
	Bind(address string)
	Close()
}
