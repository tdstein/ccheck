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

package ccheck

import (
	"github.com/spf13/afero"
)

type CCheckApplication struct {
	linter *CCheckLinter
}

func NewCCheckApplication(path string, afs *afero.Afero) (*CCheckApplication, error) {
	scanner, err := NewCCheckScanner(path, afs)
	if err != nil {
		return nil, err
	}

	ignore, err := NewCCheckIgnore(".ccheckignore", afs)
	if err != nil {
		return nil, err
	}

	linter := NewCCheckLinter(scanner, ignore)
	return &CCheckApplication{linter: linter}, nil
}

func (app *CCheckApplication) Execute(config *CCheckConfig) error {
	validations := []CCheckValidation{
		NewCCheckValidationC000(config.copyright),
	}
	return app.linter.Lint(validations)
}
