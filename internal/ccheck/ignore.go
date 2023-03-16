package ccheck

import (
	"bufio"
	"bytes"
	"log"
	"strings"

	"github.com/becheran/wildmatch-go"
	"github.com/spf13/afero"
)

const CCheckIgnoreFileName string = ".ccheckignore"

type CCheckIgnore struct {
	lines []string
}

func NewCCheckIgnore(afs afero.Afero) *CCheckIgnore {
	exists, err := afs.Exists(CCheckIgnoreFileName)
	if err != nil {
		log.Panicf("failed to check file '%s', '%v'", CCheckIgnoreFileName, err)
	}
	if !exists {
		log.Printf("'%s' not found", CCheckIgnoreFileName)
		return new(CCheckIgnore)
	}
	return parse(afs)
}

func (ccheckignore CCheckIgnore) contains(path string) (bool, error) {
	var hit bool
	for _, pattern := range ccheckignore.lines {

		// A blank line matches no files.
		if len(pattern) == 0 || pattern == "" {
			continue
		}

		// A line starting with # serves as a comment.
		if pattern[0] == '#' {
			continue
		}

		// A prefix "!" negates the pattern.
		// Any matching file excluded by a previous pattern will become included again.
		var negate bool
		if pattern[0] == '!' {
			pattern = pattern[1:]
			negate = true
		}

		// If there is a separator at the beginning or middle (or both) of the
		// pattern, then the pattern is relative to the directory level of the
		// particular .ccheckignore file itself. Otherwise the pattern may also
		// match at any level below the .ccheckignore level.
		if !strings.Contains(pattern[:len(pattern)-1], "/") {
			pattern = "*" + pattern
		}

		// If a match exists, set the hit flag according to the negation flag
		pattern = pattern + "*"
		wildmatcher := wildmatch.NewWildMatch(pattern)
		if wildmatcher.IsMatch(path) {
			if negate {
				hit = false
			} else {
				hit = true
			}
		}
	}
	return hit, nil
}

func parse(afs afero.Afero) *CCheckIgnore {
	contents, err := afs.ReadFile(CCheckIgnoreFileName)
	if err != nil {
		log.Panicf("failed to read file '%s', '%v'", CCheckIgnoreFileName, err)
	}
	buffer := bytes.NewBuffer(contents)
	scanner := bufio.NewScanner(buffer)
	ccheckignore := new(CCheckIgnore)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		ccheckignore.lines = append(ccheckignore.lines, line)
	}
	return ccheckignore
}
