//go:build !windows

package special_permissions

import (
	"fmt"

	"github.com/goEnum/goEnum/utilities"
	"github.com/goEnum/goEnum/utilities/permissions"
)

func Enumeration() ([]string, bool) {
	var files []string

	checkSpecialPermissionsPath(&files)
	checkSpecialPermissions("/bin/", &files)
	checkSpecialPermissions("/sbin/", &files)
	checkSpecialPermissions("/usr/", &files)

	return files, len(files) != 0
}

func checkSpecialPermissions(path string, files *[]string) {
	for file := range utilities.IterateOverDir(path) {
		if permissions.SUID(file) {
			if !utilities.Contains(*files, file) {
				*files = append(*files, file)
			}
		} else if permissions.GUID(file) {
			if !utilities.Contains(*files, file) {
				*files = append(*files, file)
			}
		}
	}
}

func checkSpecialPermissionsPath(files *[]string) {
	for _, path := range utilities.IterateOverPath() {
		path = fmt.Sprintf("%v/", path)
		for file := range utilities.IterateOverDir(path) {
			if permissions.SUID(file) {
				if !utilities.Contains(*files, file) {
					*files = append(*files, file)
				}
			} else if permissions.GUID(file) {
				if !utilities.Contains(*files, file) {
					*files = append(*files, file)
				}
			}
		}
	}
}
