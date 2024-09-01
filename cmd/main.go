package main

import (
	"log"

	"github.com/oofone-project/judge/tasks"
	"github.com/oofone-project/judge/utils"
)

func main() {

	tc, err := tasks.NewTaskClient()
	utils.FailOnError(err, "Unable to init new task client")

	taskCh := make(chan tasks.Task)

	err = tc.Run(taskCh)
	utils.FailOnError(err, "Unable to run task client")
	defer tc.Close()

	for t := range taskCh {
		log.Printf("Running task %s in %s", t.GetSubmission().Id, t.GetSubmission().Language.Name)

		err = t.TaskToJudge()
		utils.FailOnError(err, "Unable to send task to judge")

		err = t.GetSubmission().Language.RunJudge()
		utils.FailOnError(err, "Unable to run judge")

		t.Ack(false)
	}
}
