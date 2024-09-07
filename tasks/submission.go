package tasks

import (
	"github.com/oofone-project/judge/judges/langs"
)

type Submission struct {
	Language *langs.Language `json:"language"`
	Solution []byte          `json:"solution"`
	Runner   []byte          `json:"runner"`
	TestIn   []byte          `json:"testin"`
	TestOut  []byte          `json:"testout"`
	Id       string          `json:"id"`
}
