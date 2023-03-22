package ccheck

import (
	"fmt"
	"strings"
)

type CCheckIgnorePattern struct {
	IsNegation bool
	Value      string
}

// Creates a new CCheckIgnorePattern struct.
//
// If the value begins with a '!', the IsNegation flag is set to true and the character is removed.
func NewCCheckIgnorePattern(value string) (*CCheckIgnorePattern, error) {
	if len(value) == 1 && value[0] == '!' {
		return nil, fmt.Errorf("pattern cannot be a single '!'")
	}

	var IsNegation bool
	if len(value) > 0 && value[0] == '!' {
		IsNegation = true
		value = value[1:]
	}

	return &CCheckIgnorePattern{
		IsNegation: IsNegation,
		Value:      value,
	}, nil
}

// IsBlank returns true if the pattern contains only whitespace.
//
// A pattern that contains only whitespace is ignored.
func (pattern *CCheckIgnorePattern) IsBlank() bool {
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
func (pattern *CCheckIgnorePattern) IsComment() bool {
	return len(pattern.Value) > 0 && pattern.Value[0] == '#'
}
