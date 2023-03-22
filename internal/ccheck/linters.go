package ccheck

import "fmt"

type CCheckLinter struct {
	scanner *CCheckScanner
	ignore  *CCheckIgnore
}

func NewCCheckLinter(scanner *CCheckScanner, ignore *CCheckIgnore) *CCheckLinter {
	return &CCheckLinter{
		scanner: scanner,
		ignore:  ignore,
	}
}

func (linter *CCheckLinter) Lint(config *CCheckConfig) error {
	for linter.scanner.HasNext() {
		file := linter.scanner.Next()
		ignored, _ := linter.ignore.Contains(file.path)
		if !ignored {
			if !file.HasCopyright(config.copyright) {
				fmt.Println(file)
			}
		}
	}
	return nil
}
