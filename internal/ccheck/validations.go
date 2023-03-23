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
	"strings"

	"github.com/tdstein/ccheck/internal/ccheck/copyrights"
)

type CCheckValidationReport struct {
	Code            string
	HasError        bool
	Path            string
	LineNumberStart int
	LineNumberEnd   int
	Error           error
}

func (report *CCheckValidationReport) String() string {
	return fmt.Sprintf("%s:%d:%d: %s %v", report.Path, report.LineNumberStart, report.LineNumberEnd, report.Code, report.Error)
}

type CCheckValidation interface {
	Evaluate(file *CCheckFile) (*CCheckValidationReport, error)
}

type CCheckValidationC000 struct {
	copyright *copyrights.CCheckCopyright
}

func NewCCheckValidationC000(copyright *copyrights.CCheckCopyright) *CCheckValidationC000 {
	return &CCheckValidationC000{
		copyright: copyright,
	}
}

func (validation *CCheckValidationC000) Evaluate(file *CCheckFile) (*CCheckValidationReport, error) {
	contents, err := file.afs.ReadFile(file.path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file '%s': '%w'", file.path, err)
	}

	exists := true
	text := string(contents)
	if len(text) < len(validation.copyright.Text) {
		exists = false
	}

	for _, line := range validation.copyright.Text {
		if exists && !strings.Contains(text, line) {
			exists = false
		}
	}

	return &CCheckValidationReport{
		Code:     "C000",
		HasError: !exists,
		Path:     file.path,
		Error:    fmt.Errorf("missing copyright"),
	}, nil
}
