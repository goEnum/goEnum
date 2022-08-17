package container_block_devices

import (
	"github.com/goEnum/goEnum/utilities/systemInfo"
)

func Prereqs() bool {
	return systemInfo.Container
}
