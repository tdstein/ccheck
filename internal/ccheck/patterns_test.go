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
			res, _ := NewIgnorePattern(tc.value)
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
			pattern, _ := NewIgnorePattern(tc.value)
			res := pattern.IsNegation
			assert.Equal(t, tc.exp, res)
		})
	}
}
