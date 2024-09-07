package judges

import (
	"errors"

	"github.com/oofone-project/judge/judges/langs"
)

var (
	KillError = errors.New("failed to kill subprocess")
)

type Judge struct {
	Lang *langs.Language
}

func NewJudge(l *langs.Language) *Judge {
	return &Judge{
		Lang: l,
	}
}

// Generate output from user solution
// Collect stats like time to run, error code, memory
func (j Judge) RunJudge() (string, error) {
	return "", nil
}

// Grade solution output
// Read stderr, stdout
func (j Judge) Evaluate() (*Result, error) {
	return &Result{}, nil
}
