package ccheck

import "github.com/spf13/afero"

type CCheckApplication struct {
	linter *CCheckLinter
}

func NewCCheckApplication(path string, afs *afero.Afero) (*CCheckApplication, error) {
	scanner, err := NewCCheckScanner(".", afs)
	if err != nil {
		return nil, err
	}

	ignore, err := NewCCheckIgnore(".ccheckignore", afs)
	if err != nil {
		return nil, err
	}

	linter := NewCCheckLinter(scanner, ignore)
	return &CCheckApplication{linter: linter}, nil
}

func (app *CCheckApplication) Execute(config *CCheckConfig) error {
	violations := &[]CCheckViolation{
		NewCCheckViolationC000(config.copyright),
	}
	return app.linter.Lint(violations)
}
