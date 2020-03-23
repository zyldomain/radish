package iface

const (
	CLOSED    = 0
	CONNECTED = 1
	READ      = 2
	WRITE     = 3
)

type Pkg struct {
	Event int
	Data  interface{}
}
