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

var sshNone = &cobra.Command{
	Use:              "none",
	Short:            "runs no modules",
	Long:             "runs no modules. Gives quick fingerprinting output for system",
	Args:             cobra.NoArgs,
	Run:              sshNoneRun,
	PersistentPreRun: sshNoneParameters,
}

func init() {
	ssh.AddCommand(sshNone)
}

func sshNoneRun(cmd *cobra.Command, args []string) {
	if parameters.Verbose {
		fmt.Fprintln(os.Stderr, utilities.LocalParameters())
	}

	err := remote.SSH(remote.SSHExecuteGoEnum)
	if err != nil {
		color.Fprintln(color.Red, os.Stderr, err)
	}
}
func sshNoneParameters(cmd *cobra.Command, args []string) {
	sshParameters(cmd, args)
	parameters.Commands = append(parameters.Commands, "none")

}
