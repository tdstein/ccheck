package ccheck

import "fmt"

type CCheckViolation interface {
	Evaluate(file *CCheckFile) (*CCheckViolationReport, error)
}

type CCheckViolationReport struct {
	Code            string
	HasError        bool
	Path            string
	LineNumberStart int
	LineNumberEnd   int
	Error           error
}

func (report *CCheckViolationReport) String() string {
	return fmt.Sprintf("%s:%d:%d: %s %v", report.Path, report.LineNumberStart, report.LineNumberEnd, report.Code, report.Error)
}

type CCheckViolationC000 struct {
	copyright *CCheckCopyright
}

func NewCCheckViolationC000(copyright *CCheckCopyright) *CCheckViolationC000 {
	return &CCheckViolationC000{
		copyright: copyright,
	}
}

func (violation *CCheckViolationC000) Evaluate(file *CCheckFile) (*CCheckViolationReport, error) {
	exists, err := file.HasCopyright(violation.copyright)
	if err != nil {
		return nil, fmt.Errorf("violation failed: %w", err)
	}
	return &CCheckViolationReport{
		Code:     "C000",
		HasError: !exists,
		Path:     file.path,
		Error:    fmt.Errorf("missing copyright"),
	}, nil
}
