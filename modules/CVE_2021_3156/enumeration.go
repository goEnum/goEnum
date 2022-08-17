//go:build !windows

package CVE_2021_3156

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/goEnum/goEnum/utilities"
)

func Enumeration() ([]string, bool) {
	var files []string

	for _, path := range utilities.Which("sudo") {
		output, err := exec.Command(path, "--version").Output()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		version := strings.Split(strings.Split(strings.TrimSpace(string(output)), "\n")[0], " ")[2]
		patch := 0

		if strings.Contains(version, "p") {
			patch, err = strconv.Atoi(strings.Split(version, "p")[1])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			version = strings.Split(version, "p")[0]
		}

		major_version, err := strconv.Atoi(strings.Split(version, ".")[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		minor_version, err := strconv.Atoi(strings.Split(version, ".")[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		revision, err := strconv.Atoi(strings.Split(version, ".")[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if vulernable_version(major_version, minor_version, revision, patch) {
			if sudoedit(path) {
				if !utilities.Contains(files, path) {
					files = append(files, path)
					files = append(files, fmt.Sprintf("%v%v", path, "edit"))
				}
			}
		}
	}

	return files, len(files) != 0
}

func vulernable_version(major_version int, minor_version int, revision int, patch int) bool {
	// Vulnerable Versions: 1.7.7 - 1.7.10p9, 1.8.2 - 1.8.31p2, 1.9.0 - 1.9.5p1
	if major_version == 1 {
		if minor_version == 7 {
			if revision >= 7 && revision <= 9 {
				return true
			} else if revision == 10 {
				if patch <= 9 {
					return true
				}
			}
		} else if minor_version == 8 {
			if revision >= 2 && revision < 30 {
				return true
			} else if revision == 31 {
				if patch <= 2 {
					return true
				}
			}
		} else if minor_version == 9 {
			if revision >= 0 && revision <= 4 {
				return true
			} else if revision == 5 {
				if patch <= 1 {
					return true
				}
			}
		}
	}
	return false
}

func sudoedit(path string) bool {
	path = fmt.Sprintf("%v%v", strings.TrimSpace(path), "edit")

	var stderr bytes.Buffer
	command := exec.Command(path, "-s", "/")
	command.Stderr = &stderr
	err := command.Run()

	output := stderr.String()
	if err != nil {
		if strings.HasPrefix(strings.TrimSpace(string(output)), "usage") {
			return false
		} else {
			return true
		}
	}

	return false
}
