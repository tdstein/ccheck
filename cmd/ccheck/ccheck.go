// MIT License

// Copyright (c) 2023 Taylor Steinberg

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"fmt"
	"os"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/tdstein/ccheck/internal/ccheck"
)

var IgnoreFile string = ".ccheckignore"
var Version string = "dev"

var command = &cobra.Command{
	Use:     "ccheck",
	Short:   "ccheck is a fast copyright linter",
	Long:    `A fast and flexible copyright linter. Complete documentation is avaiable at https://github.com/tdstein/ccheck.`,
	Version: Version,
	Run: func(cmd *cobra.Command, args []string) {
		afs := afero.Afero{Fs: afero.NewOsFs()}
		app, err := ccheck.NewCCheckApplication(".", &afs)
		if err != nil {
			panic(err)
		}

		buffer, err := afero.ReadFile(afs, ".ccheck")
		if err != nil {
			panic(err)
		}

		config := ccheck.NewCCheckConfig(buffer, "toml")
		err = app.Execute(config)
		if err != nil {
			panic(err)
		}

		os.Exit(0)
	},
}

func main() {
	err := command.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, (fmt.Sprintf("See '%s -h' for help", command.CommandPath())))
		os.Exit(1)
	}
}
