package iface

type Channel interface {
	FD
	Read(msg interface{})

	Write(msg interface{})

	Unsafe() Unsafe

	IsActive() bool

	Bind(address string)

	Pipeline() Pipeline

	SetEventLoop(eventLoop EventLoop)

	SetNonBolcking()

	Close()

	SetActive()

	SetID(id string)
	GetID() string
}
