package ccheck

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestGetCCheckIgnore(t *testing.T) {
	var afs = &afero.Afero{Fs: afero.NewMemMapFs()}
	contents := []byte("ccheckignore")
	afs.WriteFile(".ccheckignore", contents, 0644)
	exp := &CCheckIgnore{contents: contents}
	res, _ := GetCCheckIgnore(*afs)
	assert.Equal(t, exp, res)
}

func TestGetCCheckIgnore_Error(t *testing.T) {
	var afs = &afero.Afero{Fs: afero.NewMemMapFs()}
	_, err := GetCCheckIgnore(*afs)
	assert.NotNil(t, err)
}
