package h3c

import (
	"errors"
	"github.com/smilelhong/netssh/pkg/driver"
)

type NewH3CDevice struct {
	Driver     driver.IDriver
	DeviceType string
	Prompt     string
}

func (h *NewH3CDevice) Connect() error {
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

func (h *NewH3CDevice) Disconnect() {
	h.Driver.Disconnect()
}

func (h *NewH3CDevice) SendCommand(cmd string) (string, error) {
	result, err := h.Driver.SendCommand(cmd, "<"+h.Prompt+">")

	return result, err
}

func (h *NewH3CDevice) SendConfigSet(cmds []string) (string, error) {
	_, err := h.Driver.SendCommand("sys", "[~"+h.Prompt+"]")
	results := ""

	cmds = append(cmds, "return")
	out, err := h.Driver.SendCommandsSet(cmds, "<"+h.Prompt+">")
	results += out

	return results, err
}

func (h *NewH3CDevice) sessionPreparation() error {
	_, err := h.Driver.SendCommand("screen-length disable", "<"+h.Prompt+">")
	if err != nil {
		return errors.New("failed to send command 'screen-length disable' :" + err.Error())
	}
	return nil

}
