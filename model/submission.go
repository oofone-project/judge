package model

import "github.com/oofone-project/judge/judges"

type Submission struct {
	Language judges.Language `json:"language"`
	Solution []byte          `json:"solution"`
	Runner   []byte          `json:"runner"`
	TestIn   []byte          `json:"testin"`
	TestOut  []byte          `json:"testout"`
	Id       string          `json:"id"`
}
