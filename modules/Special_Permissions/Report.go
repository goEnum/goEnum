package special_permissions

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
			fmt.Fprintln(&report, "=== Files with Special Permissions ===")
		} else {
			color.Fprintln(color.Bold, &report, "=== Files with Special Permissions ===")
		}

		fmt.Fprintln(&report, "Description: Files with either SUID or GUID enabled in current path")
		fmt.Fprint(&report, utilities.ListPrint("Files", files))
		fmt.Fprintln(&report)
	}

	return report
}

func buildReportJSON(files []string) *structs.JSONReport {
	return structs.NewJSONReport("Files with Special Permissions", files, "Files with either SUID or GUID enabled in current path.")
}

func buildReportMarkdown(files []string) bytes.Buffer {
	var report bytes.Buffer
	fmt.Fprintln(&report, "# Files with Special Permissions")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "Location(s)")
	for _, file := range files {
		fmt.Fprintf(&report, "* %v\n", file)
	}
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "## Description")
	fmt.Fprintln(&report, "These are file which have either SUID or GUID bits set on them.")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "SUID bits allow for the execution of a file as the owning user. GUID bits allow for the execution of a file as the owning group")
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
	fmt.Fprintln(&report, "There are several reason why SUID and GUID bits are needed for normal environment execution, a file being listed does not mean that it is vulernable.")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "In some conditions, these bit can be used to use other user's permissions maliciously.")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "To remove these permissions, use the following commands:")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "### SUID Removal")
	fmt.Fprintln(&report, "> chmod u-s [filepath]")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "### GUID Removal")
	fmt.Fprintln(&report, "> chmod g-u [filepath]")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "## References")
	fmt.Fprintln(&report, "* https://www.hackingarticles.in/linux-privilege-escalation-using-suid-binaries/")
	fmt.Fprintln(&report, "* https://linuxhint.com/special-permissions-suid-guid-sticky-bit/")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report)

	return report
}
