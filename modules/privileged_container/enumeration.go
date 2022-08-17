//go:build !linux

package privileged_container

func Enumeration() ([]string, bool) {
	var files []string

	return files, false
}
