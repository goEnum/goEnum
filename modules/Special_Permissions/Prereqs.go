package special_permissions

import (
	"github.com/goEnum/goEnum/utilities/systemInfo"
)

func Prereqs() bool {
	return systemInfo.OS == "darwin" || systemInfo.OS == "linux"
}
