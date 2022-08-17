//go:build linux

package docker_socket

import (
	"strings"

	"github.com/goEnum/goEnum/utilities"
)

func Enumeration() ([]string, bool) {
	var files []string

	for _, path := range []string{"/var/", "/run/"} {
		for file := range utilities.IterateOverDir(path) {
			pathSplit := strings.Split(file, "/")
			filename := pathSplit[len(pathSplit)-1]
			if strings.HasPrefix(filename, "docker") && strings.HasSuffix(filename, ".sock") {
				if !utilities.Contains(files, file) {
					files = append(files, file)
				}
			}
		}
	}

	for file := range utilities.IterateOverDirN("/", 1) {
		pathSplit := strings.Split(file, "/")
		filename := pathSplit[len(pathSplit)-1]
		if strings.HasPrefix(filename, "docker") && strings.HasSuffix(filename, ".sock") {
			if !utilities.Contains(files, file) {
				files = append(files, file)
			}
		}
	}

	return files, len(files) != 0
}
