package ccheck

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"

	"github.com/spf13/afero"
)

type CCheckFile struct {
	afs  *afero.Afero
	path string
}

func NewCCheckFile(path string, afs *afero.Afero) *CCheckFile {
	return &CCheckFile{
		afs:  afs,
		path: path,
	}
}

func (file *CCheckFile) HasCopyright(copyright *CCheckCopyright) (bool, error) {
	contents, err := file.afs.ReadFile(file.path)
	if err != nil {
		return false, fmt.Errorf("failed to read file '%s': '%w'", file.path, err)
	}

	buffer := bytes.NewBuffer(contents)
	scanner := bufio.NewScanner(buffer)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, (*copyright.text)[0]) {
			return true, nil
		}
	}
	return false, nil
}
