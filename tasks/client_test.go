package tasks_test

import (
	"testing"

	"github.com/oofone-project/judge/tasks"
	"github.com/oofone-project/judge/test"
	"github.com/oofone-project/judge/utils"
	"github.com/stretchr/testify/assert"
)

// TODO: multiple clients and concurrent tasks to test fair dispatch? need tasks that take time
func TestRun(t *testing.T) {
	b := test.NewBackend()
	defer b.Close()

	tc, err := tasks.NewTaskClient()
	if err != nil {
		utils.FailOnError(err, "Could not init task client")
	}

	taskCh := make(chan tasks.Task)

	tc.Run(taskCh)

	sub := test.SubFrom("../test/solution.txt", "../test/runner.txt", "../test/testin.txt", "../test/testout.txt")
	b.Publish(sub)

	task := <-taskCh
	assert.Equal(t, task.GetSubmission().Id, sub.Id, "Task ids not equal")
}
