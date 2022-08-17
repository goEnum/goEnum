package protected_files

import (
	"github.com/goEnum/goEnum/utilities/systemInfo"
)

func Prereqs() bool {
	return systemInfo.OS != "windows"
}
