package ccheck

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestGetCCheckIgnore(t *testing.T) {
	var afs = &afero.Afero{Fs: afero.NewMemMapFs()}
	exp := new(CCheckIgnore)
	res := NewCCheckIgnore(*afs)
	assert.Equal(t, exp, res)
}
