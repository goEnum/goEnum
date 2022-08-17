package systemInfo

import (
	"fmt"

	"github.com/goEnum/goEnum/utilities/color"
)

var (
	OS           string
	Build        string
	BuildVersion string
	Release      string
	Version      string
	Architecture string
	NodeName     string
	Container    bool
)

func init() {
	populate()
}

func String() string {
	output := color.Sprintf(color.Bold, "====== System Information ======\n")

	if OS != "" {
		output += fmt.Sprintf("OS:            %v\n", OS)
	}

	if Build != "" {
		output += fmt.Sprintf("Build:         %v\n", Build)
	}

	if BuildVersion != "" {
		output += fmt.Sprintf("Build Version: %v\n", BuildVersion)
	}

	if Release != "" {
		output += fmt.Sprintf("Release:       %v\n", Release)
	}

	if Version != "" {
		output += fmt.Sprintf("Version:       %v\n", Version)
	}

	if Architecture != "" {
		output += fmt.Sprintf("Architecture:  %v\n", Architecture)
	}

	if NodeName != "" {
		output += fmt.Sprintf("Node Name:     %v\n", NodeName)
	}

	output += fmt.Sprintf("Container:     %v\n", Container)

	return output
}
