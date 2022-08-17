package readable_files

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
			fmt.Fprintln(&report, "=== Mispermissioned Files (Readable) ===")
		} else {
			color.Fprintln(color.Bold, &report, "=== Mispermissioned Files (Readable) ===")
		}

		fmt.Fprintln(&report, "Description: Readable files owned by other users that are not base permissions and can potientally hold sensetive information")
		fmt.Fprint(&report, utilities.ListPrint("Files", files))
		fmt.Fprintln(&report)
	}

	return report
}

func buildReportJSON(files []string) *structs.JSONReport {
	return structs.NewJSONReport("Mispermissioned Files (Writable)", files, "Readable files owned by other users that are not base permissions and can potientally hold sensetive information")
}

func buildReportMarkdown(files []string) bytes.Buffer {
	var report bytes.Buffer
	fmt.Fprintln(&report, "# Mispermissioned Files (Readable)")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "## Location(s)")
	for _, file := range files {
		fmt.Fprintf(&report, "* %v\n", file)
	}
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "## Description")
	fmt.Fprintln(&report, "These are file inside of a path that generally should not have readable files owned by other user.")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "These files have no indication of whether they are malicious; however, they are files of interest to gain more context on.")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "Some of the these files are located inside of other user's home directories mean that there is the potiental for leaked information.")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "## Evidence")
	for _, file := range files {
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
	fmt.Fprintln(&report, "## Recommendations")
	fmt.Fprintln(&report, "Generally files owned by other users should only be readable by the owner or owning group.")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "Use the following command to remove read permissions:")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "### Other users")
	fmt.Fprintln(&report, "> chmod o-r [filepath]")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "### Groups")
	fmt.Fprintln(&report, "> chmod g-r [filepath]")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "## References")
	fmt.Fprintln(&report, "* https://www.geeksforgeeks.org/permissions-in-linux/")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report)

	return report
}
