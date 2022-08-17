package cmd

import (
	"fmt"
	"os"

	"github.com/goEnum/goEnum/utilities"
	"github.com/goEnum/goEnum/utilities/parameters"
	"github.com/spf13/cobra"
)

var sshTmpfsModules = &cobra.Command{
	Use:              "modules",
	Short:            "display all available modules",
	Long:             "display all available modules",
	Args:             cobra.NoArgs,
	Run:              sshTmpfsModulesRun,
	PersistentPreRun: sshTmpfsModulesParameters,
}

func init() {
	sshTmpfs.AddCommand(sshTmpfsModules)
}

func sshTmpfsModulesRun(cmd *cobra.Command, args []string) {
	if parameters.Verbose {
		fmt.Fprintln(os.Stderr, utilities.LocalParameters())
	}

	printModules()
}
func sshTmpfsModulesParameters(cmd *cobra.Command, args []string) {
	sshTmpfsParameters(cmd, args)
	parameters.Commands = append(parameters.Commands, "modules")

}
