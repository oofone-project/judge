package main

import (
	"log"

	"github.com/oofone-project/judge/consumer"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {

	var forever chan struct{}

	tc, err := consumer.NewTaskClient()
	failOnError(err, "Unable to init new task client")

	err = tc.Run()
	failOnError(err, "Unable to run task client")

	<-forever
}
