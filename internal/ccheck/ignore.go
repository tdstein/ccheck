package ccheck

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"

	"github.com/becheran/wildmatch-go"
	"github.com/spf13/afero"
)

type Ignore struct {
	path     string
	patterns []*IgnorePattern
}

// Creates a new Ignore struct.
func NewIgnore(path string, afs afero.Afero) (*Ignore, error) {
	pattern, err := ParsePatterns(path, afs)
	if err != nil {
		return nil, fmt.Errorf("failed to create ignore struct: %w", err)
	}
	return &Ignore{
		path:     path,
		patterns: pattern,
	}, nil
}

// Contains returns true if an Ignore struct contains a pattern matching the provided path.
func (ignore *Ignore) Contains(path string) (bool, error) {
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

type IgnorePattern struct {
	IsNegation bool
	Value      string
}

// Creates a new IgnorePattern struct.
//
// If the value begins with a '!', the IsNegation flag is set to true and the character is removed.
func NewIgnorePattern(value string) (*IgnorePattern, error) {
	if len(value) == 1 && value[0] == '!' {
		return nil, fmt.Errorf("pattern cannot be a single '!'")
	}

	var IsNegation bool
	if len(value) > 0 && value[0] == '!' {
		IsNegation = true
		value = value[1:]
	}

	return &IgnorePattern{
		IsNegation: IsNegation,
		Value:      value,
	}, nil
}

// ParsePatterns reads patterns from a filepath.
func ParsePatterns(filepath string, afs afero.Afero) ([]*IgnorePattern, error) {

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
	patterns := []*IgnorePattern{}
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

// IsBlank returns true if the pattern contains only whitespace.
//
// A pattern that contains only whitespace is ignored.
func (pattern *IgnorePattern) IsBlank() bool {
	return strings.TrimSpace(pattern.Value) == ""
}

// IsComment returns true if the pattern beings with a '#'.
//
// A pattern that is a comment is ignored.
// A pattern that needs to evaulate against paths that being with a literal '#' must escape the character using a backslash.
//
// Example:
//
//	IsComment("#") 	(true)
//	IsComment("\#") (false)
func (pattern *IgnorePattern) IsComment() bool {
	return len(pattern.Value) > 0 && pattern.Value[0] == '#'
}
