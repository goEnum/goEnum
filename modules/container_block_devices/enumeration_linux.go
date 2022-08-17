//go:build linux

package container_block_devices

import (
	"fmt"
	"os"
	"strings"
)

func Enumeration() ([]string, bool) {
	var files []string

	disks, err := os.ReadFile("/proc/partitions")
	if err != nil {
		return files, false
	}

	for _, line := range strings.Split(strings.TrimSpace(string(disks)), "\n") {
		line = strings.TrimSpace(line)
		split := strings.Fields(line)

		if len(split) != 4 {
			continue
		}

		name := split[3]

		if name == "name" {
			continue
		}

		if strings.HasPrefix(name, "loop") {
			continue
		}

		files = append(files, fmt.Sprintf("/dev/%v", name))
	}
	return files, len(files) != 0
}
