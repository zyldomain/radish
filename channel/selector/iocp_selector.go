// +build windows

package selector

import (
	"golang.org/x/sys/windows"
	"radish/channel/iface"
	"sync"
)

type IOCPSelector struct {
	epfd       windows.Handle
	fd_channel map[int]iface.Channel
	eplock     sync.RWMutex
	size       int
	events []interface{}
}

func OpenSelector() (iface.Selector, error) {


	return &IOCPSelector{}, nil
}

func (es *IOCPSelector) AddRead(channel iface.Channel) {

}

func (es *IOCPSelector) AddWrite(channel iface.Channel) {

}

func (es *IOCPSelector) AddReadWrite(channel iface.Channel) {

}

func (es *IOCPSelector) RemoveRead(channel iface.Channel) {

}

func (es *IOCPSelector) RemoveWrite(channel iface.Channel) {

}

func (es *IOCPSelector) RemoveReadWrite(channel iface.Channel) {

}

func (es *IOCPSelector) AddInterests(channel iface.Channel, filters int16) error {

	return nil
}

func (es *IOCPSelector) RemoveInterests(channel iface.Channel, filters int16) error {
	return nil
}

func (es *IOCPSelector) SelectWithTimeout(timeout int64) ([]iface.Key, error) {

	return nil, nil
}

func (es *IOCPSelector) Select() ([]iface.Key, error) {
	return nil, nil
}


