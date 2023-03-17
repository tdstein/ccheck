package ccheck

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestNewIgnore(t *testing.T) {
	afs := afero.Afero{Fs: afero.NewMemMapFs()}
	path := ""
	exp := &Ignore{path: path, patterns: []string{}}
	res := NewIgnore(path, afs)
	assert.Equal(t, exp, res)
}

func TestGetPatterns_FileDoesNotExist(t *testing.T) {
	afs := afero.Afero{Fs: afero.NewMemMapFs()}
	path := ".ccheckignore"
	exp := []string{}
	res := GetPatterns(path, afs)
	assert.Equal(t, exp, res)
}

func TestContains(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		patterns []string
		exp      bool
	}{
		{
			"no patterns",
			"",
			[]string{},
			false,
		},
		{
			"empty pattern",
			"empty",
			[]string{""},
			false,
		},
		{
			"comment",
			"comment",
			[]string{"#"},
			false,
		},
		{
			"empty negation",
			"empty negation",
			[]string{"!"},
			false,
		},
		{
			"file at root level",
			"/file",
			[]string{"file"},
			true,
		},
		{
			"file in any directory",
			"/directory/file",
			[]string{"directory/"},
			true,
		},
		{
			"file in any directory",
			"/parent/directory/file",
			[]string{"directory/"},
			true,
		},
		{
			"file in relative directory",
			"/directory/file",
			[]string{"/directory"},
			true,
		},
		{
			"file in relative directory",
			"/directory/file",
			[]string{"/directory/"},
			true,
		},
		{
			"file in any directory using wildcards",
			"/directory/file",
			[]string{"**/file"},
			true,
		},
		{
			"file in any directory using wildcards",
			"/directory/file",
			[]string{"**/directory/file"},
			true,
		},
		{
			"any file in directory using wildcards",
			"/directory/file",
			[]string{"directory/**"},
			true,
		},
		// TODO - Since this scenario doesn't match the gitignore documentation, check the Wildcard implementation to see if there is an bug.
		{
			"file in any subdirectory using wildcards",
			"directory/file",
			[]string{"directory/**/file"},
			false,
		},
		{
			"file in any subdirectory using wildcards",
			"directory/x/file",
			[]string{"directory/**/file"},
			true,
		},
		{
			"file in any subdirectory using wildcards",
			"directory/x/y/file",
			[]string{"directory/**/file"},
			true,
		},
		{
			"file in any subdirectory using wildcards",
			"directory/x/y/z/file",
			[]string{"directory/**/file"},
			true,
		},
		{
			"negation",
			"file",
			[]string{
				"file",
				"!file",
			},
			false,
		},
		{
			"double negation",
			"file",
			[]string{
				"file",
				"!file",
				"file",
			},
			true,
		},
		{
			"double negation",
			"file",
			[]string{
				"file",
				"!file",
				"!file",
			},
			false,
		},
		{
			"double negation",
			"file",
			[]string{
				"file",
				"!file",
				"file",
				"!file",
			},
			false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ignore := new(Ignore)
			ignore.patterns = test.patterns
			res := ignore.Contains(test.path)
			assert.Equal(t, test.exp, res)
		})
	}
}
