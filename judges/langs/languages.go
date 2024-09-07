package langs

import (
	"encoding/json"
	"errors"
	"os"
)

type Language struct {
	Name    string `json:"name"`
	Ext     string `json:"ext"`
	Command string `json:"command"`
	Setup   SetupFunc
}

type SetupFunc func() error

var (
	BASE_PATH     = "./languages"
	LanguageError = errors.New("Language does not exist")
	Languages     = map[string]Language{
		"python": Python,
	}
)

func (l Language) Reset() error {
	err := os.RemoveAll(l.SubPath())
	if err != nil {
		return err
	}

	err = os.MkdirAll(l.SubPath(), 0777)
	if err != nil {
		return err
	}

	err = l.Setup()
	return err
}

func (l *Language) SubPath() string {
	return BASE_PATH + "/" + l.Name + "/submission"
}

func (l *Language) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	temp, ok := Languages[s]
	if !ok {
		return LanguageError
	}

	*l = temp
	return nil
}
