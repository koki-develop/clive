package util

import (
	"os"
	"path/filepath"
)

func CreateFile(p string) (*os.File, error) {
	dir := filepath.Dir(p)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return nil, err
	}

	f, err := os.Create(p)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func Exists(p string) (bool, error) {
	if _, err := os.Stat(p); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
