package permissions

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
)

func RWX(path string) bool {
	return readable(path) && writable(path) && executable(path)
}

func RW(path string) bool {
	return readable(path) && writable(path)
}

func R(path string) bool {
	return readable(path)
}

func RX(path string) bool {
	return readable(path) && executable(path)
}

func WX(path string) bool {
	return writable(path) && executable(path)
}

func W(path string) bool {
	return writable(path)
}

func X(path string) bool {
	return executable(path)
}

func RWXO(path string) bool {
	return readableO(path) && writableO(path) && executableO(path)
}

func RWO(path string) bool {
	return readableO(path) && writableO(path)
}

func RO(path string) bool {
	return readableO(path)
}

func RXO(path string) bool {
	return readableO(path) && executableO(path)
}

func WXO(path string) bool {
	return writableO(path) && executableO(path)
}

func WO(path string) bool {
	return writableO(path)
}

func XO(path string) bool {
	return executableO(path)
}

func SUID(path string) bool {
	file, err := os.Stat(path)
	if err != nil {
		return false
	}
	mode := file.Mode()

	return mode&os.ModeSetuid != 0
}

func GUID(path string) bool {
	file, err := os.Stat(path)
	if err != nil {
		return false
	}
	mode := file.Mode()

	return mode&os.ModeSetgid != 0
}

func Sticky(path string) bool {
	file, err := os.Stat(path)
	if err != nil {
		return false
	}
	mode := file.Mode()

	return mode&os.ModeSticky != 0
}

func Groups() []int {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	groups, err := currentUser.GroupIds()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	var gids []int

	for _, group := range groups {
		gid, err := strconv.Atoi(group)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		gids = append(gids, gid)
	}

	return gids
}
