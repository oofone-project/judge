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
	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func subFrom(solFile string, runFile string, testIn string, testOut string) model.Submission {
	sol, err := os.ReadFile(solFile)
	failOnError(err, "Could not open file")
	run, err := os.ReadFile(runFile)
	failOnError(err, "Could not open file")
	testin, err := os.ReadFile(testIn)
	failOnError(err, "Could not open file")
	testout, err := os.ReadFile(testOut)
	failOnError(err, "Could not open file")

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
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		os.Getenv("RABBIT_MQ_QUEUE"),
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sub := subFrom("./test/solution", "./test/runner", "./test/test.in", "./test/test.out")
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
