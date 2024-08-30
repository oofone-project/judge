package model

type Submission struct {
	Solution []byte `json:"solution"`
	Runner   []byte `json:"runner"`
	TestIn   []byte `json:"testin"`
	TestOut  []byte `json:"testout"`
	Id       string `json:"id"`
}
