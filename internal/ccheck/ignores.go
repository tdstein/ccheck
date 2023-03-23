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
	"bufio"
	"bytes"
	"fmt"
	"strings"

	"github.com/becheran/wildmatch-go"
	"github.com/spf13/afero"
)

type CCheckIgnore struct {
	path     string
	patterns []CCheckIgnorePattern
}

// Creates a new CCheckIgnore struct.
func NewCCheckIgnore(path string, afs *afero.Afero) (*CCheckIgnore, error) {
	pattern, err := ParsePatterns(path, afs)
	if err != nil {
		return nil, fmt.Errorf("failed to create ignore struct: %w", err)
	}
	return &CCheckIgnore{
		path:     path,
		patterns: pattern,
	}, nil
}

// ParsePatterns reads patterns from a filepath.
func ParsePatterns(filepath string, afs *afero.Afero) ([]CCheckIgnorePattern, error) {

	exists, err := afs.Exists(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to check if file '%s' exists: '%w'", filepath, err)
	}

	if !exists {
		return nil, fmt.Errorf("'%s' does not exists", filepath)
	}

	contents, err := afs.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file '%s': '%w'", filepath, err)
	}

	buffer := bytes.NewBuffer(contents)
	scanner := bufio.NewScanner(buffer)
	patterns := []CCheckIgnorePattern{}
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		pattern, err := NewCCheckIgnorePattern(line)
		if err != nil {
			return nil, fmt.Errorf("failed to parse line '%s': %w", line, err)
		}
		patterns = append(patterns, *pattern)
	}

	return patterns, nil
}

// Contains returns true if an Ignore struct contains a pattern matching the provided path.
func (ignore *CCheckIgnore) Contains(path string) (bool, error) {
	if path == ignore.path {
		return true, nil
	}

	var hit bool = false
	for _, pattern := range ignore.patterns {
		if pattern.IsBlank() {
			continue
		}

		if pattern.IsComment() {
			continue
		}

		wildmatcher := wildmatch.NewWildMatch(pattern.Value)
		if wildmatcher.IsMatch(path) {
			hit = !pattern.IsNegation
		}
	}
	return hit, nil
}
