package iface

type Selector interface {
	AddInterests(channel Channel, filters int16) error
	RemoveInterests(channel Channel, filters int16) error
	SelectWithTimeout(timeout int64) ([]Key, error)
	Select() ([]Key, error)
	AddRead(channel Channel)
	AddWrite(channel Channel)
	AddReadWrite(channel Channel)
	RemoveRead(channel Channel)
	RemoveWrite(channel Channel)
	RemoveReadWrite(channel Channel)
}
