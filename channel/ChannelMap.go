package channel

import (
	"errors"
	"radish/channel/iface"
)

var serverSocketChannelMap = make(map[string]func(conn interface{},network string, address string, fd interface{}) iface.Channel)

func SetChannel(name string, f func(conn interface{},network string, address string, fd interface{}) iface.Channel) {
	serverSocketChannelMap[name] = f
}

func GetChannel(name string) (func(conn interface{},network string, address string, fd interface{}) iface.Channel, error) {
	if f, ok := serverSocketChannelMap[name]; ok {
		return f, nil
	}

	return nil, errors.New("can't get this channel" + name)
}
