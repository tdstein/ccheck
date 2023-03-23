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
	"fmt"
	"log"
)

type CCheckLinter struct {
	scanner *CCheckScanner
	ignore  *CCheckIgnore
}

func NewCCheckLinter(scanner *CCheckScanner, ignore *CCheckIgnore) *CCheckLinter {
	return &CCheckLinter{
		scanner: scanner,
		ignore:  ignore,
	}
}

func (linter *CCheckLinter) Lint(validations []CCheckValidation) error {
	for linter.scanner.HasNext() {
		file := linter.scanner.Next()
		ignored, _ := linter.ignore.Contains(file.path)
		if !ignored {
			for _, validation := range validations {
				report, err := validation.Evaluate(file)
				if err != nil {
					return fmt.Errorf("failed to evaulate error validation: %w", err)
				}
				if report.HasError {
					log.Println(report)
				}
			}
		}
	}
	return nil
}
