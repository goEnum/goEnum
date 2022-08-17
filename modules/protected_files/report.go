package protected_files

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
			fmt.Fprintln(&report, "=== Writable Protected Files ===")
		} else {
			color.Fprintln(color.Bold, &report, "=== Writable Protected Files ===")
		}

		fmt.Fprintln(&report, "Description: Files which are used for user managment and permissioning are writable by the running user")
		fmt.Fprint(&report, utilities.ListPrint("Files", files))
		fmt.Fprintln(&report)
	}

	return report
}

func buildReportJSON(files []string) *structs.JSONReport {
	return structs.NewJSONReport("Writable Protected Files", files, "Files which are used for user managment and permissioning are writable by the running user")
}

func buildReportMarkdown(files []string) bytes.Buffer {
	var report bytes.Buffer

	fmt.Fprintln(&report, "# Writable Protected Files")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "## Location(s)")
	for _, file := range files {
		fmt.Fprintf(&report, "* %v\n", file)
	}
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "## Description")
	fmt.Fprintln(&report, "Passwords and Permissions files keep configuration about all users on a system. This is everything from authentication identifiers, shells, and more.")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "Password files like /etc/passwd, /etc/shadow, and /etc/sudoers should not be writable by anyone other than root because you can lock out or gain access to any user by modifying it.")
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
	fmt.Fprintln(&report, "## Recomendations")
	fmt.Fprintln(&report, "These files should not be writable by anyone other than root and those is an administrator group.")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "Use the following commands to harden these files:")
	for _, file := range files {
		switch file {
		case "/etc/passwd":
			fmt.Fprintln(&report, "### /etc/passwd")
			fmt.Fprintln(&report, "> chmod 0644 /etc/passwd")
			fmt.Fprintln(&report, "> chown root:root /etc/passwd")
			fmt.Fprintln(&report)
		case "/etc/shadow":
			fmt.Fprintln(&report, "### /etc/shadow")
			fmt.Fprintln(&report, "> chmod 0640 /etc/shadow")
			fmt.Fprintln(&report, "> chown root:shadow /etc/shadow")
			fmt.Fprintln(&report)
		case "/etc/sudoers":
			fmt.Fprintln(&report, "### /etc/sudoers")
			fmt.Fprintln(&report, "> chmod 0440 /etc/sudoers")
			fmt.Fprintln(&report, "> chown root:root /etc/sudoers")
			fmt.Fprintln(&report)
		default:
			if strings.Contains(file, "/etc/cron") {
				fmt.Fprintf(&report, "### %v\n", file)
				fmt.Fprintf(&report, "> chmod 0644 %v\n", file)
				fmt.Fprintf(&report, "> chown root:root %v\n", file)
			}
		}
	}
	fmt.Fprintln(&report, "## References")
	for _, file := range files {
		switch file {
		case "/etc/passwd":
			fmt.Fprintln(&report, "* https://www.ibm.com/docs/en/aix/7.2?topic=passwords-using-etcpasswd-file")
		case "/etc/shadow":
			fmt.Fprintln(&report, "* https://linuxize.com/post/etc-shadow-file/")
		case "/etc/sudoers":
			fmt.Fprintln(&report, "* https://www.thegeekdiary.com/sudo-etc-sudoers-is-world-writable-how-to-correct-the-permissions-of-sudoers-file/")
		default:
			if strings.Contains(file, "/etc/cron") {
				if !strings.Contains(report.String(), "https://www.armourinfosec.com/linux-privilege-escalation-by-exploiting-cronjobs/") {
					fmt.Fprintln(&report, "* https://www.armourinfosec.com/linux-privilege-escalation-by-exploiting-cronjobs/")
				}
			}
		}
	}
	fmt.Fprintln(&report)
	fmt.Fprintln(&report)

	return report
}
