package main

import (
	"log"

	"github.com/oofone-project/judge/consumer"
	"github.com/oofone-project/judge/utils"
)

func main() {

	tc, err := consumer.NewTaskClient()
	utils.FailOnError(err, "Unable to init new task client")

	taskCh := make(chan consumer.Task)

	err = tc.Run(taskCh)
	utils.FailOnError(err, "Unable to run task client")
	defer tc.Close()

	for t := range taskCh {
		log.Print(t.GetSubmission().Id)
		t.Ack(false)
	}
}
