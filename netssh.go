package netssh

import (
	"netssh/connections"
	huawei "netssh/lib/huawei"
	"netssh/types"
	"strings"
)

func NewDevice(Host string, Username string, Password string, DeviceType string, Port uint8, Options ...DeviceOption) (types.Device, error) {
	var device types.Device

	//create connection
	connection, err := connections.NewConnection(Host, Username, Password, "ssh", Port)
	if err != nil {
		return nil, err
	}

	if strings.Contains(DeviceType, "huawei") {
		device, err = huawei.NewDevice(connection, DeviceType)
	}

	// running Options Functions.
	for _, option := range Options {
		err := option(device)
		if err != nil {
			return nil, err
		}
	}

	return device, nil
}
