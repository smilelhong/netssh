package main

import (
	"bufio"
	"fmt"
	"github.com/sirikothe/gotextfsm"
	"github.com/smilelhong/netssh/pkg"
	"os"
)

func main() {
	device, err := netssh.NewDevice("192.168.50.100", "xxxxx", "xxxxx", "huawei_vrp", 22)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = device.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer device.Disconnect()
	//执行单条命令
	//result, err := device.SendCommand("sys")
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//fmt.Println(result)
	//
	//result1, err := device.SendCommand("undo blacklist ip 1.1.1.1")
	//if err != nil {
	//	return
	//}
	//fmt.Println(result1)

	//执行多条命令，默认进入system视图执行
	configs := []string{
		"sys",
		"dis cu | in sysname ",
		"dis cu | in sysname ",
		"dis cu | in sysname ",
	}

	result2, err := device.SendConfigSet(configs)
	if err != nil {
		return
	}
	fmt.Println(result2)
	//temp, err := textFmsTemp(result2)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//for _, v := range temp {d
	//	for i, n := range v {
	//		fmt.Println(i, n)
	//	}
	//	fmt.Println("===")
	//}
}

func textFmsTemp(text string) ([]map[string]interface{}, error) {
	fsm := gotextfsm.TextFSM{}
	file, err := os.Open("./template/display_interface_brief_6881.template")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(file)
	err = fsm.ParseScanner(bufio.NewScanner(file))
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	parser := gotextfsm.ParserOutput{}
	err = parser.ParseTextString(text, fsm, true)
	if err != nil {
		fmt.Printf("Error while parsing input '%s'\n", err.Error())
		return nil, err
	}
	return parser.Dict, nil
}
