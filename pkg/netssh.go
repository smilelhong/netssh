package netssh

import (
	"github.com/smilelhong/netssh/pkg/connections"
	"github.com/smilelhong/netssh/pkg/lib/h3c"
	"github.com/smilelhong/netssh/pkg/lib/huawei"
	"github.com/smilelhong/netssh/pkg/types"
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
	if strings.Contains(DeviceType, "h3c") {
		device, err = h3c.NewDevice(connection, DeviceType)
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
