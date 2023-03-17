package ccheck

import (
	"bufio"
	"bytes"
	"log"
	"path/filepath"
	"strings"

	"github.com/becheran/wildmatch-go"
	"github.com/spf13/afero"
)

type Ignore struct {
	path     string
	patterns []string
}

func NewIgnore(path string, afs afero.Afero) *Ignore {
	return &Ignore{
		path:     path,
		patterns: GetPatterns(path, afs),
	}
}

func (ignore Ignore) Contains(path string) bool {
	path = filepath.Join("/", path)
	var hit bool
	for _, pattern := range ignore.patterns {
		// A blank line matches nothing
		if len(pattern) == 0 || pattern == "" {
			continue
		}

		// A line starting with # is as a comment.
		if pattern[0] == '#' {
			continue
		}

		// A prefix "!" negates the pattern.
		// Any matching file excluded by a previous pattern will become included again.
		var negate bool
		if pattern[0] == '!' {
			// A single '!' matches nothing
			if len(pattern) == 1 {
				continue
			}
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

		if !strings.HasPrefix(pattern, "**") {
			pattern = filepath.Join("/", pattern)
		}

		if !strings.HasSuffix(pattern, "**") {
			pattern = pattern + "*"
		}

		// If a match exists, set the hit flag according to the negation flag
		wildmatcher := wildmatch.NewWildMatch(pattern)
		if wildmatcher.IsMatch(path) {
			if negate {
				hit = false
			} else {
				hit = true
			}
		}
	}
	return hit
}

func GetPatterns(path string, afs afero.Afero) []string {

	exists, err := afs.Exists(path)
	if err != nil {
		log.Panicf("failed to check file '%s', '%v'", path, err)
	}

	if !exists {
		log.Printf("'%s' not found", path)
		return []string{}
	}

	contents, err := afs.ReadFile(path)
	if err != nil {
		log.Panicf("failed to read file '%s', '%v'", path, err)
	}

	buffer := bytes.NewBuffer(contents)
	scanner := bufio.NewScanner(buffer)
	patterns := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		patterns = append(patterns, line)
	}

	return patterns
}
