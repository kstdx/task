package path

import (
	"os"
	"path/filepath"
)

func Getwd() string {
	cwd, _ := os.Getwd()
	return cwd
}

func Get(one string) string {
	return filepath.Join(Getwd(), one)
}
