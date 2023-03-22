package ccheck

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"

	"github.com/becheran/wildmatch-go"
	"github.com/spf13/afero"
)

type CCheckIgnore struct {
	path     string
	patterns []*CCheckIgnorePattern
}

// Creates a new CCheckIgnore struct.
func NewCCheckIgnore(path string, afs afero.Afero) (*CCheckIgnore, error) {
	pattern, err := ParsePatterns(path, afs)
	if err != nil {
		return nil, fmt.Errorf("failed to create ignore struct: %w", err)
	}
	return &CCheckIgnore{
		path:     path,
		patterns: pattern,
	}, nil
}

// ParsePatterns reads patterns from a filepath.
func ParsePatterns(filepath string, afs afero.Afero) ([]*CCheckIgnorePattern, error) {

	exists, err := afs.Exists(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to check if file '%s' exists: '%w'", filepath, err)
	}

	if !exists {
		return nil, fmt.Errorf("'%s' does not exists", filepath)
	}

	contents, err := afs.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file '%s': '%w'", filepath, err)
	}

	buffer := bytes.NewBuffer(contents)
	scanner := bufio.NewScanner(buffer)
	patterns := []*CCheckIgnorePattern{}
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		pattern, err := NewIgnorePattern(line)
		if err != nil {
			return nil, fmt.Errorf("failed to parse line '%s': %w", line, err)
		}
		patterns = append(patterns, pattern)
	}

	return patterns, nil
}

// Contains returns true if an Ignore struct contains a pattern matching the provided path.
func (ignore *CCheckIgnore) Contains(path string) (bool, error) {
	if path == ignore.path {
		return true, nil
	}

	var hit bool = false
	for _, pattern := range ignore.patterns {
		if pattern.IsBlank() {
			continue
		}

		if pattern.IsComment() {
			continue
		}

		wildmatcher := wildmatch.NewWildMatch(pattern.Value)
		if wildmatcher.IsMatch(path) {
			hit = !pattern.IsNegation
		}
	}
	return hit, nil
}
