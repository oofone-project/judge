package judges

import (
	"encoding/json"
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
	Python = Language{
		Name:    "python",
		Ext:     "py",
		Command: "python3 ./python/gen_out.py",
		Setup: func() error {
			err := os.MkdirAll(BASE_PATH+"/python/submission", 0777)
			if err != nil {
				return err
			}

			f, err := os.Create(BASE_PATH + "/python/submission/__init__.py")
			if err != nil {
				return err
			}

			err = f.Close()
			return err
		},
	}
	Languages = map[string]Language{
		"python": Python,
	}
)

func (l *Language) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*l = Languages[s]
	return nil
}
