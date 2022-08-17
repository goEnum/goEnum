//go:build !linux

package services

func Enumeration() ([]string, bool) {
	var files []string

	return files, false
}
