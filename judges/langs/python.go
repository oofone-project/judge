package langs

import "os"

var Python = Language{
	Name:    "python",
	Ext:     "py",
	Command: "python3",
	Setup: func() error {
		f, err := os.Create(BASE_PATH + "/python/submission/__init__.py")
		if err != nil {
			return err
		}

		err = f.Close()
		return err
	},
}
