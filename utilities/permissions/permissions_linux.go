//go:build linux

package permissions

import (
	"os"
	"syscall"
)

func writable(path string) bool {
	file, err := os.Stat(path)
	if err != nil {
		return false
	}
	mode := file.Mode()

	stat := file.Sys().(*syscall.Stat_t)
	uid := int(stat.Uid)
	gid := int(stat.Gid)

	if os.Getuid() == 0 {
		return false
	}

	if uid == os.Getuid() {
		return false
	}

	if mode.Perm().String()[8] == 'w' {
		return true
	}

	for _, group := range Groups() {
		if gid == group {
			if mode.Perm().String()[5] == 'w' {
				return true
			}
		}
	}

	return false
}

func readable(path string) bool {
	file, err := os.Stat(path)
	if err != nil {
		return false
	}
	mode := file.Mode()

	stat := file.Sys().(*syscall.Stat_t)
	uid := int(stat.Uid)
	gid := int(stat.Gid)

	if os.Getuid() == 0 {
		return false
	}

	if uid == os.Getuid() {
		return false
	}

	if mode.Perm().String()[7] == 'r' {
		return true
	}

	for _, group := range Groups() {
		if gid == group {
			if mode.Perm().String()[4] == 'r' {
				return true
			}
		}
	}

	return false
}

func executable(path string) bool {
	file, err := os.Stat(path)
	if err != nil {
		return false
	}
	mode := file.Mode()

	stat := file.Sys().(*syscall.Stat_t)
	uid := int(stat.Uid)
	gid := int(stat.Gid)

	if os.Getuid() == 0 {
		return false
	}

	if uid == os.Getuid() {
		return false
	}

	if mode.Perm().String()[9] == 'x' {
		return true
	}

	for _, group := range Groups() {
		if gid == group {
			if mode.Perm().String()[6] == 'x' {
				return true
			}
		}
	}

	return false
}

func writableO(path string) bool {
	file, err := os.Stat(path)
	if err != nil {
		return false
	}
	mode := file.Mode()

	stat := file.Sys().(*syscall.Stat_t)
	uid := int(stat.Uid)
	gid := int(stat.Gid)

	if os.Getuid() == 0 {
		return false
	}

	if uid == os.Getuid() {
		return true
	}

	if mode.Perm().String()[8] == 'w' {
		return true
	}

	for _, group := range Groups() {
		if gid == group {
			if mode.Perm().String()[5] == 'w' {
				return true
			}
		}
	}

	return false
}

func readableO(path string) bool {
	file, err := os.Stat(path)
	if err != nil {
		return false
	}
	mode := file.Mode()

	stat := file.Sys().(*syscall.Stat_t)
	uid := int(stat.Uid)
	gid := int(stat.Gid)

	if os.Getuid() == 0 {
		return false
	}

	if uid == os.Getuid() {
		return true
	}

	if mode.Perm().String()[7] == 'r' {
		return true
	}

	for _, group := range Groups() {
		if gid == group {
			if mode.Perm().String()[4] == 'r' {
				return true
			}
		}
	}

	return false
}

func executableO(path string) bool {
	file, err := os.Stat(path)
	if err != nil {
		return false
	}
	mode := file.Mode()

	stat := file.Sys().(*syscall.Stat_t)
	uid := int(stat.Uid)
	gid := int(stat.Gid)

	if os.Getuid() == 0 {
		return false
	}

	if uid == os.Getuid() {
		return true
	}

	if mode.Perm().String()[9] == 'x' {
		return true
	}

	for _, group := range Groups() {
		if gid == group {
			if mode.Perm().String()[6] == 'x' {
				return true
			}
		}
	}

	return false
}
