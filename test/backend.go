// Simulate the backend adding tasks to the queue
package test

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/oofone-project/judge/judges"
	"github.com/oofone-project/judge/model"
	"github.com/oofone-project/judge/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Backend struct {
	ctx     context.Context
	cancel  context.CancelFunc
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
}

func NewBackend() Backend {
	godotenv.Load("../.env")

	// TODO: use test broker instead, see https://github.com/marketplace/actions/rabbitmq-in-github-actions
	conn, err := amqp.Dial(os.Getenv("RABBIT_MQ_URI"))
	utils.FailOnError(err, "Could not open file")

	ch, err := conn.Channel()
	utils.FailOnError(err, "Could not open file")

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
	return Backend{
		ctx:     ctx,
		cancel:  cancel,
		conn:    conn,
		channel: ch,
		queue:   q,
	}
}

func (b Backend) Publish(s *model.Submission) {
	body, err := json.Marshal(s)
	if err != nil {
		utils.FailOnError(err, "Could not marshal submission")
		return
	}

	err = b.channel.PublishWithContext(b.ctx,
		"",           // exchange
		b.queue.Name, // routing key
		false,        // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         body,
		})
	utils.FailOnError(err, "Could not publish submission task")
	log.Printf(" [x] Sent %s to %s", s.Id, s.Language)
}

func SubFrom(solFile string, runFile string, testIn string, testOut string) *model.Submission {
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
		Language: judges.Python,
		Solution: sol,
		Runner:   run,
		TestIn:   testin,
		TestOut:  testout,
		Id:       id.String(),
	}

	return &sub
}

func (b Backend) Close() {
	b.conn.Close()
	b.cancel()
	b.channel.Close()
}
