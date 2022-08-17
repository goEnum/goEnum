//go:build linux

package services

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/goEnum/goEnum/utilities"
	"github.com/goEnum/goEnum/utilities/permissions"
)

func Enumeration() ([]string, bool) {
	var files []string

	for file := range utilities.IterateOverDirN("/usr/lib/systemd/system/", 1) {
		if filepath, ok := checkService(file); ok {
			files = append(files, filepath)
		}

		if filepath, ok := checkServiceBinary(file); ok {
			files = append(files, filepath)
		}
	}

	return files, len(files) != 0
}

func checkService(file string) (string, bool) {
	output := ""

	if permissions.RW(file) {
		output += file
	}

	return output, output != ""
}

func getExecStart(file string) (string, error) {
	output, err := os.ReadFile(file)
	if err != nil {
		return string(output), err
	}

	for _, line := range strings.Split(string(output), "\n") {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "[") {
			continue
		}

		split := strings.SplitN(line, "=", 2)

		if len(split) != 2 {
			continue
		}

		if split[0] == "ExecStart" {
			execStart := strings.Split(split[1], " ")[0]
			return execStart, nil
		}
	}

	return "", errors.New("ExecStart not in file")
}

func checkServiceBinary(file string) (string, bool) {
	output := ""

	execStartBinary, err := getExecStart(file)
	if err == nil {
		if permissions.RW(execStartBinary) {
			output = fmt.Sprintf("%v|%v", file, execStartBinary)
		}
	}

	return output, output != ""
}
