package utilities

import (
	"bytes"
	"fmt"

	"github.com/goEnum/goEnum/utilities/color"
	"github.com/goEnum/goEnum/utilities/parameters"
)

func Parameters() string {
	var output bytes.Buffer

	color.Fprintln(color.Bold, &output, "====== Execution Parameters ======")
	fmt.Fprintf(&output, "Output File:   %v\n", parameters.Output)
	fmt.Fprintf(&output, "Output Format: %v\n", parameters.Format)
	fmt.Fprintf(&output, "Verbose:       %v\n", parameters.Verbose)
	fmt.Fprintf(&output, "Color:         %v\n", parameters.Color)
	fmt.Fprintf(&output, "SSH:           %v\n", parameters.SSH)
	fmt.Fprintf(&output, "Host:          %v\n", parameters.Host)
	fmt.Fprintf(&output, "Port:          %v\n", parameters.Port)
	fmt.Fprintf(&output, "Username:      %v\n", parameters.Username)
	fmt.Fprintf(&output, "Password:      %v\n", parameters.Password)
	fmt.Fprintf(&output, "OS:            %v\n", parameters.OS)
	fmt.Fprintf(&output, "Tmpfs:         %v\n", parameters.Tmpfs)
	fmt.Fprint(&output, ListPrint("Commmands", parameters.Commands))
	fmt.Fprint(&output, ListPrint("Args", parameters.Args))

	return output.String()
}

func LocalParameters() string {
	var output bytes.Buffer

	color.Fprintln(color.Bold, &output, "====== Local Parameters ======")
	fmt.Fprintf(&output, "Output File:   %v\n", parameters.Output)
	fmt.Fprintf(&output, "Output Format: %v\n", parameters.Format)
	fmt.Fprintf(&output, "Verbose:       %v\n", parameters.Verbose)
	fmt.Fprintf(&output, "Color:         %v\n", parameters.Color)
	fmt.Fprintf(&output, "SSH:           %v\n", parameters.SSH)
	fmt.Fprintf(&output, "Host:          %v\n", parameters.Host)
	fmt.Fprintf(&output, "Port:          %v\n", parameters.Port)
	fmt.Fprintf(&output, "Username:      %v\n", parameters.Username)
	fmt.Fprintf(&output, "Password:      %v\n", parameters.Password)
	fmt.Fprintf(&output, "OS:            %v\n", parameters.OS)
	fmt.Fprintf(&output, "Tmpfs:         %v\n", parameters.Tmpfs)
	fmt.Fprint(&output, ListPrint("Commmands", parameters.Commands))
	fmt.Fprint(&output, ListPrint("Args", parameters.Args))

	return output.String()
}
