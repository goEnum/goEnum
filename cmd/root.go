package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/goEnum/goEnum/utilities/color"
	"github.com/goEnum/goEnum/utilities/parameters"
	"github.com/spf13/cobra"
)

var (
	report     bytes.Buffer
	wg         sync.WaitGroup
	output_buf chan *bytes.Buffer = make(chan *bytes.Buffer)
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:              "goEnum",
	Short:            "System-Agnostic and Modular Enumeration Framework",
	Long:             "System-Agnostic and Modular Enumeration Framework by Maxwell Fusco",
	Args:             cobra.ArbitraryArgs,
	Run:              rootRun,
	PersistentPreRun: rootParameters,
}

func init() {
	rootCmd.PersistentFlags().StringP(
		"output",
		"o",
		"",
		"output file",
	)

	rootCmd.PersistentFlags().StringP(
		"format",
		"f",
		"",
		"output format [json, markdown]",
	)

	rootCmd.PersistentFlags().BoolP(
		"verbose",
		"v",
		false,
		"verbose output",
	)

	rootCmd.PersistentFlags().BoolP(
		"no-color",
		"c",
		false,
		"disable color output",
	)
}

func rootRun(cmd *cobra.Command, args []string) {
	verboseInformation()

	if len(parameters.Args) != 0 {
		runModulesFromArgs()
	} else {
		runAllModules()
	}

	outputFormat()
}

func rootParameters(cmd *cobra.Command, args []string) {
	if !(output(cmd) &&
		format(cmd) &&
		verbose(cmd) &&
		noColor(cmd)) {
		os.Exit(1)
	}

	parameters.Args = args
}

func output(cmd *cobra.Command) bool {
	correctUsage := true

	value, err := cmd.Flags().GetString("output")
	if err != nil {
		correctUsage = false
		color.Fprintln(color.Red, os.Stderr, "Error: \"output\" parameter failure =>", value)
	}

	if _, err := os.Stat(value); !errors.Is(err, os.ErrNotExist) {
		color.Fprintln(color.Red, os.Stderr, "Error: output file already exists =>", value)
		correctUsage = false
	}

	parameters.Output = value

	return correctUsage
}

func format(cmd *cobra.Command) bool {
	correctUsage := true

	value, err := cmd.Flags().GetString("format")
	if err != nil {
		correctUsage = false
		color.Fprintln(color.Red, os.Stderr, "Error: \"format\" parameter failure =>", value)
	}

	validFormat := false
	for _, format := range []string{"", "json", "markdown"} {
		if format == value {
			validFormat = true
		}
	}

	if !validFormat {
		color.Fprintln(color.Red, os.Stderr, "Error: Unsupported format =>", value)
		color.Fprintln(color.Red, os.Stderr, "\t> Valid Formats: [json, markdown]")
		correctUsage = false
	}

	parameters.Format = value

	return correctUsage
}

func verbose(cmd *cobra.Command) bool {
	correctUsage := true

	value, err := cmd.Flags().GetBool("verbose")
	if err != nil {
		correctUsage = false
		fmt.Fprintln(os.Stderr, "Error: \"verbose\" parameter failure =>", value)
	}

	parameters.Verbose = value

	return correctUsage
}

func noColor(cmd *cobra.Command) bool {
	correctUsage := true

	value, err := cmd.Flags().GetBool("no-color")
	if err != nil {
		correctUsage = false
		fmt.Fprintln(os.Stderr, "Error: \"color\" parameter failure =>", value)
	}

	parameters.Color = !value

	return correctUsage
}
