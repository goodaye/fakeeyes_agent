package machine

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGOOS(t *testing.T) {
	osname := runtime.GOOS
	fmt.Println(osname)
}

func TestMacMachine(t *testing.T) {
	mac := MacMachine{}
	info, err := mac.CollectDeviceInfo()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(info)
}
