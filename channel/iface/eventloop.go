package iface

import (
	"radish/channel/util"
)

type EventLoop interface {
	StartWork()
	AddTask(task *util.Task)
	Register(channel Channel, interests []int16)
	InEventLoop() bool
	ID() int64
	Shutdown()
}
