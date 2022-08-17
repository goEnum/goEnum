package docker_socket

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
		fmt.Fprintln(&report, "=== Container with Docker Socket ===")
	} else {
		color.Fprintln(color.Bold, &report, "=== Container with Docker Socket ===")
	}

	fmt.Fprintln(&report, "Description: Container contains a Docker socket which allows for creation and execution of vulnernable containers allowing for privileged escape")
	fmt.Fprint(&report, utilities.ListPrint("Files", files))
	fmt.Fprintln(&report)

	return report
}

func buildReportJSON(files []string) *structs.JSONReport {
	return structs.NewJSONReport("Container with Docker Socket", files, "Container contains a Docker socket which allows for creation and execution of vulnernable containers allowing for privileged escape")
}

func buildReportMarkdown(files []string) bytes.Buffer {
	var report bytes.Buffer

	fmt.Fprintln(&report, "# Container with Docker Socket")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "## Locations(s)")
	fmt.Fprintln(&report, "Current container")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "## Description")
	fmt.Fprintln(&report, "Container is running in with a docker socket that connects to the parent machine.")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "This is an insecure container configurations as it allows for the execution of a vulernable container that can be escaped to give a privileged shell.")
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
	fmt.Fprintln(&report, "Running a docker container in a priviledged mode potientally allows for docker escapes.")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "Ensure that container permissions are only what you want them to be and if possible limit them as much as possible.")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "## References")
	fmt.Fprintln(&report, "* https://book.hacktricks.xyz/linux-hardening/privilege-escalation/docker-breakout/docker-breakout-privilege-escalation")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report)

	return report
}
