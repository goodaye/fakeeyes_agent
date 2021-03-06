package controller

import (
	"fakeeyes_agent/drivers"
	"fakeeyes_agent/drivers/macos"
	"fakeeyes_agent/drivers/raspberry"
	"fakeeyes_agent/drivers/x86linux"
	"fmt"
	"runtime"
	"strings"
)

func DetectMachine() (m drivers.Machine, err error) {

	osname := runtime.GOOS
	arch := runtime.GOARCH
	// hostname, err := os.Hostname()
	// if err != nil {
	// 	return
	// }
	if osname == DevicePlatform.Darwin {
		// 如果是mac电脑
		m = macos.MacOSDriver{}
	} else if osname == DevicePlatform.Linux && strings.HasPrefix(arch, "arm") {
		// 判断是Arm Linux设备
		m = &raspberry.RaspberryMachine{}
	} else if osname == DevicePlatform.Linux {
		// 判断是普通Linux
		m = x86linux.X86LinuxMachine{}
	} else {
		err = fmt.Errorf("can't detect machine environment")
	}
	return m, err
}
