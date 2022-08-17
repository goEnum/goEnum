package cmd

import (
	"github.com/goEnum/goEnum/utilities/parameters"
	"github.com/spf13/cobra"
)

var all = &cobra.Command{
	Use:              "all",
	Short:            "run all available modules",
	Long:             "run all available modules",
	Args:             cobra.NoArgs,
	Run:              allRun,
	PersistentPreRun: allParameters,
}

func init() {
	rootCmd.AddCommand(all)
}

func allRun(cmd *cobra.Command, args []string) {
	verboseInformation()
	runAllModules()
	outputFormat()
}
func allParameters(cmd *cobra.Command, args []string) {
	rootParameters(cmd, args)
	parameters.Commands = append(parameters.Commands, "all")
}
