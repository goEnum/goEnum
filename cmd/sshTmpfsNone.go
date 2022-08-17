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

var sshTmpfsNone = &cobra.Command{
	Use:              "none",
	Short:            "runs no modules",
	Long:             "runs no modules. Gives quick fingerprinting output for system",
	Args:             cobra.NoArgs,
	Run:              sshTmpfsNoneRun,
	PersistentPreRun: sshTmpfsNoneParameters,
}

func init() {
	sshTmpfs.AddCommand(sshTmpfsNone)
}

func sshTmpfsNoneRun(cmd *cobra.Command, args []string) {
	if parameters.Verbose {
		fmt.Fprintln(os.Stderr, utilities.LocalParameters())
	}

	err := remote.SSH(remote.SSHTmpfsExecuteGoEnum)
	if err != nil {
		color.Fprintln(color.Red, os.Stderr, err)
	}
}
func sshTmpfsNoneParameters(cmd *cobra.Command, args []string) {
	sshTmpfsParameters(cmd, args)
	parameters.Commands = append(parameters.Commands, "none")

}
