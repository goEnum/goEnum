//go:build !(linux || darwin)

package permissions

func writable(path string) bool {
	return false
}

func readable(path string) bool {
	return false
}

func executable(path string) bool {
	return false
}

func writableO(path string) bool {
	return false
}

func readableO(path string) bool {
	return false
}

func executableO(path string) bool {
	return false
}
