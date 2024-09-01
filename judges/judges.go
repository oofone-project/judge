package judges

import (
	"os"
	"os/exec"
)

type Results struct{}

// TODO: create submission/temp folder on run and delete on reset so it can be gitignored
func (l Language) RunJudge() error {
	_, err := exec.Command("echo Running judge...").Output()
	if err != nil {
		return err
	}
	return nil
}

func (l Language) ResetJudge() error {
	err := l.removeSubmission()
	if err != nil {
		return err
	}
	return nil
}

func (l Language) removeSubmission() error {
	//err := os.Remove("./judges/" + l.Name + "/submission/" + filename)
	err := os.RemoveAll("./judges/" + l.Name + "/submission/")
	return err
}

func (l Language) Evaluate() (*Results, error) {
	return &Results{}, nil
}
