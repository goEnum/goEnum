package privileged_container

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
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
		fmt.Fprintln(&report, "=== Privileged Container ===")
	} else {
		color.Fprintln(color.Bold, &report, "=== Privileged Container ===")
	}

	fmt.Fprintln(&report, "Description: Current container is running in privileged mode. Privileged mode is insecure and allow command execution on system running container.")
	fmt.Fprintln(&report)

	return report
}

func buildReportJSON(files []string) *structs.JSONReport {
	return structs.NewJSONReport("Privileged Container", files, "Current container is running in privileged mode. Privileged mode is insecure and allow command execution on system running container.")
}

func buildReportMarkdown(files []string) bytes.Buffer {
	var report bytes.Buffer

	fmt.Fprintln(&report, "# Priviledged Container")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "## Locations(s)")
	fmt.Fprintln(&report, "Current container")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "## Description")
	fmt.Fprintln(&report, "Container is running in a privileged role that allows for exploitation on the parent container.")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "This is an insecure container configurations which can allow for escapes unless properly mitigated.")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "## Evidence")

	file, err := os.ReadFile("/proc/mounts")
	if err == nil {
		for _, line := range strings.Split(strings.TrimSpace(string(file)), "\n") {
			line = strings.TrimSpace(line)
			fmt.Fprintf(&report, "> %v\n", line)
		}
		fmt.Fprintln(&report)
	}

	fmt.Fprintln(&report, "## Recommendations")
	fmt.Fprintln(&report, "Running a docker container in a priviledged mode potientally allows for docker escapes.")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "Ensure that container permissions are only what you want them to be and if possible limit them as much as possible.")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "## References")
	fmt.Fprintln(&report, "* https://blog.trailofbits.com/2019/07/19/understanding-docker-container-escapes/")
	fmt.Fprintln(&report, "* https://phoenixnap.com/kb/docker-privileged")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report)

	return report
}
