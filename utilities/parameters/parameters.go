package parameters

import (
	"bytes"
	"fmt"
	"strings"
)

var (
	Output   string
	Format   string
	Verbose  bool
	Color    bool
	Host     string
	Port     int
	Username string
	Password string
	OS       string
	Binary   string
	Args     []string
	SSH      bool
	Tmpfs    bool
	Commands []string
)

func RemoteParameters(commands ...string) string {
	var output bytes.Buffer

	for _, command := range Commands {
		jump := false
		for _, denied := range []string{
			"ssh",
			"tmpfs",
		} {
			if command == denied {
				jump = true
				break
			}
		}

		if jump {
			continue
		}

		fmt.Fprintf(&output, "%v ", command)
	}

	if Format != "" {
		fmt.Fprintf(&output, "-f %v ", Format)
	}

	if Verbose {
		fmt.Fprint(&output, "-v ")
	}

	if Color {
		fmt.Fprint(&output, "-c ")
	}

	if len(Args) != 0 {
		fmt.Fprint(&output, strings.Join(Args, " "))
	}

	return output.String()
}
