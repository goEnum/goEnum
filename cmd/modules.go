package cmd

import (
	"github.com/goEnum/goEnum/utilities/parameters"
	"github.com/spf13/cobra"
)

var module = &cobra.Command{
	Use:              "modules",
	Short:            "display all available modules",
	Long:             "display all available modules",
	Args:             cobra.NoArgs,
	Run:              modulesRun,
	PersistentPreRun: modulesParameters,
}

func init() {
	rootCmd.AddCommand(module)
}

func modulesRun(cmd *cobra.Command, args []string) {
	printModules()
}

func modulesParameters(cmd *cobra.Command, args []string) {
	rootParameters(cmd, args)
	parameters.Commands = append(parameters.Commands, "modules")
}
