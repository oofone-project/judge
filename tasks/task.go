package tasks

import (
	"os"

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

func (t Task) TaskToJudge(basePath string) error {
	err := t.submission.Language.Setup(basePath)
	if err != nil {
		return err
	}
	submissionPath := "../judges/" + t.submission.Language.Name + "/submission/"
	ext := t.submission.Language.Ext

	err = writeTo(submissionPath+"testin.txt", t.submission.TestIn)
	if err != nil {
		return err
	}
	err = writeTo(submissionPath+"testout.txt", t.submission.TestOut)
	if err != nil {
		return err
	}
	err = writeTo(submissionPath+"solution."+ext, t.submission.Solution)
	if err != nil {
		return err
	}
	err = writeTo(submissionPath+"runner."+ext, t.submission.Runner)
	if err != nil {
		return err
	}
	return nil
}

func writeTo(filename string, input []byte) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer func() {
		if tempErr := f.Close(); tempErr != nil {
			err = tempErr
		}
	}()

	_, err = f.Write(input)
	return err
}
