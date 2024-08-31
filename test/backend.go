// Simulate the backend adding tasks to the queue
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/oofone-project/judge/model"
	"github.com/oofone-project/judge/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

func subFrom(solFile string, runFile string, testIn string, testOut string) model.Submission {
	sol, err := os.ReadFile(solFile)
	utils.FailOnError(err, "Could not open file")
	run, err := os.ReadFile(runFile)
	utils.FailOnError(err, "Could not open file")
	testin, err := os.ReadFile(testIn)
	utils.FailOnError(err, "Could not open file")
	testout, err := os.ReadFile(testOut)
	utils.FailOnError(err, "Could not open file")

	id := uuid.New()

	sub := model.Submission{
		Solution: sol,
		Runner:   run,
		TestIn:   testin,
		TestOut:  testout,
		Id:       id.String(),
	}

	return sub
}

func main() {
	godotenv.Load(".env")

	conn, err := amqp.Dial(os.Getenv("RABBIT_MQ_URI"))
	utils.FailOnError(err, "Could not open file")
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError(err, "Could not open file")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		os.Getenv("RABBIT_MQ_QUEUE"),
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	utils.FailOnError(err, "Could not open file")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sub := subFrom("./test/solution.txt", "./test/runner.txt", "./test/testin.txt", "./test/testout.txt")
	body, err := json.Marshal(sub)
	if err != nil {
		log.Panic(err)
		return
	}

	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         []byte(body),
		})
	utils.FailOnError(err, "Could not open file")
	log.Printf(" [x] Sent %s", sub.Id)
}
