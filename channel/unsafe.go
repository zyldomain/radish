package channel

import "radish/channel/util"

type Unsafe interface {
	Read(links *util.ArrayList)
	Write([]byte) (int, error)
	Bind(address string)
}
