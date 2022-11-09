package huawei

import (
	"errors"
	"netssh/connections"
	"netssh/driver"
	"netssh/types"
)

func NewDevice(connection connections.Connection, DeviceType string) (types.Device, error) {

	newDriver := driver.NewDriver(connection, "\n")
	switch DeviceType {
	case "huawei_vrp":
		return &HuaWeiV8Device{HuaWeiDevice{
			Driver:     newDriver,
			DeviceType: DeviceType,
			Prompt:     "",
		}}, nil
	default:
		return nil, errors.New("unsupported DeviceType: " + DeviceType)

	}

}
