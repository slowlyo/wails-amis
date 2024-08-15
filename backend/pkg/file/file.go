package file

import (
	"os"
)

func Exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}

	return true
}

func InitFile(path string) {
	if !Exists(path) {
		err := os.WriteFile(path, []byte(""), 0644)

		if err != nil {
			panic(err)
		}
	}
}
