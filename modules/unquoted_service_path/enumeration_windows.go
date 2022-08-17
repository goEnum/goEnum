//go:build windows

package unquoted_service_path

import (
	"fmt"
	"strings"

	"golang.org/x/sys/windows/registry"
)

func Enumeration() ([]string, bool) {
	var files []string
	path := `SYSTEM\CurrentControlSet\Services`

	key, err := registry.OpenKey(registry.LOCAL_MACHINE, path, registry.ENUMERATE_SUB_KEYS)
	defer key.Close()
	if err != nil {
		return files, false
	}

	subkeys, err := key.ReadSubKeyNames(-1)
	if err != nil {
		fmt.Println(err)
		return files, false
	}

	for _, subkeyName := range subkeys {
		subkey, err := registry.OpenKey(registry.LOCAL_MACHINE, fmt.Sprintf("%v%v%v", path, `\`, subkeyName), registry.QUERY_VALUE)
		if err == nil {
			value, _, err := subkey.GetStringValue("ImagePath")
			if err == nil {
				value = strings.TrimSpace(value)
				if strings.Contains(value, " ") {
					if !strings.HasPrefix(value, `"`) {
						for _, extension := range []string{".bat", ".exe", ".sys"} {
							if strings.Contains(value, extension) {
								servicePath := strings.SplitN(value, extension, 2)[0] + extension
								if strings.Contains(servicePath, " ") {
									files = append(files, fmt.Sprintf(`HKLM\%v\%v`, path, subkeyName))
								}
								break
							} else {
								continue
							}
						}
					}
				}
			}
		}
		subkey.Close()
	}

	return files, len(files) != 0
}
