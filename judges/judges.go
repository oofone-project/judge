package judges

import (
	"os"
	"os/exec"
)

type Language struct {
	Name    string `json:"name"`
	Ext     string `json:"ext"`
	Command string `json:"command"`
}

var (
	Python = Language{
		Name:    "python",
		Ext:     "py",
		Command: "python3 ./judges/python/gen_out.py",
	}
)

// TODO: create submission/temp folder on run and delete on reset so it can be gitignored
func (l Language) RunJudge() error {
	_, err := exec.Command(l.Ext).Output()
	if err != nil {
		return err
	}
	return nil
}

func (l Language) ResetJudge() error {
	err := l.removeSubmission("testin.txt")
	if err != nil {
		return err
	}
	err = l.removeSubmission("testout.txt")
	if err != nil {
		return err
	}
	err = l.removeSubmission("solution." + l.Ext)
	if err != nil {
		return err
	}
	err = l.removeSubmission("runner." + l.Ext)
	if err != nil {
		return err
	}

	return nil
}

func (l Language) removeSubmission(filename string) error {
	err := os.Remove("./judges/" + l.Name + "/submission/" + filename)
	return err
}
