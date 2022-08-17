package privileged_container

import (
	"github.com/goEnum/goEnum/utilities/systemInfo"
)

func Prereqs() bool {
	return systemInfo.Container
}
