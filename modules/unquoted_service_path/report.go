//go:build !windows

package unquoted_service_path

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

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

	return report
}
