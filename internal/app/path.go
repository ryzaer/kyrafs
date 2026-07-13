package app

import (
	"os"
	"path/filepath"
)

var root string

func init() {

	exe, err := os.Executable()

	if err != nil {
		root = "."
		return
	}

	root = filepath.Dir(exe)
}

func Root() string {
	return root
}

func Path(name string) string {
	return filepath.Join(
		ExecutableDir(),
		name,
	)
}

func ExecutableDir() string {

	exe, err := os.Executable()

	if err != nil {
		return "."
	}

	return filepath.Dir(exe)
}
