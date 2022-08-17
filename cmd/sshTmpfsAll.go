package cmd

import (
	"fmt"
	"os"

	"github.com/goEnum/goEnum/utilities"
	"github.com/goEnum/goEnum/utilities/color"
	"github.com/goEnum/goEnum/utilities/parameters"
	"github.com/goEnum/goEnum/utilities/remote"
	"github.com/spf13/cobra"
)

var sshTmpfsAll = &cobra.Command{
	Use:              "all",
	Short:            "run all available modules",
	Long:             "run all available modules",
	Args:             cobra.NoArgs,
	Run:              sshTmpfsAllRun,
	PersistentPreRun: sshTmpfsAllParameters,
}

func init() {
	sshTmpfs.AddCommand(sshTmpfsAll)
}

func sshTmpfsAllRun(cmd *cobra.Command, args []string) {
	if parameters.Verbose {
		fmt.Fprintln(os.Stderr, utilities.LocalParameters())
	}

	err := remote.SSH(remote.SSHTmpfsExecuteGoEnum)
	if err != nil {
		color.Fprintln(color.Red, os.Stderr, err)
	}
}
func sshTmpfsAllParameters(cmd *cobra.Command, args []string) {
	sshTmpfsParameters(cmd, args)
	parameters.Commands = append(parameters.Commands, "all")

}
