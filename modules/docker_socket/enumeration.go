//go:build !linux

package docker_socket

func Enumeration() ([]string, bool) {
	var files []string

	return files, false
}
