package ccheck

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/spf13/afero"
)

var filename string = ".ccheckignore"

type CCheckIgnore struct {
	contents []byte
}

func walk(path string, info os.FileInfo, err error) error {
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(path)
	return nil
}

func (ccheckignore CCheckIgnore) PrintEachFile(afs afero.Afero) (*CCheckIgnore, error) {
	buffer := bytes.NewBuffer(ccheckignore.contents)
	scanner := bufio.NewScanner(buffer)
	for scanner.Scan() {
		line := scanner.Text()
		matches, err := afero.Glob(afs, line)
		if err != nil {
			log.Panicf("failed to glob '%v'", err)
		}
		for _, match := range matches {
			afero.Walk(afs, match, walk)
		}

	}
	return &ccheckignore, nil
}

func GetCCheckIgnore(afs afero.Afero) (*CCheckIgnore, error) {
	exists, _ := afs.Exists(filename)
	if !exists {
		log.Printf("'%s' not found", filename)
		contents := []byte{}
		return &CCheckIgnore{contents: contents}, nil
	}
	contents, err := afs.ReadFile(filename)
	if err != nil {
		log.Panicf("failed to read file '%v'", err)
	}
	return &CCheckIgnore{contents: contents}, nil
}
