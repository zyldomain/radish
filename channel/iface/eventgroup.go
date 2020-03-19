package iface

type EventGroup interface {
	Next() EventLoop
}
