package cronjobs

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
			fmt.Fprintln(&report, "=== Cronjobs with Writable Executables ===")
		} else {
			color.Fprintln(color.Bold, &report, "=== Cronjobs with Writable Executables ===")
		}

		fmt.Fprintln(&report, "Description: Cronjobs are using executables with insecure permissions with the potiental to be highjacked")
		fmt.Fprint(&report, utilities.ListPrint("Files", files))
		fmt.Fprintln(&report)
	}

	return report
}

func buildReportJSON(files []string) *structs.JSONReport {
	return structs.NewJSONReport("Cronjobs with Writable Executables", files, "Cronjobs are using executables with insecure permissions with the potiental to be highjacked")
}

func buildReportMarkdown(files []string) bytes.Buffer {
	var report bytes.Buffer
	fmt.Fprintln(&report, "# Cronjobs with Wrtiable Executables")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "## Location(s)")
	for _, file := range files {
		fmt.Fprintf(&report, "* %v\n", file)
	}
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "## Description")
	fmt.Fprintln(&report, "Cronjobs execute on a regular pattern using a designated users permissions.")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "If a cronjob is using an insecure binary (one which an attacker has write access to), an attacker can hijack the cronjob and execute their own commands when the cronjob executes.")
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
	fmt.Fprintln(&report, "Executables should only be writable by root and those in an administator group.")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "Use the following command to harden these executables:")
	fmt.Fprintln(&report, "### Regular Binaries")
	fmt.Fprintln(&report, "Ensure these binaries are in /bin/, /usr/bin/, or similar directories.")
	fmt.Fprintln(&report, "> chmod 755 [filepath]")
	fmt.Fprintln(&report, "> chown root:root [filepath]")
	fmt.Fprintln(&report, "> mv [filepath] /usr/bin/")
	fmt.Fprintln(&report, "> mv [filepath] /bin/")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "### Superuser Binaries")
	fmt.Fprintln(&report, "Ensure binaries which require superuser privledges to execure are located in /sbin/, /usr/sbin/, or similar directories.")
	fmt.Fprintln(&report, "> chmod 755 [filepath]")
	fmt.Fprintln(&report, "> chown root:root [filepath]")
	fmt.Fprintln(&report, "> mv [filepath] /usr/sbin/")
	fmt.Fprintln(&report, "> mv [filepath] /sbin/")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "## References")
	fmt.Fprintln(&report, "* https://www.armourinfosec.com/linux-privilege-escalation-by-exploiting-cronjobs/")
	fmt.Fprintln(&report, "* https://www.hackingarticles.in/linux-privilege-escalation-using-path-variable/")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report)

	return report
}
