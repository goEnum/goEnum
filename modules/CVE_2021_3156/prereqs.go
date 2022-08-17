package CVE_2021_3156

import (
	"github.com/goEnum/goEnum/utilities/systemInfo"
)

func Prereqs() bool {
	return systemInfo.OS != "windows"
}
