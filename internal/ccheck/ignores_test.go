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
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestNewIgnore(t *testing.T) {
	afs := afero.Afero{Fs: afero.NewMemMapFs()}
	path := ""
	exp := &CCheckIgnore{path: path, patterns: []CCheckIgnorePattern{}}
	res, _ := NewCCheckIgnore(path, &afs)
	assert.Equal(t, exp, res)
}

func TestContainsSelf(t *testing.T) {
	filepath := ".ccheckignore"
	afs := afero.Afero{Fs: afero.NewMemMapFs()}
	afs.Create(filepath)
	ignore, _ := NewCCheckIgnore(filepath, &afs)
	res, _ := ignore.Contains(filepath)
	assert.True(t, res)
}

func TestGetPatterns(t *testing.T) {
	afs := afero.Afero{Fs: afero.NewMemMapFs()}
	path := ".ccheckignore"
	res, _ := ParsePatterns(path, &afs)
	assert.Nil(t, res)
}
