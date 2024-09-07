package main

import (
	"fmt"
	"log"

	"github.com/oofone-project/judge/tasks"
	"github.com/oofone-project/judge/utils"
)

func main() {

	tc, err := tasks.NewTaskClient()
	utils.FailOnError(err, "Unable to init new task client")

	taskCh := make(chan *tasks.Task)

	err = tc.Run(taskCh)
	utils.FailOnError(err, "Unable to run task client")
	defer tc.Close()

	for t := range taskCh {
		log.Printf("Running task %s in %s", t.GetSubmission().Id, t.GetSubmission().Language.Name)

		res, err := t.Run()
		utils.FailOnError(err, "Unable to run task")

		// TODO: Send result to backend
		fmt.Printf("Time: %f\nScore: %d/%d", res.Time, res.Score, res.Total)

		t.Ack(false)
	}
}
