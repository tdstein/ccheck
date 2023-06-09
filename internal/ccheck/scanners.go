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
	"io/fs"

	"github.com/spf13/afero"
)

type CCheckScanner struct {
	files []CCheckFile
	size  int
	pos   int
}

func NewCCheckScanner(path string, afs *afero.Afero) (*CCheckScanner, error) {
	files, err := GetFiles(path, afs)
	if err != nil {
		return nil, err
	}
	return &CCheckScanner{
		files: files,
		size:  len(files),
		pos:   0,
	}, nil
}

func (scanner *CCheckScanner) HasNext() bool {
	return scanner.pos < scanner.size
}

func (scanner *CCheckScanner) Next() *CCheckFile {
	cur := scanner.files[scanner.pos]
	scanner.pos += 1
	return &cur
}

func GetFiles(path string, afs *afero.Afero) ([]CCheckFile, error) {
	files := []CCheckFile{}
	err := afs.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			file := NewCCheckFile(path, afs)
			files = append(files, *file)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to walk filesystem at path '%s': %w", path, err)
	}
	return files, nil
}
