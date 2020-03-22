package channel

import (
	"errors"
	"radish/channel/iface"
)

var serverSocketChannelMap = make(map[string]func(network string, address string, fd int) iface.Channel)

func SetChannel(name string, f func(network string, address string, fd int) iface.Channel) {
	serverSocketChannelMap[name] = f
}

func GetChannel(name string) (func(network string, address string, fd int) iface.Channel, error) {
	if f, ok := serverSocketChannelMap[name]; ok {
		return f, nil
	}

	return nil, errors.New("can't get this channel" + name)
}
