package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = &cobra.Command{
	Use:              "version",
	Short:            "current goEnum version",
	Long:             "current goEnum version",
	Args:             cobra.NoArgs,
	Run:              versionRun,
	PersistentPreRun: versionParameters,
}

var Version string = "v1.0.0"

func init() {
	rootCmd.AddCommand(version)
}

func versionRun(cmd *cobra.Command, args []string) {
	fmt.Printf("goEnum: %v\n", Version)
}
func versionParameters(cmd *cobra.Command, args []string) {
}
