package services

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

	if len(files) != 0 {

		if parameters.Output != "" {
			fmt.Fprintln(&report, "=== Insecure Services and Utilized Binaries ===")
		} else {
			color.Fprintln(color.Bold, &report, "=== Insecure Services and Utilized Binaries ===")
		}

		fmt.Fprintln(&report, "Description: Insecure Services and Utilized Binaries")
		fmt.Fprint(&report, utilities.ListPrint("Files", files))
		fmt.Fprintln(&report)
	}

	return report
}

func buildReportJSON(files []string) *structs.JSONReport {
	return structs.NewJSONReport("Insecure Services and Utilized Binaries", files, "Writable services files and writable called service executables")
}

func buildReportMarkdown(files []string) bytes.Buffer {
	var report bytes.Buffer

	fmt.Fprintln(&report, "# Insecure Services and Utilized Binaries")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "## Location(s)")
	for _, file := range files {
		fmt.Fprintf(&report, "* %v\n", file)
	}
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "## Description")
	fmt.Fprintln(&report, "These files are either writable service files or services that utilized an unsecured binary that allows for a confused deputy attack.")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "Attacking these files would allow for ambigous execution whenever the service is running.")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "## Evidence")
	for _, file := range files {
		if strings.Contains(file, "|") {
			split := strings.SplitN(file, "|", 2)
			service := split[0]
			execStart := split[1]

			serviceOutput, err := os.ReadFile(service)
			if err != nil {
				continue
			}

			execStartOutput, err := exec.Command("stat", execStart).Output()
			if err != nil {
				continue
			}

			serviceBasename, err := utilities.Basename(service)
			if err != nil {
				continue
			}

			fmt.Fprintf(&report, "### %v Insecure Binary\n", serviceBasename)
			fmt.Fprintf(&report, "#### %v Configuration\n", serviceBasename)

			for _, line := range strings.Split(strings.TrimSpace(string(serviceOutput)), "\n") {
				line = strings.TrimSpace(line)
				fmt.Fprintf(&report, "> %v\n", line)
			}

			fmt.Fprintln(&report)

			fmt.Fprintf(&report, "#### %v Permissions\n", serviceBasename)
			for _, line := range strings.Split(strings.TrimSpace(string(execStartOutput)), "\n") {
				line = strings.TrimSpace(line)
				fmt.Fprintf(&report, "> %v\n", line)
			}

			fmt.Fprintln(&report)

		} else {
			output, err := exec.Command("stat", file).Output()
			if err == nil {
				fmt.Fprintf(&report, "### %v Permissions\n", file)

				for _, line := range strings.Split(strings.TrimSpace(string(output)), "\n") {
					line = strings.TrimSpace(line)
					fmt.Fprintf(&report, "> %v\n", line)
				}

				fmt.Fprintln(&report)
			}
		}
	}

	fmt.Fprintln(&report, "## Recommendations")
	fmt.Fprintln(&report, "Ensure that service files are not writable by anyone besides those who need to")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "Use the following command to remove write permissions:")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "### Other users")
	fmt.Fprintln(&report, "> chmod o-w [filepath]")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "### Groups")
	fmt.Fprintln(&report, "> chmod g-w [filepath]")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "## References")
	fmt.Fprintln(&report, "* https://www.geeksforgeeks.org/permissions-in-linux/")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report)

	return report
}
