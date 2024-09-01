package judges_test

import (
	"log"
	"testing"

	"github.com/oofone-project/judge/tasks"
	"github.com/oofone-project/judge/test"
	"github.com/oofone-project/judge/utils"
	// "github.com/stretchr/testify/assert"
)

// TODO: finish test - run gen_out, check if files created
func TestJudge(t *testing.T) {
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
	log.Printf("Running task %s in %s", task.GetSubmission().Id, task.GetSubmission().Language)

	err = task.TaskToJudge()
	utils.FailOnError(err, "Unable to send task to judge")

	err = task.GetSubmission().Language.RunJudge()
	utils.FailOnError(err, "Unable to run judge")

	task.Ack(false)
}
