package ccheck

import (
	"path/filepath"
)

type CCheckFile struct {
	path string
}

func NewCCheckFile(path string) *CCheckFile {
	return &CCheckFile{
		path: path,
	}
}

func (file *CCheckFile) HasCopyright(copyright CCheckCopyright) bool {
	ext := filepath.Ext(file.path)
	// route to service
	switch ext {
	case ".go":
		return false
	default:
		return false
	}
}
