package tasks_test

import (
	"testing"

	"github.com/oofone-project/judge/judges/langs"
	"github.com/oofone-project/judge/tasks"
	"github.com/oofone-project/judge/test"
	"github.com/oofone-project/judge/utils"
	"github.com/stretchr/testify/assert"
)

// TODO: multiple clients and concurrent tasks to test fair dispatch? need tasks that take time
func TestClientRun(t *testing.T) {
	b := test.NewBackend()
	defer b.Close()

	tc, err := tasks.NewTaskClient()
	if err != nil {
		utils.FailOnError(err, "Could not init task client")
	}

	resCh := make(chan *tasks.ResultAndSubmission)

	tc.Run(resCh)

	sub := test.SubFrom("../test/solution.txt", "../test/runner.txt", "../test/testin.txt", "../test/testout.txt")
	b.Publish(sub)

	res := <-resCh

	assert.Equal(t, res.Submission.Id, sub.Id, "Task ids not equal")
	assert.Equal(t, res.Submission.Language.Name, langs.Languages["python"].Name, "Langs not equal")
}
