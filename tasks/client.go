package tasks

import (
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/oofone-project/judge/judges"
	amqp "github.com/rabbitmq/amqp091-go"
)

type TaskClient struct {
	conn     *amqp.Connection
	channel  *amqp.Channel
	queue    amqp.Queue
	msgs     <-chan amqp.Delivery
	delivery *amqp.Delivery
}

type ResultAndSubmission struct {
	Submission *Submission
	Result     *judges.Result
}

func NewTaskClient() (*TaskClient, error) {
	godotenv.Load(".env")
	conn, err := amqp.Dial(os.Getenv("RABBIT_MQ_URI"))
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		os.Getenv("RABBIT_MQ_QUEUE"), // name
		false,                        // durable
		false,                        // delete when unused
		false,                        // exclusive
		false,                        // no-wait
		nil,                          // arguments
	)
	if err != nil {
		return nil, err
	}

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		return nil, err
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return nil, err
	}

	tc := TaskClient{
		conn:    conn,
		channel: ch,
		queue:   q,
		msgs:    msgs,
	}

	return &tc, nil
}

func (tc *TaskClient) Run(result chan *ResultAndSubmission) error {
	go func() {
		for d := range tc.msgs {
			var sub Submission
			err := json.Unmarshal(d.Body, &sub)
			if err != nil {
				// TODO: send error back and tell client something went wrong
				log.Panic(err)
				continue
			}

			log.Printf("Received new task %s in %s", sub.Id, sub.Language.Name)
			t := NewTask(&sub, &d)

			res, err := t.Run()
			t.Ack(false)
			result <- &ResultAndSubmission{
				Submission: &sub,
				Result:     res,
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	return nil
}

func (tc *TaskClient) Close() {
	tc.conn.Close()
	tc.channel.Close()
}
