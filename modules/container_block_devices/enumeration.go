//go:build !linux

package container_block_devices

func Enumeration() ([]string, bool) {
	var files []string

	return files, false
}
