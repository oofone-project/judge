package judges

type Result struct {
	Stdout string
	Stderr string
	Status JudgeCode
	Total  int
	Score  int
	Time   float64
}

type JudgeCode int

const (
	PASS JudgeCode = iota
	FAIL
	TLE
	ERR
	MEM
)
