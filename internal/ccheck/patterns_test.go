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

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	type TestCase struct {
		name  string
		value string
		exp   CCheckIgnorePattern
	}

	cases := []TestCase{
		{"Empty", "", CCheckIgnorePattern{IsNegation: false, Value: ""}},
		{"Pattern", "pattern", CCheckIgnorePattern{IsNegation: false, Value: "pattern"}},
		{"Negation", "!pattern", CCheckIgnorePattern{IsNegation: true, Value: "pattern"}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res, _ := NewCCheckIgnorePattern(tc.value)
			assert.Equal(t, tc.exp, *res)
		})
	}
}

func TestNewErrors(t *testing.T) {
	type TestCase struct {
		name  string
		value string
	}

	cases := []TestCase{
		{"Exclamation", "!"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			pattern, err := NewCCheckIgnorePattern(tc.value)
			assert.Nil(t, pattern)
			assert.Error(t, err)
		})
	}
}

func TestIsBlank(t *testing.T) {
	type TestCase struct {
		name  string
		value string
		exp   bool
	}

	cases := []TestCase{
		{"Empty", "", true},
		{"Space", " ", true},
		{"Feed", "\f", true},
		{"NewLine", "\n", true},
		{"CarrigeReturn", "\r", true},
		{"Tab", "\t", true},
		{"Alpha", "a", false},
		{"Numeric", "1", false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			pattern := CCheckIgnorePattern{Value: tc.value}
			res := pattern.IsBlank()
			assert.Equal(t, tc.exp, res)
		})
	}
}

func TestIsComment(t *testing.T) {
	type TestCase struct {
		name  string
		value string
		exp   bool
	}

	cases := []TestCase{
		{"Empty", "", false},
		{"Number", "#", true},
		{"Escape", "\\#", false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			pattern := CCheckIgnorePattern{Value: tc.value}
			res := pattern.IsComment()
			assert.Equal(t, tc.exp, res)
		})
	}
}

func TestIsNegation(t *testing.T) {
	type TestCase struct {
		name  string
		value string
		exp   bool
	}

	cases := []TestCase{
		{"Empty", "", false},
		{"Pattern", "pattern", false},
		{"PatternNegation", "!pattern", true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			pattern, _ := NewCCheckIgnorePattern(tc.value)
			res := pattern.IsNegation
			assert.Equal(t, tc.exp, res)
		})
	}
}
