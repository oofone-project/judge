package tasks_test

import (
	"encoding/json"
	"testing"

	"github.com/oofone-project/judge/judges"
	"github.com/oofone-project/judge/judges/langs"
	"github.com/oofone-project/judge/tasks"
	"github.com/oofone-project/judge/test"
	"github.com/oofone-project/judge/utils"
	"github.com/stretchr/testify/assert"
)

func TestTaskRun(t *testing.T) {
	langs.BASE_PATH = "../languages"

	var sub tasks.Submission
	cs := test.SubFrom("../test/solution.txt", "../test/runner.txt", "../test/testin.txt", "../test/testout.txt")
	body, err := json.Marshal(cs)
	utils.FailOnError(err, "JSON marshal error")

	err = json.Unmarshal(body, &sub)
	utils.FailOnError(err, "JSON unmarshal error")

	task := tasks.NewTask(&sub, nil)

	res, err := task.Run()
	utils.FailOnError(err, "Unable to run task")

	assert.Equal(t, res.Status, judges.PASS, "Judge did not evaluate correctly")
}
