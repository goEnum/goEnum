//go:build linux

package systemInfo

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func populate() {
	OS = runtime.GOOS
	Architecture = runtime.GOARCH

	uname()
	lsb_release()
	container()
}

func uname() {
	release, err := exec.Command("uname", "-r").Output()
	if err == nil {
		Release = strings.TrimSpace(string(release))

		version, err := exec.Command("uname", "-v").Output()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		Version = strings.TrimSpace(string(version))

		nodeName, err := exec.Command("uname", "-n").Output()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		NodeName = strings.TrimSpace(string(nodeName))
	}
}

func lsb_release() {
	output, err := exec.Command("lsb_release", "-a").Output()
	if err == nil {
		for _, line := range strings.Split(string(output), "\n") {
			if strings.HasPrefix(line, "Description") {
				Build = strings.TrimSpace(strings.Split(line, "\t")[1])
			} else if strings.HasPrefix(line, "Release") {
				BuildVersion = strings.TrimSpace(strings.Split(line, "\t")[1])
			}
		}
	}
}

func container() {
	Container = false

	output, err := os.ReadFile("/proc/mounts")
	if err == nil {
		if strings.HasPrefix(strings.TrimSpace(strings.Split(strings.TrimSpace(string(output)), "\n")[0]), "overlay") {
			Container = true
			return
		}
	}

	_, err = os.Stat("/.dockerenv")
	if !errors.Is(err, os.ErrNotExist) {
		Container = true
		return
	}
}
