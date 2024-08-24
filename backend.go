// Simulate the backend adding tasks to the queue
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Submission struct {
	Solution []byte `json:"solution"`
	Runner   []byte `json:"runner"`
	TestIn   []byte `json:"testin"`
	TestOut  []byte `json:"testout"`
	Id       string `json:"id"`
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func subFrom(solFile string, runFile string, testIn string, testOut string) Submission {
	sol, err := os.ReadFile(solFile)
	if err != nil {
		log.Panic(err)
	}
	run, err := os.ReadFile(runFile)
	if err != nil {
		log.Panic(err)
	}
	testin, err := os.ReadFile(testIn)
	if err != nil {
		log.Panic(err)
	}
	testout, err := os.ReadFile(testOut)
	if err != nil {
		log.Panic(err)
	}
	id := uuid.New()

	sub := Submission{
		Solution: sol,
		Runner:   run,
		TestIn:   testin,
		TestOut:  testout,
		Id:       id.String(),
	}

	return sub
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sub := subFrom("./test/submission", "./test/runner", "./test/test.in", "./test/test.out")
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
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", sub.Id)
}
