package channel

type Channel interface {
	FD() int
	Read(msg interface{})

	Write(msg interface{})

	Unsafe() Unsafe

	IsActive() bool

	Bind(address string)

	Pipeline() Pipeline

	SetEventLoop(eventLoop *EpollEventLoop)
}
