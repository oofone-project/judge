package model

import "github.com/oofone-project/judge/judges"

// TODO: language should just be string
// client is just sending the lang, not the struct
// use map of strings to struct or something else
type Submission struct {
	Language judges.Language `json:"language"`
	Solution []byte          `json:"solution"`
	Runner   []byte          `json:"runner"`
	TestIn   []byte          `json:"testin"`
	TestOut  []byte          `json:"testout"`
	Id       string          `json:"id"`
}
