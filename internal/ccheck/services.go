package ccheck

import (
	"io/fs"
	"log"

	"github.com/spf13/afero"
)

type Service struct {
	afs afero.Afero
}

func (service Service) GetFiles() ([]string, error) {
	ignore := NewIgnore(".ccheckignore", service.afs)
	service.afs.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if ignore.Contains(path) {
			log.Printf("'%s' matches file at path '%s'", ignore.path, path)
			return nil
		}
		return nil
	})
	return []string{}, nil
}
