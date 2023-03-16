package ccheck

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"

	"github.com/spf13/afero"
)

func walk(afs afero.Afero) filepath.WalkFunc {
	ccheckignore := NewCCheckIgnore(afs)
	return func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			ignores, err := ccheckignore.contains(path)
			if err != nil {
				if err != nil {
					log.Panicf("failed to process ignore for path '%s', '%v'", path, err)
				}
			}
			if ignores {
				// check for copyright
				fmt.Println(path)
			}
		}
		return nil
	}
}

func CCheck() error {
	var afs = &afero.Afero{Fs: afero.NewOsFs()}
	walk := walk(*afs)
	afs.Walk(".", walk)
	return nil
}
