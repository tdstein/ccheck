package ccheck

import (
	"fmt"
	"io/fs"

	"github.com/spf13/afero"
)

type CCheckScanner struct {
	files *[]CCheckFile
	size  int
	pos   int
}

func NewCCheckScanner(path string, afs afero.Afero) (*CCheckScanner, error) {
	files, err := GetFiles(path, afs)
	if err != nil {
		return nil, err
	}
	return &CCheckScanner{
		files: files,
		size:  len(*files),
		pos:   0,
	}, nil
}

func (scanner *CCheckScanner) HasNext() bool {
	return scanner.pos < scanner.size
}

func (scanner *CCheckScanner) Next() CCheckFile {
	cur := (*scanner.files)[scanner.pos]
	scanner.pos += 1
	return cur
}

func GetFiles(path string, afs afero.Afero) (*[]CCheckFile, error) {
	files := new([]CCheckFile)
	err := afs.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			file := NewCCheckFile(path)
			*files = append(*files, *file)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to walk filesystem at path '%s': %w", path, err)
	}
	return files, nil
}
