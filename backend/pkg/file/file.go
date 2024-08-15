package file

import (
	"os"
	"path/filepath"
)

func Exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}

	return true
}

func InitFile(path string) {
	dir := filepath.Dir(path)

	if !Exists(dir) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}

	if !Exists(path) {
		err := os.WriteFile(path, []byte(""), 0644)
		if err != nil {
			panic(err)
		}
	}
}
