//go:build linux

package privileged_container

import (
	"fmt"
	"os"
	"strings"
)

func Enumeration() ([]string, bool) {
	var files []string

	output, err := os.ReadFile("/proc/mounts")
	if err == nil {
		for _, line := range strings.Split(strings.TrimSpace(string(output)), "\n") {
			if strings.HasPrefix(line, "cgroup") {
				perms := strings.Split(line, " ")[3]
				cgroups_perms := strings.Split(perms, ",")[0]
				if cgroups_perms == "ro" {
					return files, false
				} else if cgroups_perms == "rw" {
					return files, true
				} else {
					fmt.Fprintln(os.Stderr, "Invalid cgroups permissions:", cgroups_perms)
				}
			}
		}
	}
	return files, false
}
