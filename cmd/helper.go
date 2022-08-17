package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/goEnum/goEnum/modules"
	"github.com/goEnum/goEnum/structs"
	"github.com/goEnum/goEnum/utilities"
	"github.com/goEnum/goEnum/utilities/color"
	"github.com/goEnum/goEnum/utilities/parameters"
	"github.com/goEnum/goEnum/utilities/systemInfo"
)

func moduleRunner(module *structs.Module, report *bytes.Buffer, wg *sync.WaitGroup, output_buf *chan *bytes.Buffer) {
	defer wg.Done()
	var (
		files  []string
		passed bool
		output bytes.Buffer
	)
	color.Fprintln(color.Bold, &output, "======", module.Name, "======")
	if module.Prereqs() {
		color.Fprintln(color.Green, &output, "[+] Prereqs: Passed")
		files, passed = module.Enumeration()

		if !passed {
			color.Fprintln(color.Red, &output, "[-] Enumeration: Failed")
			color.Fprintln(color.Yellow, &output, "[*] Reporting: Skipping")
		} else {
			color.Fprintln(color.Green, &output, "[+] Enumeration: Succeeded")
			if parameters.Output != "" {
				color.Fprintln(color.Green, &output, "[+] Reporting")
			} else {
				color.Fprintln(color.Yellow, &output, "[*] Reporting: Skipping")
			}
			fmt.Fprint(report, module.Report(files))
		}
	} else {
		color.Fprintln(color.Red, &output, "[-] Prereqs: Failed")
		color.Fprintln(color.Yellow, &output, "[*] Enumeration: Skipping")
		color.Fprintln(color.Yellow, &output, "[*] Reporting: Skipping")
	}
	if parameters.Verbose {
		fmt.Fprintln(&output)
		fmt.Fprintf(&output, "Vulnerable:       %v\n", passed)
		fmt.Fprintln(&output, utilities.ListPrint("Vulnerable Files", files))
	}
	fmt.Fprintln(&output)
	*output_buf <- &output
}

func verboseInformation() {
	if parameters.Verbose {
		fmt.Fprintln(os.Stderr, systemInfo.String())
	}

	if parameters.Verbose {
		whoami := utilities.Whoami()
		fmt.Fprintln(os.Stderr, utilities.WhoamiString(whoami))
	}

	if parameters.Verbose {
		fmt.Fprintln(os.Stderr, utilities.Parameters())
	}
}

func outputReport(report bytes.Buffer) {
	if parameters.Output == "" {
		color.Fprintln(color.Bold, os.Stderr, "====== Report ======")
		fmt.Fprintln(os.Stderr)
		if report.Len() == 0 {
			fmt.Fprintln(os.Stderr, "No findings")
			fmt.Fprintln(os.Stderr)
		} else {
			io.Copy(os.Stdout, &report)
		}
	}
}

func printModules() {
	format := fmt.Sprintf("[+] %%-%vs => %%v\n", modules.Padding)
	color.Fprintln(color.Bold, os.Stdout, "====== Modules ======")
	fmt.Fprintln(os.Stdout)
	for key, module := range modules.Modules {
		fmt.Fprintf(os.Stdout, format, key, module.Name)
	}
}

func moduleOutput(buffer chan *bytes.Buffer, length int) {
	for i := 0; i < length; i++ {
		output, ok := <-buffer
		if ok {
			io.Copy(os.Stderr, output)
		}
	}
}

func runModulesFromArgs() {
	var args []string
	for _, arg := range parameters.Args {
		arg = strings.TrimSpace(arg)
		if module, ok := modules.Modules[arg]; ok {
			args = append(args, arg)
			wg.Add(1)
			go moduleRunner(module, &report, &wg, &output_buf)
		} else {
			color.Fprintln(color.Red, os.Stderr, "[-] Invalid Module Name: ", color.Sprintf(color.Bold, "\"%v\"", arg))
		}
	}
	moduleOutput(output_buf, len(args))
}

func runAllModules() {
	wg.Add(len(modules.Modules))
	for _, module := range modules.Modules {
		go moduleRunner(module, &report, &wg, &output_buf)
	}
	moduleOutput(output_buf, len(modules.Modules))
}

func outputFormat() {
	close(output_buf)
	wg.Wait()

	outputReport(report)
	fmt.Fprintln(os.Stderr)
}
