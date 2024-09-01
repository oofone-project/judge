package judges_test

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/oofone-project/judge/tasks"
	"github.com/oofone-project/judge/test"
	"github.com/oofone-project/judge/utils"
	// "github.com/stretchr/testify/assert"
)

// TODO: finish tests - run gen_out, check if files created
func TestJudge(t *testing.T) {
	b := test.NewBackend()
	defer b.Close()

	tc, err := tasks.NewTaskClient()
	utils.FailOnError(err, "Could not init task client")

	taskCh := make(chan tasks.Task)

	tc.Run(taskCh)

	sub := test.SubFrom("../test/solution.txt", "../test/runner.txt", "../test/testin.txt", "../test/testout.txt")
	b.Publish(sub)

	task := <-taskCh
	fmt.Printf("Running task %s in %s", task.GetSubmission().Id, task.GetSubmission().Language.Name)
	err = task.GetSubmission().Language.ResetJudge()
	utils.FailOnError(err, "Unable to reset judge")

	err = task.TaskToJudge(".")
	utils.FailOnError(err, "Unable to send task to judge")

	task.Ack(false)

	_, err = os.Stat("./python/submission/solution.py")
	if errors.Is(err, os.ErrNotExist) {
		t.Error("Submission not created properly")
	}

	// Put this in a separate test?
	// err = task.GetSubmission().Language.RunJudge()
	// utils.FailOnError(err, "Unable to run judge")
}
