//go:build windows

package systemInfo

import (
	"fmt"
	"os"
	"runtime"

	"golang.org/x/sys/windows/registry"
)

func populate() {
	OS = runtime.GOOS
	Architecture = runtime.GOARCH
	Container = false

	registeryPopulate()
	nodeName()
}

func registeryPopulate() {

	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if err == nil {
		defer k.Close()

		cv, _, err := k.GetStringValue("CurrentVersion")
		if err == nil {
			Version = fmt.Sprint(cv)
		}

		pn, _, err := k.GetStringValue("ProductName")
		if err == nil {
			Build = fmt.Sprint(pn)
		}

		maj, _, err := k.GetIntegerValue("CurrentMajorVersionNumber")
		if err == nil {
			min, _, err := k.GetIntegerValue("CurrentMinorVersionNumber")
			if err == nil {
				Release = fmt.Sprintf("%d.%d", maj, min)
			}
			Release = fmt.Sprintf("%d", maj)
		}

		cb, _, err := k.GetStringValue("CurrentBuild")
		if err == nil {
			BuildVersion = fmt.Sprint(cb)
		}
	}

}

func nodeName() {
	hostname, err := os.Hostname()
	if err == nil {
		NodeName = hostname
	}
}
