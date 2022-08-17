//go:build !windows

package writable_files

import (
	"github.com/goEnum/goEnum/utilities"
	"github.com/goEnum/goEnum/utilities/permissions"
)

func Enumeration() ([]string, bool) {
	var files []string

	checkWritable("/home/", &files)
	checkWritable("/etc/", &files)
	checkWritable("/Users/", &files)

	return files, len(files) != 0
}

func checkWritable(path string, files *[]string) {
	for filepath := range utilities.IterateOverDirN(path, 10) {
		if permissions.RW(filepath) {
			*files = append(*files, filepath)
		}
	}
}
