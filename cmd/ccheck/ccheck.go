package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/afero"
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
		afs := afero.Afero{Fs: afero.NewOsFs()}
		application := ccheck.NewApplication(afs, ".")
		exit, err := application.Execute()
		if err != nil {
			log.Panic(err)
		}
		os.Exit(exit)
	},
}

func main() {
	err := command.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, (fmt.Sprintf("See '%s -h' for help", command.CommandPath())))
		os.Exit(1)
	}
}
