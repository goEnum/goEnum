//go:build !windows

package unquoted_service_path

func Enumeration() ([]string, bool) {
	var files []string

	return files, len(files) != 0
}
