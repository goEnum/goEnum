package container_block_devices

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/goEnum/goEnum/structs"
	"github.com/goEnum/goEnum/utilities"
	"github.com/goEnum/goEnum/utilities/color"
	"github.com/goEnum/goEnum/utilities/parameters"
)

func Report(files []string) *bytes.Buffer {
	var report bytes.Buffer
	switch parameters.Format {
	case "markdown":
		report = buildReportMarkdown(files)
		if parameters.Output != "" {
			utilities.Append(parameters.Output, report)
		}
	case "json":
		jsonReport := buildReportJSON(files)
		if parameters.Output != "" {
			utilities.WriteJSON(parameters.Output, jsonReport)
		}
		data, err := json.Marshal(jsonReport)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		fmt.Fprint(&report, string(data))
	default:
		report = buildReportOutput(files)
	}
	return &report
}

func buildReportOutput(files []string) bytes.Buffer {
	var report bytes.Buffer

	if parameters.Output != "" {
		fmt.Fprintln(&report, "=== Block Devices in Containers ===")
	} else {
		color.Fprintln(color.Bold, &report, "=== Block Devices in Containers ===")
	}

	fmt.Fprintln(&report, "Description: Block Devices can be accidentially included in containers allows for parent filesystem access")
	fmt.Fprint(&report, utilities.ListPrint("Files", files))
	fmt.Fprintln(&report)

	return report
}

func buildReportJSON(files []string) *structs.JSONReport {
	return structs.NewJSONReport("Block Devices in Containers", files, "Block Devices can be accidentially included in containers allows for parent filesystem access")
}

func buildReportMarkdown(files []string) bytes.Buffer {
	var report bytes.Buffer

	fmt.Fprintln(&report, "# Block Devices in Containers")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "## Locations(s)")
	for _, file := range files {
		fmt.Fprintf(&report, "* %v\n", file)
	}
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "## Description")
	fmt.Fprintln(&report, "Containers by default do not container any block device, including a block device could be part of a vulernable misconfiguration.")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "Block devices can be used to obtains read and write access to the parent filesystem and allow for escapes and privlege escaltion.")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "## Evidence")

	for _, path := range files {
		output, err := exec.Command("stat", path).Output()
		if err == nil {
			for _, line := range strings.Split(strings.TrimSpace(string(output)), "\n") {
				line = strings.TrimSpace(line)
				fmt.Fprintf(&report, "> %v\n", line)
			}
			fmt.Fprintln(&report)
		}
	}

	fmt.Fprintln(&report, "## Recommendations")
	fmt.Fprintln(&report, "Review usage and permission of block devices that are included.")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "Most of the time don't have to be included of that, of those occurances, even fewer have to have read and write access.")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "## References")
	fmt.Fprintln(&report, "* https://man7.org/linux/man-pages/man8/debugfs.8.html")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report)

	return report
}
