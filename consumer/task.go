package consumer

import (
	"github.com/oofone-project/judge/model"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Task struct {
	submission *model.Submission
	delivery   *amqp.Delivery
}

func NewTask(s *model.Submission, d *amqp.Delivery) Task {
	return Task{
		submission: s,
		delivery:   d,
	}
}

func (t Task) GetSubmission() *model.Submission {
	return t.submission
}

func (t Task) Ack(multiple bool) {
	t.delivery.Ack(multiple)
}
