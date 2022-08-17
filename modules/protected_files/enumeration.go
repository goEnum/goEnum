//go:build !windows

package protected_files

import (
	"github.com/goEnum/goEnum/utilities"
	"github.com/goEnum/goEnum/utilities/permissions"
	"github.com/goEnum/goEnum/utilities/systemInfo"
)

func Enumeration() ([]string, bool) {
	var files []string
	var checkFiles []string
	checkFiles = append(checkFiles, "/etc/passwd")
	checkFiles = append(checkFiles, "/etc/shadow")
	checkFiles = append(checkFiles, "/etc/sudoers")
	if systemInfo.OS == "linux" {
		checkFiles = append(checkFiles, "/etc/crontab")
		for _, paths := range []string{"/etc/cron.d/", "/etc/cron.hourly/", "/etc/cron.daily/", "/etc/cron.weekly/", "/etc/cron.monthly"} {
			for path := range utilities.IterateOverDir(paths) {
				checkFiles = append(checkFiles, path)
			}
		}
	}

	for _, file := range checkFiles {
		if permissions.RWO(file) {
			files = append(files, file)
		}
	}

	return files, len(files) != 0
}
