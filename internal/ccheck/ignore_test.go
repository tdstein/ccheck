package ccheck

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestNewIgnore(t *testing.T) {
	afs := afero.Afero{Fs: afero.NewMemMapFs()}
	path := ""
	exp := &Ignore{path: path, patterns: []*IgnorePattern{}}
	res, _ := NewIgnore(path, afs)
	assert.Equal(t, exp, res)
}

func TestGetPatterns(t *testing.T) {
	afs := afero.Afero{Fs: afero.NewMemMapFs()}
	path := ".ccheckignore"
	res, _ := ParsePatterns(path, afs)
	assert.Nil(t, res)
}

func TestNewIgnorePattern(t *testing.T) {
	type TestCase struct {
		name  string
		value string
		exp   IgnorePattern
	}

	cases := []TestCase{
		{"Empty", "", IgnorePattern{IsNegation: false, Value: ""}},
		{"Pattern", "pattern", IgnorePattern{IsNegation: false, Value: "pattern"}},
		{"Negation", "!pattern", IgnorePattern{IsNegation: true, Value: "pattern"}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res, _ := NewIgnorePattern(tc.value)
			assert.Equal(t, tc.exp, *res)
		})
	}
}

func TestNewIgnorePattern_Errors(t *testing.T) {
	type TestCase struct {
		name  string
		value string
	}

	cases := []TestCase{
		{"Exclamation", "!"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			pattern, err := NewIgnorePattern(tc.value)
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
			pattern := IgnorePattern{Value: tc.value}
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
			pattern := IgnorePattern{Value: tc.value}
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
			pattern, _ := NewIgnorePattern(tc.value)
			res := pattern.IsNegation
			assert.Equal(t, tc.exp, res)
		})
	}
}
