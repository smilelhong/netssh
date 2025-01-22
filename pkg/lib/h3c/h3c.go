package h3c

import (
	"errors"
	"github.com/smilelhong/netssh/pkg/connections"
	"github.com/smilelhong/netssh/pkg/driver"
	"github.com/smilelhong/netssh/pkg/types"
)

func NewDevice(connection connections.Connection, DeviceType string) (types.Device, error) {

	newDriver := driver.NewDriver(connection, "\n")
	switch DeviceType {
	case "h3c":
		return &NewH3CDevice{
			Driver:     newDriver,
			DeviceType: DeviceType,
			Prompt:     "",
		}, nil
	default:
		return nil, errors.New("unsupported DeviceType: " + DeviceType)

	}

}
