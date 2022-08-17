package cmd

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/goEnum/goEnum/utilities"
	"github.com/goEnum/goEnum/utilities/color"
	"github.com/goEnum/goEnum/utilities/parameters"
	"github.com/goEnum/goEnum/utilities/remote"
	"github.com/spf13/cobra"
)

var ssh = &cobra.Command{
	Use:              "ssh",
	Short:            "execute goEnum over ssh remote connection",
	Long:             "execute goEnum over ssh remote connection",
	Args:             cobra.ArbitraryArgs,
	Run:              sshRun,
	PersistentPreRun: sshParameters,
}

func init() {
	ssh.PersistentFlags().StringP(
		"host",
		"d",
		"",
		"destination host for remote connection",
	)
	ssh.MarkPersistentFlagRequired("host")

	ssh.PersistentFlags().IntP(
		"port",
		"p",
		0,
		"port for remote connections",
	)

	ssh.PersistentFlags().StringP(
		"user",
		"u",
		"",
		"username for remote connection",
	)

	ssh.PersistentFlags().StringP(
		"pass",
		"P",
		"",
		"password for remote connection",
	)

	ssh.PersistentFlags().StringP(
		"os",
		"O",
		"",
		"destination OS for remote connection",
	)

	rootCmd.AddCommand(ssh)
}

func sshRun(cmd *cobra.Command, args []string) {
	if parameters.Verbose {
		fmt.Fprintln(os.Stderr, utilities.LocalParameters())
	}

	err := remote.SSH(remote.SSHExecuteGoEnum)
	if err != nil {
		color.Fprintln(color.Red, os.Stderr, err)
	}
}

func sshParameters(cmd *cobra.Command, args []string) {
	if !(sshPort(cmd) &&
		sshHost(cmd) &&
		sshUser(cmd) &&
		sshPassword(cmd) &&
		sshOS(cmd)) {
		os.Exit(1)
	}

	rootParameters(cmd, args)
	parameters.Commands = append(parameters.Commands, "ssh")
}

func sshPort(cmd *cobra.Command) bool {
	correctUsage := true

	value, err := cmd.Flags().GetInt("port")
	if err != nil {
		correctUsage = false
		color.Fprintln(color.Red, os.Stderr, "Error: \"port\" parameter failure =>", value)
	}

	if value == 0 {
		value = 22
	}

	if 0 > value || value > 65536 {
		color.Fprintln(color.Red, os.Stderr, "Error: invalid port number =>", value)
	}

	parameters.Port = value

	return correctUsage
}

func sshHost(cmd *cobra.Command) bool {
	correctUsage := true

	value, err := cmd.Flags().GetString("host")
	if err != nil {
		correctUsage = false
		color.Fprintln(color.Red, os.Stderr, "Error: \"host\" parameter failure =>", value)
	}

	if strings.Contains(value, "@") {
		split := strings.Split(value, "@")

		if len(split) != 2 {
			correctUsage = false
			color.Fprintln(color.Red, os.Stderr, "Error: invalid host =>", value)
			color.Fprintln(color.Red, os.Stderr, "\t> Valid Usage: username@host || host", value)
		} else {
			parameters.Username = split[0]
			value = split[1]
		}
	}

	parameters.Host = value

	return correctUsage
}

func sshUser(cmd *cobra.Command) bool {
	correctUsage := true

	if parameters.Username == "" {
		value, err := cmd.Flags().GetString("user")
		if err != nil {
			correctUsage = false
			color.Fprintln(color.Red, os.Stderr, "Error: \"user\" parameter failure =>", value)
		}

		if value == "" {
			correctUsage = false
			color.Fprintln(color.Red, os.Stderr, "Error: required flag(s) \"user\" not set")
		} else {
			parameters.Username = value
		}
	}

	return correctUsage
}

func sshPassword(cmd *cobra.Command) bool {
	correctUsage := true

	value, err := cmd.Flags().GetString("pass")
	if err != nil {
		correctUsage = false
		color.Fprintln(color.Red, os.Stderr, "Error: \"pass\" parameter failure =>", value)
	}

	parameters.Password = value

	return correctUsage
}

func sshOS(cmd *cobra.Command) bool {
	correctUsage := true

	value, err := cmd.Flags().GetString("os")
	if err != nil {
		correctUsage = false
		color.Fprintln(color.Red, os.Stderr, "Error: \"os\" parameter failure =>", value)
	}

	if value == "" {
		value = runtime.GOOS
	}

	switch value {
	case "linux":
		parameters.Binary = "/Users/max/go/goEnum/build/goEnum_linux-amd64"
	case "darwin":
		parameters.Binary = "/Users/max/go/goEnum/build/goEnum_darwin-amd64"
	case "windows":
		parameters.Binary = "/Users/max/go/goEnum/build/goEnum_windows-amd64.exe"
	default:
		color.Fprintln(color.Red, os.Stderr, "Error: Invalid OS =>", value)
		correctUsage = false
	}

	return correctUsage
}
