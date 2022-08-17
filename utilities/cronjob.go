//go:build !windows

package utilities

import (
	"github.com/goEnum/goEnum/utilities/permissions"
	"os"
	"strings"
)

func CrontabDirectory(directoryPath string, files *[]string) {
	for path := range IterateOverDir(directoryPath) {
		Crontab(path, files)
	}
}

func Crontab(path string, files *[]string) {
	commands, err := crontabCommands(path)
	if err == nil {
		for _, command := range commands {
			if strings.Contains(command, "/") {
				if permissions.RW(command) {
					*files = append(*files, command)
				}
			} else {
				for _, path := range Which(command) {
					if permissions.RW(path) {
						*files = append(*files, path)
					}
				}
			}
		}
	}
}

func crontabCommands(path string) ([]string, error) {
	var commands []string

	file, err := os.ReadFile(path)
	if err != nil {
		return commands, err
	}

	for _, line := range strings.Split(strings.TrimSpace(string(file)), "\n") {
		line = strings.ReplaceAll(strings.TrimSpace(line), "\t", " ")
		line = strings.ReplaceAll(line, "(", "")
		line = strings.ReplaceAll(line, ")", "")
		split := strings.Split(line, " ")
		if strings.HasPrefix(line, "#") {
			continue
		} else if line == "" {
			continue
		}

		if strings.HasPrefix(line, "@") {
			line = strings.TrimSpace(strings.SplitN(line, " ", 2)[1])
		}

		if strings.Contains(line, "*") && len(split) > 6 {
			line = strings.TrimSpace(strings.SplitN(line, " ", 7)[6])
		}

		if strings.Contains(line, "||") {
			for _, subline := range strings.Split(line, "||") {
				subline = strings.TrimSpace(subline)
				if strings.Contains(subline, "&&") {
					for _, subsubline := range strings.Split(subline, "&&") {
						subsubline = strings.TrimSpace(subsubline)
						command := crontabCommandLineParser(subsubline)
						if !Contains(commands, command) {
							commands = append(commands, command)
						}
					}
				} else {
					command := crontabCommandLineParser(subline)
					if !Contains(commands, command) {
						commands = append(commands, command)
					}
				}
			}
		} else if strings.Contains(line, "&&") {
			for _, subline := range strings.Split(line, "&&") {
				subline = strings.TrimSpace(subline)
				command := crontabCommandLineParser(subline)
				if !Contains(commands, command) {
					commands = append(commands, command)
				}
			}
		} else {
			command := crontabCommandLineParser(line)
			if !Contains(commands, command) {
				commands = append(commands, command)
			}
		}
	}
	return commands, nil
}

func crontabCommandLineParser(line string) string {
	split := strings.Split(line, " ")

	if strings.Contains(split[0], "=") {
		if len(split) == 1 {
			return ""
		} else {
			line = strings.SplitN(line, " ", 2)[1]
			split = strings.Split(line, " ")
			for i := 0; i < len(split)-1; i++ {
				if strings.Contains(split[0], "=") {
					line = strings.SplitN(line, " ", 2)[1]
					split = strings.Split(line, " ")
				} else {
					break
				}
			}
		}
	}

	if len(split) > 2 && split[0] == "test" && split[1] == "-x" {
		line = strings.SplitN(line, " ", 3)[2]
		split = strings.Split(line, " ")
	}

	return split[0]
}
