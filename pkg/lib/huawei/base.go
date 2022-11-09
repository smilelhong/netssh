package huawei

import (
	"errors"
	"github.com/smilelhong/netssh/pkg/driver"
)

type HuaWeiDevice struct {
	Driver     driver.IDriver
	DeviceType string
	Prompt     string
}

func (h *HuaWeiDevice) Connect() error {
	if err := h.Driver.Connect(); err != nil {
		return err
	}
	prompt, err := h.Driver.FindDevicePrompt("[<[](.*)[]>]", "#|>")
	if err != nil {
		return err
	}
	h.Prompt = prompt
	return h.sessionPreparation()
}

func (h *HuaWeiDevice) Disconnect() {
	h.Driver.Disconnect()
}

func (h *HuaWeiDevice) SendCommand(cmd string) (string, error) {
	result, err := h.Driver.SendCommand(cmd, "<"+h.Prompt+">")

	return result, err
}

func (h *HuaWeiDevice) SendConfigSet(cmds []string) (string, error) {

	//results, _ := h.Driver.SendCommand("sys", "["+h.Prompt+"]")

	results := ""

	//cmds = append(cmds, "commit")

	cmds = append(cmds, "return")

	out, err := h.Driver.SendCommandsSet(cmds, "["+h.Prompt+"]")
	results += out

	return results, err
}

func (h *HuaWeiDevice) sessionPreparation() error {
	_, err := h.Driver.SendCommand("screen-length 0 temporary", "<"+h.Prompt+">")
	if err != nil {
		return errors.New("failed to send command 'screen-length 0 temporary' :" + err.Error())
	}
	return nil

}
