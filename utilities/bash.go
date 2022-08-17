//go:build !windows

package utilities

import (
	"fmt"
	"github.com/goEnum/goEnum/utilities/permissions"
	"os"
	"strings"
)

func ShellDirectory(directoryPath string, files *[]string) {
	for path := range IterateOverDir(directoryPath) {
		Shell(path, files)
	}
}

func Shell(path string, files *[]string) {
	commands, err := shellCommands(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
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

func shellCommands(path string) ([]string, error) {
	var commands []string

	file, err := os.ReadFile(path)
	if err != nil {
		return commands, err
	}

	for _, line := range strings.Split(strings.TrimSpace(string(file)), "\n") {
		line = strings.ReplaceAll(strings.TrimSpace(line), "\t", " ")
		line = strings.ReplaceAll(strings.ReplaceAll(line, "(", ""), ")", ")")

		if strings.HasPrefix(line, "#") {
			continue
		} else if line == "" {
			continue
		}

		if strings.Contains(line, ";") {
			for _, subline := range strings.Split(line, ";") {
				subline = strings.TrimSpace(subline)
				if strings.Contains(subline, "&&") {
					for _, subsubline := range strings.Split(subline, "&&") {
						subsubline = strings.TrimSpace(subsubline)
						if strings.Contains(subsubline, "||") {
							for _, subsubsubline := range strings.Split(subsubline, "||") {
								subsubsubline = strings.TrimSpace(subsubsubline)
								command := shellCommandLineParser(subsubsubline)
								if !Contains(commands, command) {
									commands = append(commands, command)
								}

							}
						} else {
							command := shellCommandLineParser(subsubline)
							if !Contains(commands, command) {
								commands = append(commands, command)
							}
						}
					}
				} else {
					command := shellCommandLineParser(subline)
					if !Contains(commands, command) {
						commands = append(commands, command)
					}
				}
			}
		} else {
			if strings.Contains(line, "&&") {
				for _, subline := range strings.Split(line, "&&") {
					subline = strings.TrimSpace(subline)
					if strings.Contains(subline, "||") {
						for _, subsubline := range strings.Split(subline, "||") {
							subsubline = strings.TrimSpace(subsubline)
							command := shellCommandLineParser(subsubline)
							if !Contains(commands, command) {
								commands = append(commands, command)
							}

						}
					} else {
						command := shellCommandLineParser(subline)
						if !Contains(commands, command) {
							commands = append(commands, command)
						}
					}
				}
			} else {
				command := shellCommandLineParser(line)
				if !Contains(commands, command) {
					commands = append(commands, command)
				}
			}
		}
	}

	return commands, nil
}

func shellCommandLineParser(line string) string {
	bashReservedWords := []string{
		"!",
		"case",
		"coproc",
		"do",
		"done",
		"elif",
		"else",
		"esac",
		"fi",
		"for",
		"function",
		"if",
		"in",
		"select",
		"until",
		"while",
		"{",
		"}",
		"time",
		"[[",
		"]]",
	}

	invalidCharacters := []string{
		"[",
		"]",
		":",
		"%",
		"\\",
		"$",
		"\"",
		"'",
		"(",
		")",
	}

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

	if split[0] == "then" && len(split) > 1 {
		line = strings.SplitN(line, " ", 2)[1]
		split = strings.Split(line, " ")
	}

	if strings.HasPrefix(split[0], "-") {
		if len(split) == 1 {
			return ""
		} else {
			line = strings.SplitN(line, " ", 2)[1]
			split = strings.Split(line, " ")
			for i := 0; i < len(split)-1; i++ {
				if strings.HasPrefix(split[0], "-") {
					line = strings.SplitN(line, " ", 2)[1]
					split = strings.Split(line, " ")
				} else {
					break
				}
			}
		}
	}

	for _, invalidCharacter := range invalidCharacters {
		if strings.Contains(split[0], invalidCharacter) {
			return ""
		}
	}

	if strings.HasPrefix(line, "-") {
		return ""
	}

	if split[0] == "if" {
		return ""
	}

	if Contains(bashReservedWords, split[0]) {
		return ""
	}

	return split[0]
}
