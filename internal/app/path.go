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

func Path(parts ...string) string {

	items := []string{root}

	items = append(items, parts...)

	return filepath.Join(items...)
}
