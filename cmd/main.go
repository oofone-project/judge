package main

import (
	"github.com/oofone-project/judge/consumer"
	"github.com/oofone-project/judge/utils"
)

func main() {

	var forever chan struct{}

	tc, err := consumer.NewTaskClient()
	utils.FailOnError(err, "Unable to init new task client")

	err = tc.Run()
	utils.FailOnError(err, "Unable to run task client")
	defer tc.Close()

	<-forever
}
