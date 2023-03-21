package ccheck

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"

	"github.com/spf13/afero"
)

type Service struct {
	afs afero.Afero
}

func (service Service) GetFiles() ([]string, error) {
	ignore, err := NewIgnore(".ccheckignore", service.afs)
	if err != nil {
		return nil, fmt.Errorf("failed to create ignore struct: %w", err)
	}
	function := GetWalkFunction(ignore)
	service.afs.Walk(".", function)
	return nil, nil
}

func GetWalkFunction(ignore *Ignore) filepath.WalkFunc {
	return func(path string, info fs.FileInfo, _ error) error {
		if info.IsDir() {
			return nil
		}
		contains, err := ignore.Contains(path)
		if err != nil {
			return fmt.Errorf("failed to determine if '%s' is contained in the ignorefile: %w", path, err)
		}
		if contains {
			log.Printf("'%s' matches file at path '%s'", ignore.path, path)
			return nil
		}
		return nil
	}
}
