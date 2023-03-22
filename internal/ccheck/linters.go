// Copyright (c) 2023 Taylor Steinberg

package ccheck

import (
	"fmt"
	"log"
)

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

func (linter *CCheckLinter) Lint(violations *[]CCheckViolation) error {
	for linter.scanner.HasNext() {
		file := linter.scanner.Next()
		ignored, _ := linter.ignore.Contains(file.path)
		if !ignored {
			for _, violation := range *violations {
				report, err := (violation).Evaluate(file)
				if err != nil {
					return fmt.Errorf("failed to evaulate error violation: %w", err)
				}
				if report.HasError {
					log.Println(report)
				}
			}
		}
	}
	return nil
}
