//go:build darwin

package systemInfo

import (
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
	sw_vers()
	Container = false
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

func sw_vers() {
	output, err := exec.Command("sw_vers").Output()
	if err == nil {
		split := strings.Split(string(output), "\n")
		Build = strings.TrimSpace(strings.Split(split[1], "\t")[1])
		BuildVersion = strings.TrimSpace(strings.Split(split[2], "\t")[1])
	}
}
