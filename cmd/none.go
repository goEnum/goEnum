package cmd

import (
	"github.com/goEnum/goEnum/utilities/parameters"
	"github.com/spf13/cobra"
)

var none = &cobra.Command{
	Use:              "none",
	Short:            "runs no modules",
	Long:             "runs no modules. Gives quick fingerprinting output for system",
	Args:             cobra.NoArgs,
	Run:              noneRun,
	PersistentPreRun: noneParameters,
}

func init() {
	rootCmd.AddCommand(none)
}

func noneRun(cmd *cobra.Command, args []string) {
	parameters.Verbose = true
	verboseInformation()
}
func noneParameters(cmd *cobra.Command, args []string) {
	rootParameters(cmd, args)
	parameters.Commands = append(parameters.Commands, "none")

}
