package CVE_2021_3156

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
			fmt.Fprintln(&report, "=== CVE-2021-315 ===")
		} else {
			color.Fprintln(color.Bold, &report, "=== CVE-2021-3156 ===")
		}

		fmt.Fprintln(&report, "Description: Vulnerable version of sudo allowing for Heap-Based Buffer Overflow Privelege Escalation")
		fmt.Fprint(&report, utilities.ListPrint("Files", files))
		fmt.Fprintln(&report)
	}

	return report
}

func buildReportJSON(files []string) *structs.JSONReport {
	return structs.NewJSONReport("CVE-2021-3156", files, "Vulnerable version of sudo allowing for Heap-Based Buffer Overflow Privelege Escalation")
}

func buildReportMarkdown(files []string) bytes.Buffer {
	var report bytes.Buffer
	fmt.Fprintln(&report, "# CVE-2021-3156 Sudo Heap-Based Buffer Overflow Privledge Escalation")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "## Location(s)")
	for _, file := range files {
		fmt.Fprintf(&report, "* %v\n", file)
	}
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "## Description")
	fmt.Fprintln(&report, "CVE-2021-3156 is a heap-based buffer overflow privlege escalation vulnerability in the sudo binary.")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "Having a version of the sudo binary that is vulnerable to this allows for root privledge escalation for all users of a system.")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "Sudo release version between **1.7.7 - 1.7.10p9**,  **1.8.2 - 1.8.31p2**, **1.9.0 - 1.9.5p1** is vulnerable to this attack.")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "## Evidence")
	for _, file := range files {
		version, err := exec.Command(file, "--version").Output()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		} else {
			fmt.Fprintf(&report, "### %v Version\n", file)
			for _, line := range strings.Split(strings.TrimSpace(string(version)), "\n") {
				fmt.Fprintf(&report, "> %v\n", line)
			}
			fmt.Fprintln(&report)
		}
	}
	fmt.Fprintln(&report, "## Recommendations")
	fmt.Fprintln(&report, "Update sudo version being used.")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "Ensure that new version is not in the following (inclusive):")
	fmt.Fprintln(&report, "* 1.7.7 - 1.7.10p9")
	fmt.Fprintln(&report, "* 1.8.2 - 1.8.31p2")
	fmt.Fprintln(&report, "* 1.9.0 - 1.9.5p1")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "## Reference(s)")
	fmt.Fprintln(&report, "* https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2021-3156")
	fmt.Fprintln(&report, "* https://nvd.nist.gov/vuln/detail/CVE-2021-3156")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report)

	return report
}
