package utilities

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Exists(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func Basename(command string) (string, error) {
	command = strings.TrimSpace(command)
	split := strings.Split(command, "/")
	if len(split) == 0 {
		return "", errors.New("empty string")
	} else if split[len(split)-1] == "" {
		return "", errors.New("basename is empty")
	}
	return split[len(split)-1], nil
}

func Which(command string) []string {
	var paths []string

	for _, path := range IterateOverPath() {
		for commandPath := range IterateOverDirN(path, 1) {
			commandBasename, err := Basename(commandPath)
			commandBasename = strings.TrimSpace(commandBasename)
			if err == nil {
				if commandBasename == command {
					paths = append(paths, commandPath)
				}
			}
		}
	}
	return paths
}

func IterateOverDir(path string) chan string {
	channel := make(chan string)

	go func() {
		defer close(channel)
		iterateOverDir(path, channel)
	}()

	return channel
}

func iterateOverDir(path string, channel chan string) {
	files, err := ioutil.ReadDir(path)
	if err == nil {
		for _, file := range files {
			if file.IsDir() {
				if strings.HasSuffix(file.Name(), "/") {
					iterateOverDir(fmt.Sprintf("%v%v", path, file.Name()), channel)
				} else {
					iterateOverDir(fmt.Sprintf("%v%v/", path, file.Name()), channel)
				}
			} else {
				channel <- strings.ReplaceAll(fmt.Sprintf("%v%v", path, file.Name()), "//", "/")
			}
		}

	}
}

func IterateOverDirN(path string, depth int) chan string {
	channel := make(chan string)

	go func() {
		defer close(channel)
		iterateOverDirN(path, depth, channel)
	}()

	return channel
}

func iterateOverDirN(path string, depth int, channel chan string) {
	if depth > 0 {
		if !strings.HasPrefix(path, "/") {
			path += "/"
		}
		files, err := ioutil.ReadDir(path)
		if err == nil {
			for _, file := range files {
				if file.IsDir() {
					if strings.HasSuffix(file.Name(), "/") {
						iterateOverDirN(fmt.Sprintf("%v%v", path, file.Name()), depth-1, channel)
					} else {
						iterateOverDirN(fmt.Sprintf("%v%v/", path, file.Name()), depth-1, channel)
					}
				} else {
					channel <- strings.ReplaceAll(fmt.Sprintf("%v%v", path, file.Name()), "//", "/")
				}
			}
		}
	}
}

func IterateOverPath() []string {
	var paths []string

	path := os.Getenv("PATH")

	if path == "" {
		return paths
	}

	var output []string

	for _, path := range strings.Split(strings.TrimSpace(path), ":") {
		path = strings.TrimSpace(path)
		if strings.HasSuffix(path, "/") {
			output = append(output, path)
		} else {
			output = append(output, path+"/")
		}
	}

	return output
}
