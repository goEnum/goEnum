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

var sshTmpfs = &cobra.Command{
	Use:              "tmpfs",
	Short:            "execute goEnum over ssh with in memory exection using tmpfs",
	Long:             "execute goEnum over ssh with in memory exection using tmpfs (require remote root priviledges)",
	Run:              sshTmpfsRun,
	PersistentPreRun: sshTmpfsParameters,
}

func init() {
	ssh.AddCommand(sshTmpfs)
}

func sshTmpfsRun(cmd *cobra.Command, args []string) {
	if parameters.Verbose {
		fmt.Fprintln(os.Stderr, utilities.LocalParameters())
	}

	err := remote.SSH(remote.SSHTmpfsExecuteGoEnum)
	if err != nil {
		color.Fprintln(color.Red, os.Stderr, err)
	}
}

func sshTmpfsParameters(cmd *cobra.Command, args []string) {
	sshParameters(cmd, args)
	parameters.Commands = append(parameters.Commands, "tmpfs")

}
