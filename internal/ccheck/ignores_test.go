package ccheck

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestNewIgnore(t *testing.T) {
	afs := afero.Afero{Fs: afero.NewMemMapFs()}
	path := ""
	exp := &CCheckIgnore{path: path, patterns: []*CCheckIgnorePattern{}}
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
