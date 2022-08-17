package unquoted_service_path

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/goEnum/goEnum/structs"
	"github.com/goEnum/goEnum/utilities"
	"github.com/goEnum/goEnum/utilities/color"
	"github.com/goEnum/goEnum/utilities/parameters"
	"golang.org/x/sys/windows/registry"
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
			fmt.Println(err)
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
			fmt.Fprintln(&report, "=== Unquoted Service Path ===")
		} else {
			color.Fprintln(color.Bold, &report, "=== Unquoted Service Path ===")
		}

		fmt.Fprintln(&report, "Description: Unquoted Service Paths allows for code execution as the service runner")
		fmt.Fprint(&report, utilities.ListPrint("Files", files))
		fmt.Fprintln(&report)
	}

	return report
}

func buildReportJSON(files []string) *structs.JSONReport {
	return structs.NewJSONReport("Mispermissioned Files (Writable)", files, "Unquoted Service Paths allows for code execution as the service runner")
}

func buildReportMarkdown(files []string) bytes.Buffer {
	var report bytes.Buffer

	fmt.Fprintln(&report, "# Unquoted Service Path")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "## Location(s)")
	for _, file := range files {
		fmt.Fprintf(&report, "* %v\n", file)
	}
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "## Description")
	fmt.Fprintln(&report, "Unquoted Sevice Paths can result in the code execution as the service runner")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "This vulernability is caused by a service path which has a space in without being quoted, other files matching the partial path can be executed instead of the actual service binary")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "## Evidence")
	fmt.Fprintln(&report)
	for _, file := range files {
		fmt.Fprintln(&report, "###", file)

		path := strings.SplitN(file, `\`, 2)[1]
		key, err := registry.OpenKey(registry.LOCAL_MACHINE, path, registry.QUERY_VALUE)
		if err == nil {
			value, _, err := key.GetStringValue("ImagePath")
			if err == nil {
				fmt.Fprintln(&report, "#### ImagePath")
				fmt.Fprintln(&report, ">", value)
			}
		}
		fmt.Fprintln(&report)
	}
	fmt.Fprintln(&report, "## Recommendations")
	fmt.Fprintln(&report, "Quote the binary file in the \"ImagePath\" value in the service registry key.")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "### Example")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "#### Vulnerable Image Path")
	fmt.Fprintln(&report, `> c:\windows\System32\Vulnerable Service.exe -f option`)
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "#### Secure Image Path")
	fmt.Fprintln(&report, `> "c:\windows\System32\Vulnerable Service.exe" -f option`)
	fmt.Fprintln(&report)
	fmt.Fprintln(&report, "## References")
	fmt.Fprintln(&report, "* https://medium.com/@SumitVerma101/windows-privilege-escalation-part-1-unquoted-service-path-c7a011a8d8ae")
	fmt.Fprintln(&report, "* https://www.ired.team/offensive-security/privilege-escalation/unquoted-service-paths")
	fmt.Fprintln(&report)
	fmt.Fprintln(&report)

	return report
}
