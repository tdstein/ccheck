package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tdstein/ccheck/internal/ccheck"
)

var IgnoreFile string = ".ccheckignore"
var Version string = "local"

var command = &cobra.Command{
	Use:     "ccheck",
	Short:   "ccheck is a fast copyright linter",
	Long:    `A fast and flexible copyright linter. Complete documentation is avaiable at https://github.com/tdstein/ccheck.`,
	Version: Version,
	Run: func(cmd *cobra.Command, args []string) {
		// config := ccheck.GetCCheckConfig()
		ccheck.CCheck()
	},
}

func main() {
	err := command.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, (fmt.Sprintf("See '%s -h' for help", command.CommandPath())))
		os.Exit(1)
	}
}
