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
)

type CCheckIgnorePattern struct {
	IsNegation bool
	Value      string
}

// Creates a new CCheckIgnorePattern struct.
//
// If the value begins with a '!', the IsNegation flag is set to true and the character is removed.
func NewCCheckIgnorePattern(value string) (*CCheckIgnorePattern, error) {
	if len(value) == 1 && value[0] == '!' {
		return nil, fmt.Errorf("pattern cannot be a single '!'")
	}

	var IsNegation bool
	if len(value) > 0 && value[0] == '!' {
		IsNegation = true
		value = value[1:]
	}

	return &CCheckIgnorePattern{
		IsNegation: IsNegation,
		Value:      value,
	}, nil
}

// IsBlank returns true if the pattern contains only whitespace.
//
// A pattern that contains only whitespace is ignored.
func (pattern *CCheckIgnorePattern) IsBlank() bool {
	return strings.TrimSpace(pattern.Value) == ""
}

// IsComment returns true if the pattern beings with a '#'.
//
// A pattern that is a comment is ignored.
// A pattern that needs to evaulate against paths that being with a literal '#' must escape the character using a backslash.
//
// Example:
//
//	IsComment("#") 	(true)
//	IsComment("\#") (false)
func (pattern *CCheckIgnorePattern) IsComment() bool {
	return len(pattern.Value) > 0 && pattern.Value[0] == '#'
}
