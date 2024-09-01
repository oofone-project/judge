package judges

import (
	"os"
	"os/exec"
)

var (
	BASE_PATH = "./judges"
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
	//err := os.Remove("./judges/" + l.Name + "/submission/" + filename)
	err := os.RemoveAll(BASE_PATH + "/" + l.Name + "/submission/")
	return err
}

func (l Language) Evaluate() (*Results, error) {
	return &Results{}, nil
}
