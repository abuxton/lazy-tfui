// Copyright Authors of lazy-tfui

package cmd

import (
	"fmt"

	"github.com/abuxton/lazy-tfui/internal/colour"
	"github.com/spf13/cobra"
)

func versionCmd() *cobra.Command {
	var short bool

	command := cobra.Command{
		Use:   "version",
		Short: "Print version/build info",
		Long:  "Print version/build information",
		Run: func(cmd *cobra.Command, args []string) {
			printVersion(short)
		},
	}

	command.PersistentFlags().BoolVarP(&short, "short", "s", false, "Prints TFUI version info in short format")

	return &command
}

func printVersion(short bool) {
	const fmat = "%-20s %s\n"
	var outputcolour colour.Paint

	if short {
		outputcolour = -1
	} else {
		outputcolour = colour.Cyan
		printLogo(outputcolour)
	}
	printTuple(fmat, "Version", version, outputcolour)
	printTuple(fmat, "Commit", commit, outputcolour)
	printTuple(fmat, "Date", date, outputcolour)
}

func printTuple(fmat, section, value string, outputcolour colour.Paint) {
	if outputcolour != -1 {
		fmt.Fprintf(out, fmat, colour.colourize(section+":", outputcolour), value)
		return
	}
	fmt.Fprintf(out, fmat, section, value)
}
